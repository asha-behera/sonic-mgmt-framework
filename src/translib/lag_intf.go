package translib

import (
	"errors"
	log "github.com/golang/glog"
	"translib/db"
	"translib/tlerr"
)

/******** CONFIG FUNCTIONS ********/

func (app *IntfApp) translateUpdateLagIntf(d *db.DB, lagName *string, inpOp reqType) ([]db.WatchKeys, error) {
	log.Info("----INSIDE LAG INTF FILE-translate----")
	var err error
	var keys []db.WatchKeys

	intfObj := app.getAppRootObject()

	m := make(map[string]string)
	entryVal := db.Value{Field: m}
	entryVal.Field["admin_status"]= "up"
	entryVal.Field["mtu"]= "9100"
	if err != nil {
		return keys, err
	}

	lag := intfObj.Interface[*lagName]
	curr, _ := d.GetEntry(app.lagD.lagTs, db.Key{Comp: []string{*lagName}})
	if !curr.IsPopulated() {
		log.Info("LAG-" + *lagName + " not present in DB, need to create it!!")
		app.ifTableMap[*lagName] = dbEntry{op: opCreate, entry: entryVal}
		return keys, nil
	}
	log.Info("----> LAGkeys are ", keys)
	app.translateUpdateIntfConfig(lagName, lag, &curr)
	return keys, err
}

func (app *IntfApp) processUpdateLagIntfConfig(d *db.DB) error {
	log.Info("INSIDE LAG INTF FILE--processUPDATE---")
	var err error

	for lagName, lagEntry := range app.ifTableMap {
	log.Info("--->lagName and lagentry are", lagName, lagEntry) //Portchannel10{1 {map[admin_status:up mtu:9100]}}
		switch lagEntry.op {
		case opCreate:
			err = d.CreateEntry(app.lagD.lagTs, db.Key{Comp: []string{lagName}}, lagEntry.entry)
			if err != nil {
				errStr := "Creating LAG entry for LAG : " + lagName + " failed"
				return errors.New(errStr)
			}
		case opUpdate:
			err = d.SetEntry(app.lagD.lagTs, db.Key{Comp: []string{lagName}}, lagEntry.entry)
			if err != nil {
				errStr := "Updating LAG entry for LAG : " + lagName + " failed"
				return errors.New(errStr)
			}
		}
	}
	return err
}

func (app *IntfApp) processUpdateLagIntf(d *db.DB) error {
	var err error
	err = app.processUpdateLagIntfConfig(d)
	if err != nil {
		return err
	}

	return err
}

/********* DELETE FUNCTIONS ********/

func (app *IntfApp) translateDeleteLagIntf(d *db.DB, lagName string) ([]db.WatchKeys, error) {
	var err error
	var keys []db.WatchKeys
	curr, err := d.GetEntry(app.lagD.lagTs, db.Key{Comp: []string{lagName}})
	if err != nil {
		errStr := "Invalid Lag: " + lagName
		return keys, tlerr.InvalidArgsError{Format: errStr}
	}
	app.ifTableMap[lagName] = dbEntry{entry: curr, op: opDelete}
	return keys, err
}

func (app *IntfApp) processDeleteLagIntfAndMembers(d *db.DB) error {
	var err error

	for lagKey, dbentry := range app.ifTableMap {
		memberPortsVal, ok := dbentry.entry.Field["members@"]
		if ok {
			memberPorts := generateMemberPortsSliceFromString(&memberPortsVal)
			if memberPorts == nil {
				return nil
			}
			log.Info("MemberPorts = ", memberPortsVal)

			for _, memberPort := range memberPorts {
				log.Infof("Member Port:%s part of lag:%s to be deleted!", memberPort, lagKey)
				err = d.DeleteEntry(app.lagD.lagMemberTs, db.Key{Comp: []string{lagKey, memberPort}})
				if err != nil {
					return err
				}
			}
		}
		err = d.DeleteEntry(app.lagD.lagTs, db.Key{Comp: []string{lagKey}})
		if err != nil {
			return err
		}
	}
	return err
}

func (app *IntfApp) processDeleteLagIntf(d *db.DB) error {
	var err error

	err = app.processDeleteLagIntfAndMembers(d)
	if err != nil {
		return err
	}
	return err
}
