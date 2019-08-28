package translib

import (
	"errors"
	"fmt"
	log "github.com/golang/glog"
	"net"
	"regexp"
	"strconv"
	"strings"
	"translib/db"
	"translib/tlerr"
)

/* Extract Interface type from Interface name */
func (app *IntfApp) getIntfTypeFromIntf(ifName *string) error {
	var err error

	if len(*ifName) == 0 {
		return errors.New("Interface name received is empty! Fetching if-type from interface failed!")
	}
	if strings.HasPrefix(*ifName, "Ethernet") {
		app.intfType = ETHERNET
	} else if strings.HasPrefix(*ifName, "Vlan") {
		app.intfType = VLAN
	} else if strings.HasPrefix(*ifName, "portchannel") {
		app.intfType = LAG
	} else {
		return errors.New("Fetching Interface type from Interface name failed!")
	}
	return err
}

/* Validates whether the specific IP exists in the DB for an Interface*/
/* TODO: Change the name, it does updating DS as well */
func (app *IntfApp) validateIp(dbCl *db.DB, ifName string, ip string) error {
	app.allIpKeys, _ = app.doGetAllIpKeys(dbCl, app.intfD.intfIPTs)

	for _, key := range app.allIpKeys {
		if len(key.Comp) < 2 {
			continue
		}
		if key.Get(0) != ifName {
			continue
		}
		ipAddr, _, _ := net.ParseCIDR(key.Get(1))
		ipStr := ipAddr.String()
		if ipStr == ip {
			log.Infof("IP address %s exists, updating the DS for deletion!", ipStr)
			ipInfo, err := dbCl.GetEntry(app.intfD.intfIPTs, key)
			if err != nil {
				log.Error("Error found on fetching Interface IP info from App DB for Interface Name : ", ifName)
				return err
			}
			if len(app.intfD.ifIPTableMap[key.Get(0)]) == 0 {
				app.intfD.ifIPTableMap[key.Get(0)] = make(map[string]dbEntry)
				app.intfD.ifIPTableMap[key.Get(0)][key.Get(1)] = dbEntry{entry: ipInfo}
			} else {
				app.intfD.ifIPTableMap[key.Get(0)][key.Get(1)] = dbEntry{entry: ipInfo}
			}
			return nil
		}
	}
	return errors.New(fmt.Sprintf("IP address : %s doesn't exist!", ip))
}

/* Validate whether the Interface has IP configuration */
func (app *IntfApp) validateIpExistsForInterface(dbCl *db.DB, ifName *string) bool {
	app.allIpKeys, _ = app.doGetAllIpKeys(dbCl, app.intfD.intfIPTs)

	for _, key := range app.allIpKeys {
		if len(key.Comp) < 2 {
			continue
		}
		if key.Get(0) == *ifName {
			return false
		}
	}
	return true
}

/* Check for IP overlap */
func (app *IntfApp) translateIpv4(d *db.DB, intf string, ip string, prefix int) error {
	var err error
	var ifsKey db.Key

	ifsKey.Comp = []string{intf}

	ipPref := ip + "/" + strconv.Itoa(prefix)
	ifsKey.Comp = []string{intf, ipPref}

	log.Info("ifsKey:=", ifsKey)

	log.Info("Checking for IP overlap ....")
	ipA, ipNetA, _ := net.ParseCIDR(ipPref)

	for _, key := range app.allIpKeys {
		if len(key.Comp) < 2 {
			continue
		}
		ipB, ipNetB, _ := net.ParseCIDR(key.Get(1))

		if ipNetA.Contains(ipB) || ipNetB.Contains(ipA) {
			log.Info("IP ", ipPref, "overlaps with ", key.Get(1), " of ", key.Get(0))

			if intf != key.Get(0) {
				//IP overlap across different interface, reject
				log.Error("IP ", ipPref, " overlaps with ", key.Get(1), " of ", key.Get(0))

				errStr := "IP " + ipPref + " overlaps with IP " + key.Get(1) + " of Interface " + key.Get(0)
				err = tlerr.InvalidArgsError{Format: errStr}
				return err
			} else {
				//IP overlap on same interface, replace
				var entry dbEntry
				entry.op = opDelete

				log.Info("Entry ", key.Get(1), " on ", intf, " needs to be deleted")
				if app.intfD.ifIPTableMap[intf] == nil {
					app.intfD.ifIPTableMap[intf] = make(map[string]dbEntry)
				}
				app.intfD.ifIPTableMap[intf][key.Get(1)] = entry
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
		if app.intfD.ifIPTableMap[intf] == nil {
			app.intfD.ifIPTableMap[intf] = make(map[string]dbEntry)
		}
		app.intfD.ifIPTableMap[intf][ipPref] = entry
	}
	return err
}

/* Validate whether VLAN exists in DB */
func (app *IntfApp) validateVlanExists(d *db.DB, vlanName *string) error {
	if len(*vlanName) == 0 {
		return errors.New("Length of VLAN name is zero")
	}
	entry, err := d.GetEntry(app.vlanD.vlanTs, db.Key{Comp: []string{*vlanName}})
	if err != nil || !entry.IsPopulated() {
		errStr := "Invalid Vlan:" + *vlanName
		return errors.New(errStr)
	}
	return nil
}

/* Validate whether physical interface is valid or not */
func (app *IntfApp) validateInterface(dbCl *db.DB, ifName string, ifKey db.Key) error {
	var err error
	if len(ifName) == 0 {
		return errors.New("Empty Interface name")
	}
	app.intfD.portTblTs = &db.TableSpec{Name: "PORT_TABLE"}
	_, err = dbCl.GetEntry(app.intfD.portTblTs, ifKey)
	if err != nil {
		log.Errorf("Error found on fetching Interface info from App DB for If Name : %s", ifName)
		errStr := "Invalid Interface:" + ifName
		err = tlerr.InvalidArgsError{Format: errStr}
		return err
	}
	return err
}

/* Generate Member Ports string from Slice to update VLAN table in CONFIG DB */
func generateMemberPortsStringFromSlice(memberPortsList []string) (*string, error) {
	if len(memberPortsList) == 0 {
		return nil, nil
	}
	var memberPortsStr strings.Builder

	for _, memberPort := range memberPortsList {
		idx := 1
		if idx != len(memberPortsList) {
			memberPortsStr.WriteString(memberPort + ",")
		} else {
			memberPortsStr.WriteString(memberPort)
		}
		idx = idx + 1
	}
	memberPorts := memberPortsStr.String()
	return &(memberPorts), nil
}

/* Generate list of member-ports from string */
func generateMemberPortsSliceFromString(memberPortsStr *string) ([]string) {
	if len(*memberPortsStr) == 0 {
		return nil
	}
	memberPorts := strings.Split(*memberPortsStr, ",")
	return memberPorts
}

/* Extract VLAN-Id from Vlan String */
func getVlanIdFromVlanName(vlanName *string) (string, error) {
	if !strings.HasPrefix(*vlanName, "Vlan") {
		return "", errors.New("Not valid vlan name : " + *vlanName)
	}
	id := strings.SplitAfter(*vlanName, "Vlan")
	log.Info("Extracted VLAN-Id = ", id[1])
	return id[1], nil
}

/* Validate whether member port exists in the member ports list */
func checkMemberPortExistsInList(memberPortsList []string, memberPort *string) bool {
	for _, port := range memberPortsList {
		if *memberPort == port {
			return true
		}
	}
	return false
}

/* Validate IPv4 address */
func validIPv4(ipAddress string) bool {
	ipAddress = strings.Trim(ipAddress, " ")

	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	if re.MatchString(ipAddress) {
		return true
	}
	return false
}

/* Validate IPv6 address */
func validIPv6(ip6Address string) bool {
	ip6Address = strings.Trim(ip6Address, " ")
	re, _ := regexp.Compile(`(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`)
	if re.MatchString(ip6Address) {
		return true
	}
	return false
}

/* Get all the IP keys from INTERFACE table */
func (app *IntfApp) doGetAllIpKeys(d *db.DB, dbSpec *db.TableSpec) ([]db.Key, error) {

	var keys []db.Key

	intfTable, err := d.GetTable(dbSpec)
	if err != nil {
		return keys, err
	}

	keys, err = intfTable.GetKeys()
	log.Infof("Found %d INTF table keys", len(keys))
	return keys, err
}
