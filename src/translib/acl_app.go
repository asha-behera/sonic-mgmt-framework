package translib

import (
    "fmt"
    "bytes"
    "reflect"
    "strconv"
    "strings"
    "errors"
    "translib/db"
    "translib/ocbinds"
    "github.com/openconfig/ygot/ygot"
    "github.com/openconfig/ygot/util"
    log "github.com/golang/glog"
    "github.com/openconfig/gnmi/proto/gnmi"
)

const (
    TABLE_SEPARATOR  = "|"
    KEY_SEPARATOR = "|"
    ACL_TABLE = "ACL_TABLE"
    RULE_TABLE = "ACL_RULE"
    ACL_TYPE = "type"
    ACL_DESCRIPTION = "policy_desc"
    SONIC_ACL_TYPE_L2 = "L2"
    SONIC_ACL_TYPE_IPV4 = "L3"
    SONIC_ACL_TYPE_IPV6 = "L3V6"
    OPENCONFIG_ACL_TYPE_IPV4 = "ACL_IPV4"
    OPENCONFIG_ACL_TYPE_IPV6 = "ACL_IPV6"
    OPENCONFIG_ACL_TYPE_L2 = "ACL_L2"

    MIN_PRIORITY = 1
    MAX_PRIORITY = 65535
)

var IP_PROTOCOL_MAP = map[ocbinds.E_OpenconfigPacketMatchTypes_IP_PROTOCOL]uint8 {
    ocbinds.OpenconfigPacketMatchTypes_IP_PROTOCOL_IP_ICMP:   1,
    ocbinds.OpenconfigPacketMatchTypes_IP_PROTOCOL_IP_IGMP:   2,
    ocbinds.OpenconfigPacketMatchTypes_IP_PROTOCOL_IP_TCP:    6,
    ocbinds.OpenconfigPacketMatchTypes_IP_PROTOCOL_IP_UDP:   17,
    ocbinds.OpenconfigPacketMatchTypes_IP_PROTOCOL_IP_RSVP:  46,
    ocbinds.OpenconfigPacketMatchTypes_IP_PROTOCOL_IP_GRE:   47,
    ocbinds.OpenconfigPacketMatchTypes_IP_PROTOCOL_IP_AUTH:  51,
    ocbinds.OpenconfigPacketMatchTypes_IP_PROTOCOL_IP_PIM:  103,
    ocbinds.OpenconfigPacketMatchTypes_IP_PROTOCOL_IP_L2TP: 115,
}

var ETHERTYPE_MAP = map[ocbinds.E_OpenconfigPacketMatchTypes_ETHERTYPE]uint32 {
    ocbinds.OpenconfigPacketMatchTypes_ETHERTYPE_ETHERTYPE_LLDP: 0x88CC,
    ocbinds.OpenconfigPacketMatchTypes_ETHERTYPE_ETHERTYPE_VLAN: 0x8100,
    ocbinds.OpenconfigPacketMatchTypes_ETHERTYPE_ETHERTYPE_ROCE: 0x8915,
    ocbinds.OpenconfigPacketMatchTypes_ETHERTYPE_ETHERTYPE_ARP:  0x0806,
    ocbinds.OpenconfigPacketMatchTypes_ETHERTYPE_ETHERTYPE_IPV4: 0x0800,
    ocbinds.OpenconfigPacketMatchTypes_ETHERTYPE_ETHERTYPE_IPV6: 0x86DD,
    ocbinds.OpenconfigPacketMatchTypes_ETHERTYPE_ETHERTYPE_MPLS: 0x8847,
}

var aclTs db.TableSpec
var ruleTs db.TableSpec

type AclApp struct {
    path       string
    ygotRoot   *ygot.GoStruct
    ygotTarget *interface{}

    aclTableMap map[string]db.Value
    ruleTableMap map[string]map[string]db.Value

    createAclFlag bool
    createRuleFlag bool
    bindAclFlag bool
}

func init() {
    log.Info("Init called for ACL module")
    err := register("/openconfig-acl:acl",
    &appInfo{appType: reflect.TypeOf(AclApp{}),
    ygotRootType: reflect.TypeOf(ocbinds.OpenconfigAcl_Acl{}),
    isNative:     false})
    if err != nil {
        log.Fatal("Register ACL app module with App Interface failed with error=", err)
    }

    err = addModel(&ModelData{Name: "openconfig-acl",
    Org: "OpenConfig working group",
    Ver:      "1.0.2"})
    if err != nil {
        log.Fatal("Adding model data to appinterface failed with error=", err)
    }
}

func (acl *AclApp) initialize(data appData) {
    log.Info("initialize:acl:path =", data.path)
    *acl = AclApp{path: data.path, ygotRoot: data.ygotRoot, ygotTarget: data.ygotTarget}

    aclTs = db.TableSpec {Name: ACL_TABLE}
    ruleTs = db.TableSpec {Name: RULE_TABLE}

    acl.aclTableMap  = make(map[string]db.Value)
    acl.ruleTableMap  = make(map[string]map[string]db.Value)

    acl.createAclFlag = false
    acl.createRuleFlag = false
    acl.bindAclFlag = false
}

func (acl *AclApp) getAppRootObject() (*ocbinds.OpenconfigAcl_Acl) {
	deviceObj := (*acl.ygotRoot).(*ocbinds.Device)
	return deviceObj.Acl
}

func (acl *AclApp) translateCreate(d *db.DB) ([]db.WatchKeys, error)  {
	var err error
	var keys []db.WatchKeys
    var acl_subtree = false
    //var intf_subtree = false
    log.Info("translateCreate:acl:path =", acl.path)

    aclObj := acl.getAppRootObject()
    if reflect.TypeOf(*acl.ygotTarget).Elem().Name() == "OpenconfigAcl_Acl" {
        acl_subtree = true
        //intf_subtree = true
    }

    acl.aclTableMap = convert_oc_acls_to_internal(aclObj)
    acl.ruleTableMap = convert_oc_acl_rules_to_internal(aclObj)
    acl.bindAclFlag = convert_oc_acl_bindings_to_internal(acl.aclTableMap, aclObj)

    // These slices will store the yangPaths derived from the URI requested to help
    // determining when to create ACL or rule or both
    var ruleBasedTargets []string = []string{}
    var aclBasedTargets []string = []string{getYangPathFromStruct(aclObj), getYangPathFromStruct(aclObj.AclSets)}

    targetUriPath, err := getYangPathFromUri(acl.path)
    if isSubtreeRequest(targetUriPath, "/openconfig-acl:acl/acl-sets") || acl_subtree {
        if aclObj.AclSets != nil && len(aclObj.AclSets.AclSet) > 0 {
            // Build Watch keys for a specific ACL
            for aclSetKey,_ := range aclObj.AclSets.AclSet {
                aclName := strings.ReplaceAll(strings.ReplaceAll(aclSetKey.Name, " ", "_"), "-", "_")
                aclType := aclSetKey.Type.ΛMap()["E_OpenconfigAcl_ACL_TYPE"][int64(aclSetKey.Type)].Name
                aclKey := aclName + "_" + aclType
                keys = append(keys, db.WatchKeys{ &aclTs, &(db.Key{Comp:[]string{aclKey}})})

                aclSet := aclObj.AclSets.AclSet[aclSetKey]
                aclBasedTargets = append(aclBasedTargets, getYangPathFromStruct(aclSet))
                ruleBasedTargets = append(ruleBasedTargets, getYangPathFromStruct(aclSet.AclEntries))

                if aclSet.AclEntries != nil && len(aclSet.AclEntries.AclEntry) > 0 {
                    // Build Watch keys for a specific Rule
                    for seqId,_ := range aclSet.AclEntries.AclEntry {
                        ruleName := "RULE_" + strconv.FormatInt(int64(seqId), 10)
                        keys = append(keys, db.WatchKeys{ &ruleTs, &(db.Key{Comp:[]string{aclKey,ruleName}})})
                        ruleBasedTargets = append(ruleBasedTargets, getYangPathFromStruct(aclSet.AclEntries.AclEntry[seqId]))
                    }
                } else {
                    // Build watch keys for all rules for a specific ACL
                    for ruleName,_ := range acl.ruleTableMap[aclKey] {
                        keys = append(keys, db.WatchKeys{ &ruleTs, &db.Key{Comp:[]string{aclKey, ruleName}} })
                    }
                }
            }
        } else {
            // Building Watch keys for Create Request for All ACLs and Rules
            for aclName,_ := range acl.aclTableMap {
                keys = append(keys, db.WatchKeys{ &aclTs, &db.Key{Comp:[]string{aclName}} })
                for ruleName,_ := range acl.ruleTableMap[aclName] {
                    keys = append(keys, db.WatchKeys{ &ruleTs, &db.Key{Comp:[]string{aclName, ruleName}} })
                }
            }
        }
    }

    /*
    if isSubtreeRequest(targetUriPath, "/openconfig-acl:acl/interfaces") || intfSubtree {
    }
    */

    if contains(aclBasedTargets, targetUriPath) {
        acl.createAclFlag = true
        acl.createRuleFlag = true
    }
    if contains(ruleBasedTargets, targetUriPath) {
        acl.createRuleFlag = true
    }

    //err = errors.New("Not implemented")
    return keys, err
}

func (acl *AclApp) translateUpdate(d *db.DB) ([]db.WatchKeys, error)  {
    var err error
    var keys []db.WatchKeys
    log.Info("translateUpdate:acl:path =", acl.path)
    err = errors.New("Not implemented")
    return keys, err
}

func (acl *AclApp) translateReplace(d *db.DB) ([]db.WatchKeys, error)  {
    var err error
    var keys []db.WatchKeys
    log.Info("translateReplace:acl:path =", acl.path)
    err = errors.New("Not implemented")
    return keys, err
}

func (acl *AclApp) translateDelete(d *db.DB) ([]db.WatchKeys, error)  {
    var err error
    var keys []db.WatchKeys
    var acl_subtree = false
    //var intf_subtree = false
    log.Info("translateDelete:acl:path =", acl.path)

    aclObj := acl.getAppRootObject()
    if reflect.TypeOf(*acl.ygotTarget).Elem().Name() == "OpenconfigAcl_Acl" {
        acl_subtree = true
        //intf_subtree = true
    }

    fmt.Println("translateDelete: Target Type: " + reflect.TypeOf(*acl.ygotTarget).Elem().Name())
    targetUriPath, err := getYangPathFromUri(acl.path)
    if isSubtreeRequest(targetUriPath, "/openconfig-acl:acl/acl-sets") || acl_subtree {
        if aclObj.AclSets != nil && len(aclObj.AclSets.AclSet) > 0 {
            // Build Watch keys for a specific ACL
            for aclSetKey,_ := range aclObj.AclSets.AclSet {
                aclName := strings.ReplaceAll(strings.ReplaceAll(aclSetKey.Name, " ", "_"), "-", "_")
                aclType := aclSetKey.Type.ΛMap()["E_OpenconfigAcl_ACL_TYPE"][int64(aclSetKey.Type)].Name
                aclKey := aclName + "_" + aclType
                keys = append(keys, db.WatchKeys{ &aclTs, &(db.Key{Comp:[]string{aclKey}})})

                aclSet := aclObj.AclSets.AclSet[aclSetKey]
                if aclSet.AclEntries != nil && len(aclSet.AclEntries.AclEntry) > 0 {
                    // Build Watch keys for a specific Rule
                    for seqId,_ := range aclSet.AclEntries.AclEntry {
                        ruleName := "RULE_" + strconv.FormatInt(int64(seqId), 10)
                        keys = append(keys, db.WatchKeys{ &ruleTs, &(db.Key{Comp:[]string{aclKey,ruleName}})})
                    }
                } else {
                    // Build watch keys for all rules for a specific ACL
                    ruleKeys,_ := d.GetKeys(&ruleTs)
                    for i,rulekey := range ruleKeys {
                        // Rulekey has two keys, first aclkey and second rulename
                        if rulekey.Comp[0] == aclKey {
                            keys = append(keys, db.WatchKeys{ &ruleTs, &ruleKeys[i]})
                        }
                    }
                }
            }
        } else {
            // Building Watch keys for Delete Request for All ACLs and Rules
            aclKeys,_ := d.GetKeys(&aclTs)
            ruleKeys,_ := d.GetKeys(&ruleTs)

            for i,_ := range aclKeys {
                keys = append(keys, db.WatchKeys{ &aclTs, &aclKeys[i]})
            }
            for i,_ := range ruleKeys {
                keys = append(keys, db.WatchKeys{ &ruleTs, &ruleKeys[i]})
            }
        }
    }

    /*
    if isSubtreeRequest(targetUriPath, "/openconfig-acl:acl/interfaces") || intfSubtree {
    }
    */

    //err = errors.New("Not implemented")
    return keys, err
}

func (acl *AclApp) translateGet(dbs [db.MaxDB]*db.DB) error  {
    var err error
    log.Info("translateGet:acl:path =", acl.path)
    return err
}

func (acl *AclApp) processCreate(d *db.DB) (SetResponse, error)  {
    var err error
    var resp SetResponse

    log.Info("processCreate:acl:path =", acl.path)
    log.Info("ProcessCreate: Target Type is " + reflect.TypeOf(*acl.ygotTarget).Elem().Name())

    if acl.createAclFlag {
        set_acl_data_in_config_db(d, acl.aclTableMap)
    }
    if acl.createRuleFlag {
        set_acl_rule_data_in_config_db(d, acl.ruleTableMap)
    }
    if acl.bindAclFlag && !acl.createAclFlag {
        //set_acl_bind_data_in_config_db(d, acl.aclTableMap)
    }

	//err = errors.New("Not implemented")
	return resp, err
}

func (acl *AclApp) processUpdate(d *db.DB) (SetResponse, error)  {
    var err error
    var resp SetResponse
    log.Info("processUpdate:acl:path =", acl.path)
    err = errors.New("Not implemented")
    return resp, err
}

func (acl *AclApp) processReplace(d *db.DB) (SetResponse, error)  {
    var err error
    var resp SetResponse
    log.Info("processReplace:acl:path =", acl.path)
    err = errors.New("Not implemented")
    return resp, err
}

func (acl *AclApp) processDelete(d *db.DB) (SetResponse, error)  {
    var err error
    var resp SetResponse
    var acl_subtree = false
    //var intf_subtree = false
    log.Info("processDelete:acl:path =", acl.path)

    aclObj := acl.getAppRootObject()
    if reflect.TypeOf(*acl.ygotTarget).Elem().Name() == "OpenconfigAcl_Acl" {
        acl_subtree = true
        //intf_subtree = true
    }
    targetUriPath, err := getYangPathFromUri(acl.path)
    if isSubtreeRequest(targetUriPath, "/openconfig-acl:acl/acl-sets") || acl_subtree {
        if aclObj.AclSets != nil && len(aclObj.AclSets.AclSet) > 0 {
            // Deletion of a specific ACL
            for aclSetKey,_ := range aclObj.AclSets.AclSet {
                aclName := strings.ReplaceAll(strings.ReplaceAll(aclSetKey.Name, " ", "_"), "-", "_")
                aclType := aclSetKey.Type.ΛMap()["E_OpenconfigAcl_ACL_TYPE"][int64(aclSetKey.Type)].Name
                aclKey := aclName + "_" + aclType

                aclSet := aclObj.AclSets.AclSet[aclSetKey]
                if aclSet.AclEntries != nil && len(aclSet.AclEntries.AclEntry) > 0 {
                    // Deletion of a specific Rule
                    for seqId,_ := range aclSet.AclEntries.AclEntry {
                        ruleName := "RULE_" + strconv.FormatInt(int64(seqId), 10)
                        d.DeleteEntry(&ruleTs, db.Key{Comp: []string {aclKey, ruleName} })
                    }
                } else {
                    // Deletion of a specific Acl and all its rule
                    if *acl.ygotTarget == aclSet {
                        d.DeleteKeys(&ruleTs, db.Key{Comp: []string {aclKey + TABLE_SEPARATOR + "*"} })
                        d.DeleteEntry(&aclTs, db.Key{Comp: []string {aclKey} })
                    }
                    // Deletion of all rules for a specific ACL but NOT ACL
                    if *acl.ygotTarget == aclSet.AclEntries {
                        d.DeleteKeys(&ruleTs, db.Key{Comp: []string {aclKey + TABLE_SEPARATOR + "*"} })
                    } else {
                        d.DeleteKeys(&ruleTs, db.Key{Comp: []string {aclKey + TABLE_SEPARATOR + "*"} })
                    }
                }
            }
        } else {
            // Deletion of All ACLs and Rules
            d.DeleteTable(&aclTs)
            d.DeleteTable(&ruleTs)
        }
    }

    /*
    if isSubtreeRequest(targetUriPath, "/openconfig-acl:acl/interfaces") || intfSubtree {
    }
    */

    //err = errors.New("Not implemented")
    return resp, err
}

func (acl *AclApp) processGet(dbs [db.MaxDB]*db.DB) (GetResponse, error)  {
    var err error
    var payload []byte
    var aclSubtree bool = false
    var intfSubtree bool = false

    configDb := dbs[db.ConfigDB]
    aclObj := acl.getAppRootObject()

    log.Info("processGet: Target Type: " + reflect.TypeOf(*acl.ygotTarget).Elem().Name())
    if reflect.TypeOf(*acl.ygotTarget).Elem().Name() == "OpenconfigAcl_Acl" {
        aclSubtree = true
        intfSubtree = true
    }

    targetUriPath, err := getYangPathFromUri(acl.path)
    if isSubtreeRequest(targetUriPath, "/openconfig-acl:acl/acl-sets") || aclSubtree {
        if aclObj.AclSets != nil && len(aclObj.AclSets.AclSet) > 0 {
            // Request for specific ACL
            for aclSetKey,_ := range aclObj.AclSets.AclSet {
                aclName := strings.ReplaceAll(strings.ReplaceAll(aclSetKey.Name, " ", "_"), "-", "_")
                aclType := aclSetKey.Type.ΛMap()["E_OpenconfigAcl_ACL_TYPE"][int64(aclSetKey.Type)].Name
                aclSet := aclObj.AclSets.AclSet[aclSetKey]
                aclKey := aclName + "_" + aclType

                if aclSet.AclEntries != nil && len(aclSet.AclEntries.AclEntry) > 0 {
                    // Request for specific Rule
                    for seqId,_ := range aclSet.AclEntries.AclEntry {
                        //ruleKey := "RULE_" + strconv.FormatInt(int64(seqId), 10)
                        entrySet := aclSet.AclEntries.AclEntry[seqId]
                        err = acl.convert_db_acl_rules_to_internal(configDb, aclKey, int64(seqId), db.Key{})
                        if (err != nil) {
                            return GetResponse{Payload:payload, ErrSrc:AppErr}, err
                        }
                        ygot.BuildEmptyTree(entrySet)
                        acl.convert_internal_to_oc_acl_rule(aclKey, aclSetKey.Type, int64(seqId), nil, entrySet)

                        if *acl.ygotTarget == entrySet {
                            payload, err = dumpIetfJson(aclSet.AclEntries)
                        } else {
                            dummyEntrySet := &ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry{}
                            if *acl.ygotTarget == entrySet.Config {
                                dummyEntrySet.Config = entrySet.Config
                                payload, err = dumpIetfJson(dummyEntrySet)
                            } else if *acl.ygotTarget == entrySet.State {
                                dummyEntrySet.State = entrySet.State
                                payload, err = dumpIetfJson(dummyEntrySet)
                            } else if *acl.ygotTarget == entrySet.Actions {
                                dummyEntrySet.Actions = entrySet.Actions
                                payload, err = dumpIetfJson(dummyEntrySet)
                            } else if *acl.ygotTarget == entrySet.InputInterface {
                                dummyEntrySet.InputInterface = entrySet.InputInterface
                                payload, err = dumpIetfJson(dummyEntrySet)
                            } else if *acl.ygotTarget == entrySet.Ipv4 {
                                dummyEntrySet.Ipv4 = entrySet.Ipv4
                                payload, err = dumpIetfJson(dummyEntrySet)
                            } else if *acl.ygotTarget == entrySet.Ipv6 {
                                dummyEntrySet.Ipv6 = entrySet.Ipv6
                                payload, err = dumpIetfJson(dummyEntrySet)
                            } else if *acl.ygotTarget == entrySet.L2 {
                                dummyEntrySet.L2 = entrySet.L2
                                payload, err = dumpIetfJson(dummyEntrySet)
                            } else if *acl.ygotTarget == entrySet.Transport {
                                dummyEntrySet.Transport = entrySet.Transport
                                payload, err = dumpIetfJson(dummyEntrySet)
                            } else {
                            }
                        }
                    }
                } else {
                    err = acl.convert_db_acl_to_internal(configDb, db.Key{Comp: []string {aclKey} })
                    if (err != nil) {
                        return GetResponse{Payload:payload, ErrSrc:AppErr}, err
                    }

                    ygot.BuildEmptyTree(aclSet)
                    acl.convert_internal_to_oc_acl(aclKey, aclObj.AclSets, aclSet)

                    if *acl.ygotTarget == aclSet {
                        payload, err = dumpIetfJson(aclObj.AclSets)
                    } else {
                        dummyAclSet := &ocbinds.OpenconfigAcl_Acl_AclSets_AclSet{}
                        if *acl.ygotTarget == aclSet.Config {
                            dummyAclSet.Config = aclSet.Config
                            payload, err = dumpIetfJson(dummyAclSet)
                        } else if *acl.ygotTarget == aclSet.State {
                            dummyAclSet.State = aclSet.State
                            payload, err = dumpIetfJson(dummyAclSet)
                        } else if *acl.ygotTarget == aclSet.AclEntries {
                            dummyAclSet.AclEntries = aclSet.AclEntries
                            payload, err = dumpIetfJson(dummyAclSet)
                        } else {
                            if targetUriPath == "/openconfig-acl:acl/acl-sets/acl-set/acl-entries/acl-entry" {
                                dummyAclSet.AclEntries = aclSet.AclEntries
                                payload, err = dumpIetfJson(dummyAclSet.AclEntries)
                            } else if targetUriPath == "/openconfig-acl:acl/acl-sets/acl-set/config/description" {
                                dummyAclSet.Config = &ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_Config{}
                                dummyAclSet.Config.Description = aclSet.Config.Description
                                payload, err = dumpIetfJson(dummyAclSet.Config)
                            }
                        }
                    }
                }
            }
        } else {
            // Request for all ACLs
            ygot.BuildEmptyTree(aclObj)
            err = acl.convert_db_acl_to_internal(configDb, db.Key{})
            if (err != nil) {
                return GetResponse{Payload:payload, ErrSrc:AppErr}, err
            }

            acl.convert_internal_to_oc_acl("", aclObj.AclSets, nil)
            if (err != nil) {
                return GetResponse{Payload:payload, ErrSrc:AppErr}, err
            }

            if reflect.TypeOf(*acl.ygotTarget).Elem().Name() == "OpenconfigAcl_Acl" {
                //payload, err = dumpIetfJson(&ocbinds.Device{Acl:aclObj})
            } else if *acl.ygotTarget == aclObj.AclSets {
                payload, err = dumpIetfJson(aclObj)
            } else {
                if targetUriPath == "/openconfig-acl:acl/acl-sets/acl-set" {
                    payload, err = dumpIetfJson(aclObj.AclSets)
                }
            }
        }
    }

    if isSubtreeRequest(targetUriPath, "/openconfig-acl:acl/interfaces") || intfSubtree {
        if aclObj.Interfaces != nil && len(aclObj.Interfaces.Interface) > 0 {
            var intfData *ocbinds.OpenconfigAcl_Acl_Interfaces_Interface
            for intfId := range aclObj.Interfaces.Interface {
                intfData = aclObj.Interfaces.Interface[intfId]
                ygot.BuildEmptyTree(intfData)
                if isSubtreeRequest(targetUriPath, "/openconfig-acl:acl/interfaces/interface/ingress-acl-sets") {
                    // Ingress ACL Specific
                    acl.get_acl_binding_info_for_subtree(configDb, intfData, intfId, "INGRESS")
                } else if isSubtreeRequest(targetUriPath, "/openconfig-acl:acl/interfaces/interface/egress-acl-sets") {
                    // Egress ACL Specific
                    acl.get_acl_binding_info_for_subtree(configDb, intfData, intfId, "EGRESS")
                } else {
                    // Direction unknown. Check ACL Table for binding information.
                    fmt.Println("Request is for specific interface, ingress and egress ACLs")
                    acl.get_acl_binding_info_for_subtree(configDb, intfData, intfId, "INGRESS")
                    acl.get_acl_binding_info_for_subtree(configDb, intfData, intfId, "EGRESS")
                }
            }

            if *acl.ygotTarget == intfData {
                payload, err = dumpIetfJson(aclObj.Interfaces)
            } else {
                dummyIntfData := &ocbinds.OpenconfigAcl_Acl_Interfaces_Interface{}
                if  *acl.ygotTarget == intfData.Config {
                    dummyIntfData.Config = intfData.Config
                    payload, err = dumpIetfJson(dummyIntfData)
                } else if  *acl.ygotTarget == intfData.State {
                    dummyIntfData.State = intfData.State
                    payload, err = dumpIetfJson(dummyIntfData)
                } else if *acl.ygotTarget == intfData.IngressAclSets {
                    dummyIntfData.IngressAclSets = intfData.IngressAclSets
                    payload, err = dumpIetfJson(dummyIntfData)
                } else if *acl.ygotTarget == intfData.EgressAclSets {
                    dummyIntfData.EgressAclSets = intfData.EgressAclSets
                    payload, err = dumpIetfJson(dummyIntfData)
                } else if *acl.ygotTarget == intfData.InterfaceRef {
                    dummyIntfData.InterfaceRef = intfData.InterfaceRef
                    payload, err = dumpIetfJson(dummyIntfData)
                }
            }
        } else {
            fmt.Println("Request is for all interfaces and all directions on which ACL is applied")
            if len(acl.aclTableMap) == 0 {
                // Get all ACLs
                acl.convert_db_acl_to_internal(configDb, db.Key{})
            }

            var interfaces []string
            for aclName := range acl.aclTableMap {
                aclData := acl.aclTableMap[aclName]
                if len(aclData.Get("ports@")) > 0 {
                    aclIntfs := aclData.GetList("ports")
                    for i,_ := range aclIntfs {
                        if !contains(interfaces, aclIntfs[i]) && aclIntfs[i] != "" {
                            interfaces = append(interfaces, aclIntfs[i])
                        }
                    }
                }
            }

            for _, intfId := range interfaces {
                var intfData *ocbinds.OpenconfigAcl_Acl_Interfaces_Interface
                intfData,ok := aclObj.Interfaces.Interface[intfId]; if !ok {
                    intfData,_ = aclObj.Interfaces.NewInterface(intfId)
                }
                ygot.BuildEmptyTree(intfData)
                acl.get_acl_binding_info_for_subtree(configDb, intfData, intfId, "INGRESS")
                acl.get_acl_binding_info_for_subtree(configDb, intfData, intfId, "EGRESS")
            }

            if *acl.ygotTarget == aclObj.Interfaces {
                payload, err = dumpIetfJson(aclObj)
            } else {
                if targetUriPath == "/openconfig-acl:acl/interfaces/interface" {
                    payload, err = dumpIetfJson(aclObj.Interfaces)
                }
            }
        }
    }

    if reflect.TypeOf(*acl.ygotTarget).Elem().Name() == "OpenconfigAcl_Acl" {
        //payload, err = dumpIetfJson((*acl.ygotRoot).(*ocbinds.Device))
        payload, err = dumpIetfJson(aclObj)
    }

    return GetResponse{Payload:payload}, err
}

/***********    These are Translation Helper Function   ***********/
func (acl *AclApp) convert_db_acl_rules_to_internal(dbCl *db.DB, aclName string , seqId int64, ruleKey db.Key) error {
    var err error
    if seqId != -1 {
        ruleKey.Comp = []string {aclName, "RULE_" + strconv.FormatInt(int64(seqId), 10)}
    }
    if ruleKey.Len() > 1 {
        ruleName := ruleKey.Get(1)
        if ruleName != "DEFAULT_RULE" {
            ruleData, err := dbCl.GetEntry(&ruleTs, ruleKey)
            if err != nil {
                return err
            }
            if acl.ruleTableMap[aclName] == nil {
                acl.ruleTableMap[aclName] = make(map[string]db.Value)
            }
            acl.ruleTableMap[aclName][ruleName] = ruleData
        }
    } else {
        ruleKeys, err := dbCl.GetKeys(&ruleTs)
        if err != nil {
            return err
        }
        for i, _ := range ruleKeys {
            if aclName == ruleKeys[i].Get(0) {
                acl.convert_db_acl_rules_to_internal(dbCl, aclName, -1, ruleKeys[i])
            }
        }
    }
    return err
}

func (acl *AclApp) convert_db_acl_to_internal(dbCl *db.DB, aclkey db.Key) error {
    var err error
    if aclkey.Len() > 0 {
        // Get one particular ACL
        entry, err := dbCl.GetEntry(&aclTs, aclkey)
        if err != nil {
            return err
        }
        if entry.IsPopulated() {
            acl.aclTableMap[aclkey.Get(0)] = entry
            acl.ruleTableMap[aclkey.Get(0)] = make(map[string]db.Value)
            err = acl.convert_db_acl_rules_to_internal(dbCl, aclkey.Get(0), -1, db.Key{})
            if err != nil {
                return err
            }
        } else {
            return errors.New("ACL is not configured")
        }
    } else {
        // Get all ACLs
        tbl,err := dbCl.GetTable(&aclTs)
        if err != nil {
            return err
        }
        keys, _ := tbl.GetKeys()
        for i, _ := range keys {
            acl.convert_db_acl_to_internal(dbCl, keys[i])
        }
    }
    return err
}

func (acl *AclApp) convert_internal_to_oc_acl(aclName string, aclSets *ocbinds.OpenconfigAcl_Acl_AclSets, aclSet *ocbinds.OpenconfigAcl_Acl_AclSets_AclSet) {
    if len(aclName) > 0 {
        aclData := acl.aclTableMap[aclName]
        if aclSet != nil {
            aclSet.Config.Name = aclSet.Name
            aclSet.Config.Type = aclSet.Type
            aclSet.State.Name = aclSet.Name
            aclSet.State.Type = aclSet.Type

            for k := range aclData.Field {
                if ACL_DESCRIPTION == k {
                    descr := aclData.Get(k)
                    aclSet.Config.Description = &descr
                    aclSet.State.Description = &descr
                } else if "ports@" == k {
                    continue
                    //convert_db_to_oc_acl_bindings
                }
            }

            acl.convert_internal_to_oc_acl_rule(aclName, aclSet.Type, -1, aclSet, nil)
        }
    } else {
        for acln := range acl.aclTableMap {
            acldata := acl.aclTableMap[acln]
            var aclNameStr string
            var aclType ocbinds.E_OpenconfigAcl_ACL_TYPE
            if acldata.Get(ACL_TYPE) == SONIC_ACL_TYPE_IPV4 {
                aclNameStr = strings.Replace(acln, "_"+OPENCONFIG_ACL_TYPE_IPV4, "", 1)
                aclType = ocbinds.OpenconfigAcl_ACL_TYPE_ACL_IPV4
            } else if acldata.Get(ACL_TYPE) == SONIC_ACL_TYPE_IPV6 {
                aclNameStr = strings.Replace(acln, "_"+OPENCONFIG_ACL_TYPE_IPV6, "", 1)
                aclType = ocbinds.OpenconfigAcl_ACL_TYPE_ACL_IPV6
            } else if acldata.Get(ACL_TYPE) == SONIC_ACL_TYPE_L2 {
                aclNameStr = strings.Replace(acln, "_"+OPENCONFIG_ACL_TYPE_L2, "", 1)
                aclType = ocbinds.OpenconfigAcl_ACL_TYPE_ACL_L2
            }
            aclSetPtr, aclErr := aclSets.NewAclSet(aclNameStr, aclType) ; if (aclErr != nil) {
                fmt.Println("Error handling: ", aclErr)
            }
            ygot.BuildEmptyTree(aclSetPtr)
            acl.convert_internal_to_oc_acl(acln, nil, aclSetPtr)
        }
    }
}

func (acl *AclApp) convert_internal_to_oc_acl_rule(aclName string, aclType ocbinds.E_OpenconfigAcl_ACL_TYPE, seqId int64, aclSet *ocbinds.OpenconfigAcl_Acl_AclSets_AclSet, entrySet *ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry) {
    if seqId != -1 {
        ruleName := "RULE_" + strconv.FormatInt(int64(seqId), 10)
        acl.convert_internal_to_oc_acl_rule_properties(acl.ruleTableMap[aclName][ruleName], aclType, nil, entrySet)
    } else {
        for ruleName := range acl.ruleTableMap[aclName] {
            acl.convert_internal_to_oc_acl_rule_properties(acl.ruleTableMap[aclName][ruleName], aclType, aclSet, nil)
        }
    }
}

func (acl *AclApp) convert_internal_to_oc_acl_rule_properties(ruleData db.Value, aclType ocbinds.E_OpenconfigAcl_ACL_TYPE, aclSet *ocbinds.OpenconfigAcl_Acl_AclSets_AclSet, entrySet *ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry) {
    priority,_ := strconv.ParseInt(ruleData.Get("PRIORITY"), 10, 32)
    seqId := uint32(MAX_PRIORITY - priority)
    ruleDescr := ruleData.Get("RULE_DESCRIPTION")

    if entrySet == nil {
        if aclSet != nil {
            entrySet_, _ := aclSet.AclEntries.NewAclEntry(seqId)
            entrySet = entrySet_
            ygot.BuildEmptyTree(entrySet)
        }
    }

    entrySet.Config.SequenceId = &seqId
    entrySet.Config.Description = &ruleDescr
    entrySet.State.SequenceId = &seqId
    entrySet.State.Description = &ruleDescr

    var num uint64
    num = 0
    entrySet.State.MatchedOctets = &num
    entrySet.State.MatchedPackets = &num

    ygot.BuildEmptyTree(entrySet.Transport)
    ygot.BuildEmptyTree(entrySet.Actions)

    for ruleKey := range ruleData.Field {
        if "L4_SRC_PORT" == ruleKey {
            port := ruleData.Get(ruleKey)
            entrySet.Transport.Config.SourcePort = getTransportConfigSrcPort(port)
            //entrySet.Transport.State.SourcePort = &addr
        } else if "L4_DST_PORT" == ruleKey {
            port := ruleData.Get(ruleKey)
            entrySet.Transport.Config.DestinationPort = getTransportConfigDestPort(port)
            //entrySet.Transport.State.DestinationPort = &addr
        } else if "PACKET_ACTION" == ruleKey {
            if "FORWARD" == ruleData.Get(ruleKey) {
                entrySet.Actions.Config.ForwardingAction = ocbinds.OpenconfigAcl_FORWARDING_ACTION_ACCEPT
                //entrySet.Actions.State.ForwardingAction = ocbinds.OpenconfigAcl_FORWARDING_ACTION_ACCEPT
            } else {
                entrySet.Actions.Config.ForwardingAction = ocbinds.OpenconfigAcl_FORWARDING_ACTION_DROP
                //entrySet.Actions.State.ForwardingAction = ocbinds.OpenconfigAcl_FORWARDING_ACTION_DROP
            }
        }
    }

    if aclType == ocbinds.OpenconfigAcl_ACL_TYPE_ACL_IPV4 {
        ygot.BuildEmptyTree(entrySet.Ipv4)
        for ruleKey := range ruleData.Field {
            if "IP_PROTOCOL" == ruleKey {
                ipProto, _ := strconv.ParseInt(ruleData.Get(ruleKey), 10, 64)
                ipv4ProElem := getIpProtocol(ipProto, ocbinds.OpenconfigAcl_ACL_TYPE_ACL_IPV4, "config")
                entrySet.Ipv4.Config.Protocol = ipv4ProElem.(*ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv4_Config_Protocol_Union_E_OpenconfigPacketMatchTypes_IP_PROTOCOL)

                ipv4ProElem = getIpProtocol(ipProto, ocbinds.OpenconfigAcl_ACL_TYPE_ACL_IPV4, "state")
                entrySet.Ipv4.State.Protocol = ipv4ProElem.(*ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv4_State_Protocol_Union_E_OpenconfigPacketMatchTypes_IP_PROTOCOL)
            } else if "DSCP" == ruleKey {
                var dscp uint8
                dscpData, _ := strconv.ParseInt(ruleData.Get(ruleKey), 10, 64)
                dscp = uint8(dscpData)
                entrySet.Ipv4.Config.Dscp = &dscp
                entrySet.Ipv4.State.Dscp = &dscp
            } else if "SRC_IP" == ruleKey {
                addr := ruleData.Get(ruleKey)
                entrySet.Ipv4.Config.SourceAddress = &addr
                entrySet.Ipv4.State.SourceAddress = &addr
            } else if "DST_IP" == ruleKey {
                addr := ruleData.Get(ruleKey)
                entrySet.Ipv4.Config.DestinationAddress = &addr
                entrySet.Ipv4.State.DestinationAddress = &addr
            }
        }
    } else if aclType == ocbinds.OpenconfigAcl_ACL_TYPE_ACL_IPV6 {
        ygot.BuildEmptyTree(entrySet.Ipv6)
        for ruleKey := range ruleData.Field {
            if "IP_PROTOCOL" == ruleKey {
                ipProto, _ := strconv.ParseInt(ruleData.Get(ruleKey), 10, 64)
                ipv6ProElem := getIpProtocol(ipProto, ocbinds.OpenconfigAcl_ACL_TYPE_ACL_IPV6, "config")
                entrySet.Ipv6.Config.Protocol = ipv6ProElem.(*ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv6_Config_Protocol_Union_E_OpenconfigPacketMatchTypes_IP_PROTOCOL)

                ipv6ProElem = getIpProtocol(ipProto, ocbinds.OpenconfigAcl_ACL_TYPE_ACL_IPV6, "state")
                entrySet.Ipv6.State.Protocol = ipv6ProElem.(*ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv6_State_Protocol_Union_E_OpenconfigPacketMatchTypes_IP_PROTOCOL)
            } else if "DSCP" == ruleKey {
                var dscp uint8
                dscpData, _ := strconv.ParseInt(ruleData.Get(ruleKey), 10, 64)
                dscp = uint8(dscpData)
                entrySet.Ipv6.Config.Dscp = &dscp
                entrySet.Ipv6.State.Dscp = &dscp
            }
        }
    } else if aclType == ocbinds.OpenconfigAcl_ACL_TYPE_ACL_L2 {
        ygot.BuildEmptyTree(entrySet.L2)
        for ruleKey := range ruleData.Field {
            if "ETHER_TYPE" == ruleKey {
                ethType, _ := strconv.ParseInt(ruleData.Get(ruleKey), 10, 64)
                fmt.Println(ethType)
                //entrySet.L2.Config.Ethertype = ""
                //entrySet.Ipv6.State.Protocol = getIpProtocolState(ipProto)
            }
        }
    }
}

func convert_internal_to_oc_acl_rule_binding(d *db.DB, priority uint32, seqId int64, direction string, aclSet ygot.GoStruct, entrySet ygot.GoStruct) {
    if seqId == -1 {
        seqId = int64(MAX_PRIORITY - priority)
    }

    var num uint64
    num = 0
    var ruleId uint32 = uint32(seqId)

    if direction == "INGRESS" {
        var ingressEntrySet *ocbinds.OpenconfigAcl_Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_AclEntries_AclEntry
        var ok bool
        if entrySet == nil {
            ingressAclSet := aclSet.(*ocbinds.OpenconfigAcl_Acl_Interfaces_Interface_IngressAclSets_IngressAclSet)
            if ingressEntrySet,ok = ingressAclSet.AclEntries.AclEntry[ruleId]; !ok {
                ingressEntrySet,_ = ingressAclSet.AclEntries.NewAclEntry(ruleId)
            }
        } else {
            ingressEntrySet = entrySet.(*ocbinds.OpenconfigAcl_Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_AclEntries_AclEntry)
        }
        if ingressEntrySet != nil {
            ygot.BuildEmptyTree(ingressEntrySet)
            ingressEntrySet.State.SequenceId = &ruleId
            ingressEntrySet.State.MatchedPackets = &num
            ingressEntrySet.State.MatchedOctets = &num
        }
    } else if direction == "EGRESS" {
        var egressEntrySet *ocbinds.OpenconfigAcl_Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_AclEntries_AclEntry
        var ok bool
        if entrySet == nil {
            egressAclSet := aclSet.(*ocbinds.OpenconfigAcl_Acl_Interfaces_Interface_EgressAclSets_EgressAclSet)
            if egressEntrySet,ok = egressAclSet.AclEntries.AclEntry[ruleId]; !ok {
                egressEntrySet,_ = egressAclSet.AclEntries.NewAclEntry(ruleId)
            }
        } else {
            egressEntrySet = entrySet.(*ocbinds.OpenconfigAcl_Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_AclEntries_AclEntry)
        }
        if egressEntrySet != nil {
            ygot.BuildEmptyTree(egressEntrySet)
            egressEntrySet.State.SequenceId = &ruleId
            egressEntrySet.State.MatchedPackets = &num
            egressEntrySet.State.MatchedOctets = &num
        }
    }
}

func (acl *AclApp) convert_internal_to_oc_acl_binding (d *db.DB, aclName string, intfId string, direction string, intfAclSet ygot.GoStruct) {
    if _,ok := acl.aclTableMap[aclName]; !ok {
        acl.convert_db_acl_to_internal(d, db.Key{Comp: []string {aclName} })
    }

    if _,ok := acl.ruleTableMap[aclName]; !ok {
        acl.convert_db_acl_rules_to_internal(d, aclName, -1, db.Key{})
    }

    for ruleName,_ := range acl.ruleTableMap[aclName] {
        ruleData := acl.ruleTableMap[aclName][ruleName]
        priority,_ := strconv.ParseInt(ruleData.Get("PRIORITY"), 10, 32)
        convert_internal_to_oc_acl_rule_binding(d, uint32(priority), -1, direction, intfAclSet, nil)
    }
}

func (acl *AclApp) get_acl_binding_info_for_subtree(d *db.DB, intfData *ocbinds.OpenconfigAcl_Acl_Interfaces_Interface, intfId string, direction string) {
    if intfData != nil {
        intfData.Config.Id = intfData.Id
        intfData.State.Id = intfData.Id
    }
    if direction == "INGRESS" {
        if intfData.IngressAclSets != nil && len(intfData.IngressAclSets.IngressAclSet) > 0 {
            for ingressAclSetKey,_ := range intfData.IngressAclSets.IngressAclSet {
                aclName := strings.ReplaceAll(strings.ReplaceAll(ingressAclSetKey.SetName, " ", "_"), "-", "_")
                aclType := ingressAclSetKey.Type.ΛMap()["E_OpenconfigAcl_ACL_TYPE"][int64(ingressAclSetKey.Type)].Name
                aclKey := aclName + "_" + aclType

                ingressAclSet := intfData.IngressAclSets.IngressAclSet[ingressAclSetKey]
                if ingressAclSet != nil && ingressAclSet.AclEntries != nil && len(ingressAclSet.AclEntries.AclEntry) > 0 {
                    for seqId,_ := range ingressAclSet.AclEntries.AclEntry {
                        entrySet := ingressAclSet.AclEntries.AclEntry[seqId]
                        convert_internal_to_oc_acl_rule_binding(d, 0, int64(seqId), direction, nil, entrySet)
                    }
                } else {
                    acl.convert_internal_to_oc_acl_binding(d, aclKey, intfId, direction, ingressAclSet)
                }
            }
        } else {
            acl.find_and_get_acl_binding_info_for(d, intfId, direction, intfData)
        }
    } else if direction == "EGRESS" {
        if intfData.EgressAclSets != nil && len(intfData.EgressAclSets.EgressAclSet) > 0 {
            for egressAclSetKey,_ := range intfData.EgressAclSets.EgressAclSet {
                aclName := strings.ReplaceAll(strings.ReplaceAll(egressAclSetKey.SetName, " ", "_"), "-", "_")
                aclType := egressAclSetKey.Type.ΛMap()["E_OpenconfigAcl_ACL_TYPE"][int64(egressAclSetKey.Type)].Name
                aclKey := aclName + "_" + aclType

                egressAclSet := intfData.EgressAclSets.EgressAclSet[egressAclSetKey]
                if egressAclSet != nil && egressAclSet.AclEntries != nil && len(egressAclSet.AclEntries.AclEntry) > 0 {
                    for seqId,_ := range egressAclSet.AclEntries.AclEntry {
                        entrySet := egressAclSet.AclEntries.AclEntry[seqId]
                        convert_internal_to_oc_acl_rule_binding(d, 0, int64(seqId), direction, nil, entrySet)
                    }
                } else {
                    acl.convert_internal_to_oc_acl_binding(d, aclKey, intfId, direction, egressAclSet)
                }
            }
        } else {
            acl.find_and_get_acl_binding_info_for(d, intfId, direction, intfData)
        }
    } else {
        log.Error("Unknown direction")
    }
}

func (acl *AclApp) find_and_get_acl_binding_info_for(d *db.DB, intfId string, direction string, intfData *ocbinds.OpenconfigAcl_Acl_Interfaces_Interface) {
    if len(acl.aclTableMap) == 0 {
        acl.convert_db_acl_to_internal(d, db.Key{})
    }

    for aclName, aclData := range acl.aclTableMap {
        aclIntfs := aclData.GetList("ports")
        aclType := aclData.Get(ACL_TYPE)
        var aclOrigName string
        var aclOrigType ocbinds.E_OpenconfigAcl_ACL_TYPE
        if SONIC_ACL_TYPE_IPV4 == aclType {
            aclOrigName = strings.Replace(aclName, "_"+OPENCONFIG_ACL_TYPE_IPV4, "", 1)
            aclOrigType = ocbinds.OpenconfigAcl_ACL_TYPE_ACL_IPV4
        } else if SONIC_ACL_TYPE_IPV6 == aclType {
            aclOrigName = strings.Replace(aclName, "_"+OPENCONFIG_ACL_TYPE_IPV4, "", 1)
            aclOrigType = ocbinds.OpenconfigAcl_ACL_TYPE_ACL_IPV6
        } else if SONIC_ACL_TYPE_L2 == aclType {
            aclOrigName = strings.Replace(aclName, "_"+OPENCONFIG_ACL_TYPE_L2, "", 1)
            aclOrigType = ocbinds.OpenconfigAcl_ACL_TYPE_ACL_L2
        }

        if contains(aclIntfs, intfId) && direction == aclData.Get("stage") {
            if direction == "INGRESS" {
                fmt.Println("Ingress find_and_get_acl_binding_info_for AclName: " + aclOrigName + " and Interface: " + intfId)
                if intfData.IngressAclSets != nil {
                    aclSetKey := ocbinds.OpenconfigAcl_Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_Key{ SetName: aclOrigName, Type: aclOrigType }
                    ingressAclSet,ok := intfData.IngressAclSets.IngressAclSet[aclSetKey]; if !ok {
                        ingressAclSet,_ = intfData.IngressAclSets.NewIngressAclSet(aclOrigName, aclOrigType)
                        ygot.BuildEmptyTree(ingressAclSet)
                        ingressAclSet.Config = &ocbinds.OpenconfigAcl_Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_Config{SetName: &aclOrigName, Type: aclOrigType}
                        ingressAclSet.State = &ocbinds.OpenconfigAcl_Acl_Interfaces_Interface_IngressAclSets_IngressAclSet_State{SetName: &aclOrigName, Type: aclOrigType}
                    }
                    acl.convert_internal_to_oc_acl_binding(d, aclName, intfId, direction, ingressAclSet)
                }
            } else if direction == "EGRESS" {
                if intfData.EgressAclSets != nil {
                    aclSetKey := ocbinds.OpenconfigAcl_Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_Key{ SetName: aclOrigName, Type: aclOrigType }
                    egressAclSet,ok := intfData.EgressAclSets.EgressAclSet[aclSetKey]; if !ok {
                        egressAclSet,_ = intfData.EgressAclSets.NewEgressAclSet(aclOrigName, aclOrigType)
                        ygot.BuildEmptyTree(egressAclSet)
                        egressAclSet.Config = &ocbinds.OpenconfigAcl_Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_Config{SetName: &aclOrigName, Type: aclOrigType}
                        egressAclSet.State = &ocbinds.OpenconfigAcl_Acl_Interfaces_Interface_EgressAclSets_EgressAclSet_State{SetName: &aclOrigName, Type: aclOrigType}
                    }
                    acl.convert_internal_to_oc_acl_binding(d, aclName, intfId, direction, egressAclSet)
                }
            }
        }
    }

}


/********************   CREATE related    *******************************/
func convert_oc_acls_to_internal(acl *ocbinds.OpenconfigAcl_Acl) map[string]db.Value {
    var aclInfo map[string]db.Value
    if acl != nil {
        aclInfo = make(map[string]db.Value)
        for aclSetKey,_ := range acl.AclSets.AclSet {
            aclSet := acl.AclSets.AclSet[aclSetKey]
            aclName := strings.ReplaceAll(strings.ReplaceAll(aclSetKey.Name, " ", "_"), "-", "_")
            aclType := aclSetKey.Type.ΛMap()["E_OpenconfigAcl_ACL_TYPE"][int64(aclSetKey.Type)].Name
            aclKey := aclName + "_" + aclType
            m := make(map[string]string)
            aclInfo[aclKey] = db.Value{Field: m}

            if aclSet.Config != nil {
                if aclSet.Config.Type == ocbinds.OpenconfigAcl_ACL_TYPE_ACL_IPV4 {
                    aclInfo[aclKey].Field[ACL_TYPE] = SONIC_ACL_TYPE_IPV4
                } else if aclSet.Config.Type == ocbinds.OpenconfigAcl_ACL_TYPE_ACL_IPV6 {
                    aclInfo[aclKey].Field[ACL_TYPE] = SONIC_ACL_TYPE_IPV6
                } else if aclSet.Config.Type == ocbinds.OpenconfigAcl_ACL_TYPE_ACL_L2 {
                    aclInfo[aclKey].Field[ACL_TYPE] = SONIC_ACL_TYPE_L2
                }

                if len(*aclSet.Config.Description) > 0 {
                    aclInfo[aclKey].Field[ACL_DESCRIPTION] = *aclSet.Config.Description
                }
            }
        }
    }

    return aclInfo
}

func convert_oc_acl_rules_to_internal(acl *ocbinds.OpenconfigAcl_Acl) map[string]map[string]db.Value {
    var rulesInfo map[string]map[string]db.Value
    if acl != nil {
        rulesInfo = make(map[string]map[string]db.Value)
        for aclSetKey,_ := range acl.AclSets.AclSet {
            aclSet := acl.AclSets.AclSet[aclSetKey]
            aclName := strings.ReplaceAll(strings.ReplaceAll(aclSetKey.Name, " ", "_"), "-", "_")
            aclType := aclSetKey.Type.ΛMap()["E_OpenconfigAcl_ACL_TYPE"][int64(aclSetKey.Type)].Name
            aclKey := aclName + "_" + aclType
            rulesInfo[aclKey] = make(map[string]db.Value)

            if aclSet.AclEntries != nil {
                for seqId,_ := range aclSet.AclEntries.AclEntry {
                    entrySet := aclSet.AclEntries.AclEntry[seqId]
                    ruleName := "RULE_" + strconv.FormatInt(int64(seqId), 10)
                    m := make(map[string]string)
                    rulesInfo[aclKey][ruleName] = db.Value{ Field:m }
                    convert_oc_to_internal_rule(rulesInfo[aclKey][ruleName], seqId, aclKey, aclSet.Type, entrySet)
                }
            }

            default_deny_rule(rulesInfo[aclKey])
        }
    }

    return rulesInfo
}

func convert_oc_acl_bindings_to_internal(aclData map[string]db.Value, acl *ocbinds.OpenconfigAcl_Acl) bool {
    var ret bool = false
    if acl.Interfaces != nil && len(acl.Interfaces.Interface) > 0 {
        for intfId,_ := range acl.Interfaces.Interface {
            intf := acl.Interfaces.Interface[intfId]
            if intf != nil {
                fmt.Println("Interface Name: " + *intf.Id)
                if intf.IngressAclSets != nil && len(intf.IngressAclSets.IngressAclSet) > 0 {
                    for inAclKey,_ := range intf.IngressAclSets.IngressAclSet {
                        //ingressAclSet := intf.IngressAclSets.IngressAclSet[inAclKey]
                        aclName := inAclKey.SetName + "_" + inAclKey.Type.ΛMap()["E_OpenconfigAcl_ACL_TYPE"][int64(inAclKey.Type)].Name
                        // TODO: Need to handle Subinterface also
                        if intf.InterfaceRef != nil && intf.InterfaceRef.Config.Interface != nil {
                            aclData[aclName].Field["ports@"] = *intf.InterfaceRef.Config.Interface
                        }
                        aclData[aclName].Field["stage"] = "INGRESS"
                        ret = true
                    }
                }

                if intf.EgressAclSets != nil && len(intf.EgressAclSets.EgressAclSet) > 0 {
                    for outAclKey,_ := range intf.EgressAclSets.EgressAclSet {
                        //egressAclSet := intf.EgressAclSets.EgressAclSet[outAclKey]
                        //aclName := strings.ReplaceAll(strings.ReplaceAll(*egressAclSet.SetName, " ", "_"), "-", "_")
                        aclName := outAclKey.SetName + "_" + outAclKey.Type.ΛMap()["E_OpenconfigAcl_ACL_TYPE"][int64(outAclKey.Type)].Name
                        if intf.InterfaceRef != nil && intf.InterfaceRef.Config.Interface != nil {
                            aclData[aclName].Field["ports@"] = *intf.InterfaceRef.Config.Interface
                        }
                        aclData[aclName].Field["stage"] = "EGRESS"
                        ret = true
                    }
                }
            }
        }
    }
    return ret
}

func default_deny_rule(rulesInfo map[string]db.Value) {
    m := make(map[string]string)
    rulesInfo["DEFAULT_RULE"] = db.Value{ Field: m }
    rulesInfo["DEFAULT_RULE"].Field["PRIORITY"] = strconv.FormatInt(int64(MIN_PRIORITY), 10)
    rulesInfo["DEFAULT_RULE"].Field["PACKET_ACTION"] = "DROP"
}

func convert_oc_to_internal_rule(ruleData db.Value, seqId uint32, aclName string, aclType ocbinds.E_OpenconfigAcl_ACL_TYPE, rule *ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry) {
    ruleIndex := seqId
    ruleData.Field["PRIORITY"] = strconv.FormatInt(int64(MAX_PRIORITY - ruleIndex), 10)
    if rule.Config != nil && rule.Config.Description != nil {
        ruleData.Field["RULE_DESCRIPTION"] = *rule.Config.Description
    }

    if ocbinds.OpenconfigAcl_ACL_TYPE_ACL_IPV4 == aclType {
        convert_oc_to_internal_ipv4(ruleData, aclName, ruleIndex, rule)
    } else if ocbinds.OpenconfigAcl_ACL_TYPE_ACL_IPV4 == aclType {
        convert_oc_to_internal_ipv6(ruleData, aclName, ruleIndex, rule)
    } else if ocbinds.OpenconfigAcl_ACL_TYPE_ACL_L2 == aclType {
        convert_oc_to_internal_l2(ruleData, aclName, ruleIndex, rule)
    } /*else if ocbinds.OpenconfigAcl_ACL_TYPE_ACL_MIXED == aclType {
    } */

    convert_oc_to_internal_transport(ruleData, aclName, ruleIndex, rule)
    convert_oc_to_internal_input_interface(ruleData, aclName, ruleIndex, rule)
    convert_oc_to_internal_action(ruleData, aclName, ruleIndex, rule)
}

func convert_oc_to_internal_l2(ruleData db.Value, aclName string, ruleIndex uint32, rule *ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry) {
    if rule.L2 == nil {
        return
    }
    if rule.L2.Config.Ethertype != nil && util.IsTypeStructPtr(reflect.TypeOf(rule.L2.Config.Ethertype)) {
        ethertypeType := reflect.TypeOf(rule.L2.Config.Ethertype).Elem()
        var b bytes.Buffer
        switch ethertypeType {
        case reflect.TypeOf(ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_L2_Config_Ethertype_Union_E_OpenconfigPacketMatchTypes_ETHERTYPE{}):
            v := (rule.L2.Config.Ethertype).(*ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_L2_Config_Ethertype_Union_E_OpenconfigPacketMatchTypes_ETHERTYPE)
            //ruleData["ETHER_TYPE"] = v.E_OpenconfigPacketMatchTypes_ETHERTYPE.ΛMap()["E_OpenconfigPacketMatchTypes_ETHERTYPE"][int64(v.E_OpenconfigPacketMatchTypes_ETHERTYPE)].Name
            fmt.Fprintf(&b, "0x%0.4x", ETHERTYPE_MAP[v.E_OpenconfigPacketMatchTypes_ETHERTYPE])
            ruleData.Field["ETHER_TYPE"] = b.String()
            break
        case reflect.TypeOf(ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_L2_Config_Ethertype_Union_Uint16{}):
            v := (rule.L2.Config.Ethertype).(*ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_L2_Config_Ethertype_Union_Uint16)
            fmt.Fprintf(&b, "0x%0.4x", v.Uint16)
            ruleData.Field["ETHER_TYPE"] = b.String()
            break
        }
    }
}

func convert_oc_to_internal_ipv4(ruleData db.Value, aclName string, ruleIndex uint32, rule *ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry) {
    if rule.Ipv4.Config.Protocol != nil && util.IsTypeStructPtr(reflect.TypeOf(rule.Ipv4.Config.Protocol)) {
        protocolType := reflect.TypeOf(rule.Ipv4.Config.Protocol).Elem()
        switch protocolType {
        case reflect.TypeOf(ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv4_Config_Protocol_Union_E_OpenconfigPacketMatchTypes_IP_PROTOCOL{}):
            v := (rule.Ipv4.Config.Protocol).(*ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv4_Config_Protocol_Union_E_OpenconfigPacketMatchTypes_IP_PROTOCOL)
            //ruleData["IP_PROTOCOL"] = v.E_OpenconfigPacketMatchTypes_IP_PROTOCOL.ΛMap()["E_OpenconfigPacketMatchTypes_IP_PROTOCOL"][int64(v.E_OpenconfigPacketMatchTypes_IP_PROTOCOL)].Name
            ruleData.Field["IP_PROTOCOL"] = strconv.FormatInt(int64(IP_PROTOCOL_MAP[v.E_OpenconfigPacketMatchTypes_IP_PROTOCOL]), 10)
            break
        case reflect.TypeOf(ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv4_Config_Protocol_Union_Uint8{}):
            v := (rule.Ipv4.Config.Protocol).(*ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv4_Config_Protocol_Union_Uint8)
            ruleData.Field["IP_PROTOCOL"] = strconv.FormatInt(int64(v.Uint8), 10)
            break
        }
    }

    if rule.Ipv4.Config.Dscp != nil {
        ruleData.Field["DSCP"] = strconv.FormatInt(int64(*rule.Ipv4.Config.Dscp), 10)
    }

    if rule.Ipv4.Config.SourceAddress != nil {
        ruleData.Field["SRC_IP"] = *rule.Ipv4.Config.SourceAddress
    }

    if rule.Ipv4.Config.DestinationAddress != nil {
        ruleData.Field["DST_IP"] = *rule.Ipv4.Config.DestinationAddress
    }
}

func convert_oc_to_internal_ipv6(ruleData db.Value, aclName string, ruleIndex uint32, rule *ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry) {
    if rule.Ipv6.Config.Protocol != nil && util.IsTypeStructPtr(reflect.TypeOf(rule.Ipv6.Config.Protocol)) {
        protocolType := reflect.TypeOf(rule.Ipv6.Config.Protocol).Elem()
        switch protocolType {
        case reflect.TypeOf(ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv6_Config_Protocol_Union_E_OpenconfigPacketMatchTypes_IP_PROTOCOL{}):
            v := (rule.Ipv6.Config.Protocol).(*ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv6_Config_Protocol_Union_E_OpenconfigPacketMatchTypes_IP_PROTOCOL)
            //ruleData["IP_PROTOCOL"] = v.E_OpenconfigPacketMatchTypes_IP_PROTOCOL.ΛMap()["E_OpenconfigPacketMatchTypes_IP_PROTOCOL"][int64(v.E_OpenconfigPacketMatchTypes_IP_PROTOCOL)].Name
            ruleData.Field["IP_PROTOCOL"] = strconv.FormatInt(int64(IP_PROTOCOL_MAP[v.E_OpenconfigPacketMatchTypes_IP_PROTOCOL]), 10)
            break
        case reflect.TypeOf(ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv6_Config_Protocol_Union_Uint8{}):
            v := (rule.Ipv6.Config.Protocol).(*ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv6_Config_Protocol_Union_Uint8)
            ruleData.Field["IP_PROTOCOL"] = strconv.FormatInt(int64(v.Uint8), 10)
            break
        }
    }

    if rule.Ipv6.Config.Dscp != nil {
        ruleData.Field["DSCP"] = strconv.FormatInt(int64(*rule.Ipv6.Config.Dscp), 10)
    }

    if rule.Ipv6.Config.SourceAddress != nil {
        ruleData.Field["SRC_IPV6"] = *rule.Ipv6.Config.SourceAddress
    }

    if rule.Ipv6.Config.DestinationAddress != nil {
        ruleData.Field["DST_IPV6"] = *rule.Ipv6.Config.DestinationAddress
    }

    if rule.Ipv6.Config.SourceFlowLabel != nil {
        ruleData.Field["SRC_FLOWLABEL"] = strconv.FormatInt(int64(*rule.Ipv6.Config.SourceFlowLabel), 10)
    }

    if rule.Ipv6.Config.DestinationFlowLabel != nil {
        ruleData.Field["DST_FLOWLABEL"] = strconv.FormatInt(int64(*rule.Ipv6.Config.DestinationFlowLabel), 10)
    }
}

func convert_oc_to_internal_transport(ruleData db.Value, aclName string, ruleIndex uint32, rule *ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry) {
    if rule.Transport == nil {
        return
    }
    if rule.Transport.Config.SourcePort != nil && util.IsTypeStructPtr(reflect.TypeOf(rule.Transport.Config.SourcePort)) {
        sourceportType := reflect.TypeOf(rule.Transport.Config.SourcePort).Elem()
        switch sourceportType {
        case reflect.TypeOf(ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_SourcePort_Union_E_OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_SourcePort{}):
            v := (rule.Transport.Config.SourcePort).(*ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_SourcePort_Union_E_OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_SourcePort)
            ruleData.Field["L4_SRC_PORT"] = v.E_OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_SourcePort.ΛMap()["E_OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_SourcePort"][int64(v.E_OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_SourcePort)].Name
            break
        case reflect.TypeOf(ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_SourcePort_Union_String{}):
            v := (rule.Transport.Config.SourcePort).(*ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_SourcePort_Union_String)
            ruleData.Field["L4_SRC_PORT_RANGE"] = strings.Replace(v.String, "..", "-", 1)
            break
        case reflect.TypeOf(ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_SourcePort_Union_Uint16{}):
            v := (rule.Transport.Config.SourcePort).(*ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_SourcePort_Union_Uint16)
            ruleData.Field["L4_SRC_PORT"] = strconv.FormatInt(int64(v.Uint16), 10)
            break
        }
    }

    if rule.Transport.Config.DestinationPort != nil && util.IsTypeStructPtr(reflect.TypeOf(rule.Transport.Config.DestinationPort)) {
        destportType := reflect.TypeOf(rule.Transport.Config.DestinationPort).Elem()
        switch destportType {
        case reflect.TypeOf(ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_DestinationPort_Union_E_OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_DestinationPort{}):
            v := (rule.Transport.Config.DestinationPort).(*ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_DestinationPort_Union_E_OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_DestinationPort)
            ruleData.Field["L4_DST_PORT"] = v.E_OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_DestinationPort.ΛMap()["E_OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_DestinationPort"][int64(v.E_OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_DestinationPort)].Name
            break
        case reflect.TypeOf(ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_DestinationPort_Union_String{}):
            v := (rule.Transport.Config.DestinationPort).(*ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_DestinationPort_Union_String)
            ruleData.Field["L4_DST_PORT_RANGE"] = strings.Replace(v.String, "..", "-", 1)
            break
        case reflect.TypeOf(ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_DestinationPort_Union_Uint16{}):
            v := (rule.Transport.Config.DestinationPort).(*ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_DestinationPort_Union_Uint16)
            ruleData.Field["L4_DST_PORT"] = strconv.FormatInt(int64(v.Uint16), 10)
            break
        }
    }

    var tcpFlags uint32 = 0x00
    if len(rule.Transport.Config.TcpFlags) > 0 {
        for _,flag := range rule.Transport.Config.TcpFlags {
            fmt.Println("TCP Flag name: " + flag.ΛMap()["E_OpenconfigPacketMatchTypes_TCP_FLAGS"][int64(flag)].Name)
            switch flag {
            case ocbinds.OpenconfigPacketMatchTypes_TCP_FLAGS_TCP_FIN:
                tcpFlags |= 0x01
                break
            case ocbinds.OpenconfigPacketMatchTypes_TCP_FLAGS_TCP_SYN:
                tcpFlags |= 0x02
                break
            case ocbinds.OpenconfigPacketMatchTypes_TCP_FLAGS_TCP_RST:
                tcpFlags |= 0x04
                break
            case ocbinds.OpenconfigPacketMatchTypes_TCP_FLAGS_TCP_PSH:
                tcpFlags |= 0x08
                break
            case ocbinds.OpenconfigPacketMatchTypes_TCP_FLAGS_TCP_ACK:
                tcpFlags |= 0x10
                break
            case ocbinds.OpenconfigPacketMatchTypes_TCP_FLAGS_TCP_URG:
                tcpFlags |= 0x20
                break
            case ocbinds.OpenconfigPacketMatchTypes_TCP_FLAGS_TCP_ECE:
                tcpFlags |= 0x40
                break
            case ocbinds.OpenconfigPacketMatchTypes_TCP_FLAGS_TCP_CWR:
                tcpFlags |= 0x80
                break
            }
        }
        var b bytes.Buffer
        fmt.Fprintf(&b, "0x%0.2x/0x%0.2x", tcpFlags, tcpFlags)
        ruleData.Field["TCP_FLAGS"] = b.String()
    }
}

func convert_oc_to_internal_input_interface(ruleData db.Value, aclName string, ruleIndex uint32, rule *ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry) {
    if rule.InputInterface != nil && rule.InputInterface.InterfaceRef != nil {
        ruleData.Field["IN_PORTS"] = *rule.InputInterface.InterfaceRef.Config.Interface
    }
}

func convert_oc_to_internal_action(ruleData db.Value, aclName string, ruleIndex uint32, rule *ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry) {
    if rule.Actions != nil && rule.Actions.Config != nil {
        switch rule.Actions.Config.ForwardingAction {
        case ocbinds.OpenconfigAcl_FORWARDING_ACTION_ACCEPT:
            ruleData.Field["PACKET_ACTION"] = "FORWARD"
            break
        case ocbinds.OpenconfigAcl_FORWARDING_ACTION_DROP:
            ruleData.Field["PACKET_ACTION"] = "DROP"
            break
        case ocbinds.OpenconfigAcl_FORWARDING_ACTION_REJECT:
            ruleData.Field["PACKET_ACTION"] = "DROP"
            break
        default:
        }
    }
}

func set_acl_data_in_config_db(d *db.DB, aclData map[string]db.Value) {
    for key := range aclData {

        /*
        existingEntry,_ := dbCl.GetEntry(&aclTs, db.Key{Comp: []string {key} })
        //Merge any ACL binds already present. Validate should take care of any checks so its safe to blindly merge here
        if len(existingEntry.Field) > 0  {
            value.Field["ports"] += "," + existingEntry.Field["ports@"]
        }
        fmt.Println(value)
        */
        err := d.SetEntry(&aclTs, db.Key{Comp: []string {key} }, aclData[key])
        if err != nil {
            fmt.Println(err)
        }
    }
}

func set_acl_rule_data_in_config_db(d *db.DB, ruleData map[string]map[string]db.Value) {
    for aclName := range ruleData {
        for ruleName := range ruleData[aclName] {
            err := d.SetEntry(&ruleTs, db.Key{Comp: []string {aclName, ruleName} }, ruleData[aclName][ruleName])
            if err != nil {
                fmt.Println(err)
            }
        }
    }
}

/*
func set_acl_bind_data_in_config_db(dbCl *db.DB, aclData map[string]map[string]map) {
}
*/

func getIpProtocol(proto int64, aclType ocbinds.E_OpenconfigAcl_ACL_TYPE, contType string) interface{} {
    foundInMap := false
    var ptype ocbinds.E_OpenconfigPacketMatchTypes_IP_PROTOCOL = ocbinds.OpenconfigPacketMatchTypes_IP_PROTOCOL_UNSET

    for k, v := range IP_PROTOCOL_MAP {
        if proto == int64(v) {
            foundInMap = true
            ptype = k
        }
    }

    switch aclType {
        case ocbinds.OpenconfigAcl_ACL_TYPE_ACL_IPV4:
            if "config" == contType {
                if foundInMap {
                    var ipProCfg *ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv4_Config_Protocol_Union_E_OpenconfigPacketMatchTypes_IP_PROTOCOL
                    ipProCfg = new (ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv4_Config_Protocol_Union_E_OpenconfigPacketMatchTypes_IP_PROTOCOL)
                    ipProCfg.E_OpenconfigPacketMatchTypes_IP_PROTOCOL = ptype
                    return ipProCfg
                } else {
                    var ipProCfgUint8 *ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv4_Config_Protocol_Union_Uint8
                    ipProCfgUint8 = new (ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv4_Config_Protocol_Union_Uint8)
                    ipProCfgUint8.Uint8 = uint8(proto)
                    return ipProCfgUint8
                }
            } else if "state" == contType {
                if foundInMap {
                    var ipProSt *ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv4_State_Protocol_Union_E_OpenconfigPacketMatchTypes_IP_PROTOCOL
                    ipProSt = new (ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv4_State_Protocol_Union_E_OpenconfigPacketMatchTypes_IP_PROTOCOL)
                    ipProSt.E_OpenconfigPacketMatchTypes_IP_PROTOCOL = ptype
                    return ipProSt
                } else {
                    var ipProStUint8 *ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv4_State_Protocol_Union_Uint8
                    ipProStUint8 = new (ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv4_State_Protocol_Union_Uint8)
                    ipProStUint8.Uint8 = uint8(proto)
                    return ipProStUint8
                }
            }
            break
        case ocbinds.OpenconfigAcl_ACL_TYPE_ACL_IPV6:
            if "config" == contType {
                if foundInMap {
                    var ipv6ProCfg *ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv6_Config_Protocol_Union_E_OpenconfigPacketMatchTypes_IP_PROTOCOL
                    ipv6ProCfg = new (ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv6_Config_Protocol_Union_E_OpenconfigPacketMatchTypes_IP_PROTOCOL)
                    ipv6ProCfg.E_OpenconfigPacketMatchTypes_IP_PROTOCOL = ptype
                    return ipv6ProCfg
                } else {
                    var ipv6ProCfgUint8 *ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv6_Config_Protocol_Union_Uint8
                    ipv6ProCfgUint8 = new (ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv6_Config_Protocol_Union_Uint8)
                    ipv6ProCfgUint8.Uint8 = uint8(proto)
                    return ipv6ProCfgUint8
                }
            } else if "state" == contType {
                if foundInMap {
                    var ipv6ProSt *ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv6_State_Protocol_Union_E_OpenconfigPacketMatchTypes_IP_PROTOCOL
                    ipv6ProSt = new (ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv6_State_Protocol_Union_E_OpenconfigPacketMatchTypes_IP_PROTOCOL)
                    ipv6ProSt.E_OpenconfigPacketMatchTypes_IP_PROTOCOL = ptype
                    return ipv6ProSt
                } else {
                    var ipv6ProStUint8 *ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv6_State_Protocol_Union_Uint8
                    ipv6ProStUint8 = new (ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Ipv6_State_Protocol_Union_Uint8)
                    ipv6ProStUint8.Uint8 = uint8(proto)
                    return ipv6ProStUint8
                }
            }
            break
    }
    return nil
}

func getTransportConfigDestPort(destPort string) ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_DestinationPort_Union {
    portNum, _ := strconv.ParseInt(destPort, 10, 64)
    var destPortCfg *ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_DestinationPort_Union_Uint16
    destPortCfg = new (ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_DestinationPort_Union_Uint16)
    destPortCfg.Uint16 = uint16(portNum)
    return destPortCfg
}

func getTransportConfigSrcPort(srcPort string) ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_SourcePort_Union {
    portNum, _ := strconv.ParseInt(srcPort, 10, 64)
    var srcPortCfg *ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_SourcePort_Union_Uint16
    srcPortCfg = new (ocbinds.OpenconfigAcl_Acl_AclSets_AclSet_AclEntries_AclEntry_Transport_Config_SourcePort_Union_Uint16)
    srcPortCfg.Uint16 = uint16(portNum)
    return srcPortCfg
}


func getYangPathFromUri(uri string) (string, error) {
    var path *gnmi.Path
    var err error

    path, err = ygot.StringToPath(uri, ygot.StructuredPath, ygot.StringSlicePath)
    if err != nil {
        fmt.Println("Error in uri to path conversion: ", err)
        return "", err
    }

    yangPath, yperr := ygot.PathToSchemaPath(path)
    if yperr != nil {
        fmt.Println("Error in Gnmi path to Yang path conversion: ", yperr)
        return "", yperr
    }

    return yangPath, err
}

func getYangPathFromStruct(s ygot.GoStruct) string {
    tn := reflect.TypeOf(s).Elem().Name()
    schema, ok := ocbinds.SchemaTree[tn]
    if !ok {
        fmt.Errorf("could not find schema for type %s", tn )
        return ""
    } else if schema != nil {
        yPath := schema.Path()
        yPath = strings.Replace(yPath, "/device/acl", "/openconfig-acl:acl", 1)
        return yPath
    }
    return ""
}

/* Check if targetUriPath is child (subtree) of nodePath
The return value can be used to decide if subtrees needs
to visited to fill the data or not.
*/
func isSubtreeRequest(targetUriPath string, nodePath string) bool {
    return strings.HasPrefix(targetUriPath, nodePath)
}

func dumpIetfJson(s ygot.ValidatedGoStruct) ([]byte, error) {
    jsonStr, err := ygot.EmitJSON(s, &ygot.EmitJSONConfig{
        Format: ygot.RFC7951,
        Indent: "  ",
        RFC7951Config: &ygot.RFC7951JSONConfig{
            AppendModuleName: true,
        },
    })
    return []byte(jsonStr), err
}

func contains(sl []string, str string) bool {
    for _,v := range sl {
        if v == str {
            return true
        }
    }
    return false
}

// remove this test method later
func (acl *AclApp) testGetParentNode() {
    deviceObj := (*acl.ygotRoot).(*ocbinds.Device)
    // parentObjIntf, _, _ := getParentNode(acl.path, deviceObj)
    uri := "/acl/acl-sets/acl-set"
    parentObjIntf, _, _ := getParentNode(&uri, deviceObj)
    fmt.Println("aclSetsObj =>", reflect.TypeOf(*parentObjIntf).Elem().Name())
    aclSetsObj, ok := (*parentObjIntf).(*ocbinds.OpenconfigAcl_Acl_AclSets)
    if ok == true {
        fmt.Println("OpenconfigAcl_Acl_AclSets object casting is success =>", aclSetsObj)
    } else {
        fmt.Println("testGetParentNode: object casting fails")
    }
}

