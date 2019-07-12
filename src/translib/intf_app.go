package translib

import (
    "fmt"
    "net"
    "reflect"
    "strconv"
    "encoding/json"
    "errors"
    "translib/db"
    "translib/ocbinds"
    "github.com/openconfig/ygot/ygot"
    log "github.com/golang/glog"
)

type reqType int

const (
    opCreate     reqType = iota + 1
    opDelete
    opUpdate
)

type dbEntry struct {
    op         reqType
    entry      db.Value
}

type IntfApp struct {
    path        *PathInfo
    reqData     []byte
    ygotRoot    *ygot.GoStruct
    ygotTarget  *interface{}

    respJSON interface{}

    dbIfMap     map[string]dbEntry
    dbIpMap     map[string]map[string]dbEntry

    allIpKeys   []db.Key

    confDB      *db.DB
    portTs      *db.TableSpec
    intfTs      *db.TableSpec
}

func init() {
    log.Info("Init called for INTF module")
    err := register("/openconfig-interfaces:interfaces",
    &appInfo{appType: reflect.TypeOf(IntfApp{}),
    ygotRootType: reflect.TypeOf(ocbinds.OpenconfigInterfaces_Interfaces{}),
    isNative:     false})
    if err != nil {
        log.Fatal("Register INTF app module with App Interface failed with error=", err)
    }

    err = addModel(&ModelData{Name: "openconfig-interfaces",
    Org: "OpenConfig working group",
    Ver:      "1.0.2"})
    if err != nil {
        log.Fatal("Adding model data to appinterface failed with error=", err)
    }
}

func (app *IntfApp) initialize(data appData) {
    log.Info("initialize:if:path =", data.path)

    app.path = NewPathInfo(data.path)
    app.reqData = data.payload
    app.ygotRoot = data.ygotRoot
    app.ygotTarget = data.ygotTarget

    app.portTs = &db.TableSpec{"PORT"}
    app.intfTs = &db.TableSpec{"INTERFACE"}

}

func (app *IntfApp) getAppRootObject() (*ocbinds.OpenconfigInterfaces_Interfaces) {
	deviceObj := (*app.ygotRoot).(*ocbinds.Device)
	return deviceObj.Interfaces
}

func (app *IntfApp) translateCreate(d *db.DB) ([]db.WatchKeys, error)  {
    var err error
    var keys []db.WatchKeys
    log.Info("translateCreate:intf:path =", app.path)

    err = errors.New("Not implemented")
    return keys, err
}

func (app *IntfApp) translateUpdate(d *db.DB) ([]db.WatchKeys, error)  {
    var err error
    var keys []db.WatchKeys

    log.Info("translateUpdate:intf:path =", app.path)

    keys, err = app.translateCommon(d, opUpdate)

    if err != nil {
        log.Info("Something wrong:=", err)
    }

    return keys, err
}

func (app *IntfApp) translateReplace(d *db.DB) ([]db.WatchKeys, error)  {
    var err error
    var keys []db.WatchKeys
    log.Info("translateReplace:intf:path =", app.path)
    err = errors.New("Not implemented")
    return keys, err
}

func (app *IntfApp) translateDelete(d *db.DB) ([]db.WatchKeys, error)  {
    var err error
    var keys []db.WatchKeys
    log.Info("translateDelete:intf:path =", app.path)

    err = errors.New("Not implemented")
    return keys, err
}

func (app *IntfApp) translateGet(dbs [db.MaxDB]*db.DB) error  {
    var err error
    log.Info("translateGet:intf:path =", app.path)
    return err
}

func (app *IntfApp) processCreate(d *db.DB) (SetResponse, error)  {
    var err error
    var resp SetResponse

    log.Info("processCreate:intf:path =", app.path)
    log.Info("ProcessCreate: Target Type is " + reflect.TypeOf(*app.ygotTarget).Elem().Name())

    err = errors.New("Not implemented")
    return resp, err
}

func (app *IntfApp) processUpdate(d *db.DB) (SetResponse, error)  {

    log.Infof("Calling processCommon()")

    resp, err := app.processCommon(d)
    return resp, err
}

func (app *IntfApp) processReplace(d *db.DB) (SetResponse, error)  {
    var err error
    var resp SetResponse
    log.Info("processReplace:intf:path =", app.path)
    err = errors.New("Not implemented")
    return resp, err
}

func (app *IntfApp) processDelete(d *db.DB) (SetResponse, error)  {
    var err error
    var resp SetResponse
    log.Info("processDelete:intf:path =", app.path)

    err = errors.New("Not implemented")
    return resp, err
}

func (app *IntfApp) processGet(dbs [db.MaxDB]*db.DB) (GetResponse, error)  {
    var err error
    var payload []byte
    pathInfo := app.path

    log.Infof("Received GET for path %s; vars=%v", pathInfo.Template, pathInfo.Vars)
    var intfSubtree bool = false

    app.confDB = dbs[db.ConfigDB]
    intfObj := app.getAppRootObject()

    switch pathInfo.Template {
    case "test":

    default:
        err = app.doGetAllPorts()
    }

    var respData []byte
    if err == nil && app.respJSON != nil {
        respData, err = json.Marshal(app.respJSON)
    }

    return GetResponse{Payload: respData}, err

    log.Info("processGet: Target Type: " + reflect.TypeOf(*app.ygotTarget).Elem().Name())
    if reflect.TypeOf(*app.ygotTarget).Elem().Name() == "OpenconfigInterfaces_Interfaces" {
        intfSubtree = true
	log.Info("subtree request = ",  intfSubtree)
    }

    targetUriPath, err := getYangPathFromUri(app.path.Path)
    log.Info("uripath:=", targetUriPath)
    if isSubtreeRequest(targetUriPath, "/openconfig-interfaces:interfaces/interface") {
	if intfObj.Interface != nil  && len(intfObj.Interface) > 0 {
	    log.Info("len:=", len(intfObj.Interface))
	    for ifKey, _ := range intfObj.Interface {
	        log.Info("Name:=", ifKey)
	    }
	}
    }

    err = errors.New("Not implemented")
    return GetResponse{Payload:payload}, err
}

///////////////////////////
func (app *IntfApp) translateCommon(d *db.DB, inpOp reqType) ([]db.WatchKeys, error)  {
    var err error
    var keys []db.WatchKeys
    pathInfo := app.path

    log.Infof("Received UPDATE for path %s; vars=%v", pathInfo.Template, pathInfo.Vars)

    app.allIpKeys,_ = app.doGetAllIpKeys(d)

    app.dbIfMap = make(map[string]dbEntry)
    app.dbIpMap = make(map[string]map[string]dbEntry)

    intfObj := app.getAppRootObject()

    targetUriPath, err := getYangPathFromUri(app.path.Path)
    log.Info("uripath:=", targetUriPath)
    log.Info("err:=", err)

    if intfObj.Interface != nil  && len(intfObj.Interface) > 0 {
        log.Info("len:=", len(intfObj.Interface))
        for ifKey, _ := range intfObj.Interface {
            log.Info("Name:=", ifKey)
            intf := intfObj.Interface[ifKey]
            curr, err := d.GetEntry(app.portTs, db.Key{Comp: []string{ifKey}})
            if err != nil {
                return keys, err
            }
            if !curr.IsPopulated(){
                log.Info("Interface ", ifKey, " doesnt exist in DB")
                err = errors.New("Interface " + ifKey + " doesnt exist in DB")
                return keys, err
            }
            if intf.Config != nil {
                if intf.Config.Description != nil {
                    log.Info("descript:= ", *intf.Config.Description)
                    curr.Field["description"] = *intf.Config.Description
                } else if intf.Config.Mtu != nil {
                    log.Info("mtu:= ", *intf.Config.Mtu)
                    curr.Field["mtu"] = strconv.Itoa(int(*intf.Config.Mtu))
                } else if intf.Config.Enabled != nil {
                    log.Info("enabled:= ", *intf.Config.Enabled)
                    if *intf.Config.Enabled == true {
                        curr.Field["admin_status"] = "up"
                    } else {
                        curr.Field["admin_status"] = "down"
                    }
                }
                log.Info("Writing to db for ", ifKey)
                var entry dbEntry
                entry.op = opUpdate
                entry.entry = curr

                app.dbIfMap[ifKey] = entry
            }
            if intf.Subinterfaces == nil {
                continue
            }
            subIf := intf.Subinterfaces.Subinterface[0]
            if subIf !=  nil {
                if subIf.Ipv4 != nil && subIf.Ipv4.Addresses !=nil {
                    for ip, _ := range subIf.Ipv4.Addresses.Address {
                        addr := subIf.Ipv4.Addresses.Address[ip]
                        if addr.Config != nil {
                            log.Info("Ip:=", *addr.Config.Ip)
                            log.Info("prefix:=", *addr.Config.PrefixLength)
                            err = app.translateIpv4(d, ifKey, *addr.Config.Ip, int(*addr.Config.PrefixLength))
                            if err != nil {
                                return keys, err
                            }
                        }
                    }
                }
                if subIf.Ipv6 != nil && subIf.Ipv6.Addresses !=nil {
                    for ip, _ := range subIf.Ipv6.Addresses.Address {
                        addr := subIf.Ipv6.Addresses.Address[ip]
                        if addr.Config != nil {
                            log.Info("Ip:=", *addr.Config.Ip)
                            log.Info("prefix:=", *addr.Config.PrefixLength)
                            err = app.translateIpv4(d, ifKey, *addr.Config.Ip, int(*addr.Config.PrefixLength))
                            if err != nil {
                                return keys, err
                            }
                        }
                    }
                }
           } else {
               err =  errors.New("Only subinterface index 0 is supported")
               return keys, err
           }
        }
    } else {
	err = errors.New("Not implemented")
    }

    return keys, err
}

func (app *IntfApp) translateIpv4(d *db.DB, intf string, ip string, prefix int) (error)  {
    var err error
    var ifsKey db.Key;

    ifsKey.Comp = []string{intf }

    ipPref := ip + "/" + strconv.Itoa(prefix)
    ifsKey.Comp = []string{intf, ipPref}

    log.Info("ifsKey:=", ifsKey)

    log.Info("Checking for IP overlap ....")
    ipA, ipNetA, _ := net.ParseCIDR(ipPref)

    for _, key := range app.allIpKeys {
        ipB, ipNetB, _ := net.ParseCIDR(key.Get(1))

        if ipNetA.Contains(ipB) || ipNetB.Contains(ipA) {
            log.Info("IP ", ipPref , "overlaps with ", key.Get(1), " of ", key.Get(0))

            if intf != key.Get(0) {
                //IP overlap across different interface, reject
                log.Info("IP ", ipPref , " overlaps with ", key.Get(1), " of ", key.Get(0))
                err = errors.New(fmt.Sprintf("IP %s overlaps with %s of %s ", ipPref, key.Get(1), key.Get(0)))
                return err
            } else {
                //IP overlap on same interface, replace
                var entry dbEntry
                entry.op = opDelete

                log.Info("Entry ", key.Get(1), " on ", intf, " needs to be deleted")
                if app.dbIpMap[intf]  == nil {
                    app.dbIpMap[intf] = make(map[string]dbEntry)
                }
                app.dbIpMap[intf][key.Get(1)] = entry
            }
        }
    }

    //At this point, we need to add the entry to db
    {
        var entry dbEntry
        entry.op = opCreate

        m := make(map[string]string)
        m["NULL"] = "NULL"
        value := db.Value{Field: m}
        entry.entry = value
        if app.dbIpMap[intf]  == nil {
            app.dbIpMap[intf] = make(map[string]dbEntry)
        }
        app.dbIpMap[intf][ipPref] = entry
    }
    return err
}

func (app *IntfApp) processCommon(d *db.DB) (SetResponse, error)  {
    var err error
    var resp SetResponse

    log.Info("processCommon:intf:path =", app.path)
    log.Info("ProcessCommon: Target Type is " + reflect.TypeOf(*app.ygotTarget).Elem().Name())

    for key, entry := range app.dbIfMap {
        if entry.op == opUpdate {
            log.Info("Updating entry for ", key)
            err = d.SetEntry(app.portTs, db.Key{Comp: []string{key}}, entry.entry)
        }
    }

    for key, entry1 := range app.dbIpMap {
        for ip, entry := range entry1 {
            if entry.op == opCreate {
                log.Info("Creating entry for ", key,":", ip)
                err = d.CreateEntry(app.intfTs, db.Key{Comp: []string{key, ip}}, entry.entry)
            } else if entry.op == opDelete {
                log.Info("Deleting entry for ", key,":", ip)
                err = d.DeleteEntry(app.intfTs, db.Key{Comp: []string{key, ip}})
            }
        }
    }
    return resp, err
}

func (app *IntfApp) doGetAllIpKeys(d *db.DB) ([]db.Key, error) {
	log.Infof("in GetAllIpKeys")
        var keys    []db.Key

	// Get all vlans from db
	intfTable, err := d.GetTable(app.intfTs)
	if err != nil {
	    return keys, err
	}

	keys, err = intfTable.GetKeys()
	log.Infof("Found %d INTF table keys", len(keys))
        return keys, err
}

func (app *IntfApp) doGetAllPorts() error {
	log.Infof("in GetAllPorts")

	// Get all vlans from db
	portTable, err := app.confDB.GetTable(app.portTs)
	if err != nil {
		return err
	}

	var allPortsJSON jsonArray

	keys, _ := portTable.GetKeys()
	log.Infof("Found %d PORT table keys", len(keys))
	for _, key := range keys {
		log.Infof("Processing %v", key.Get(0))

		portInfo, _ := portTable.GetEntry(key)
		portJSON, err := app.getPortJSON(key, &portInfo)
		if err != nil {
			return err
		}

		allPortsJSON = append(allPortsJSON, *portJSON)
	}

	app.respJSON = &allPortsJSON
	return nil
}

func (app *IntfApp) getPortJSON(key db.Key, portEntry *db.Value) (*jsonObject, error) {
	portJSON := make(jsonObject)

	log.Infof("Preparing json for port %v", key.Get(0))

	portJSON["name"] = key.Get(0)

	return &portJSON, nil
}
