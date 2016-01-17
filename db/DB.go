package db

import (
	"github.com/kalambet/mission-control/services"
	"gopkg.in/mgo.v2"
)

var collectionName = "bmcio"

// Driver is an abstrcation object that connects backend to db
type Driver struct {
	URI string
	//Connection
}

// SaveStatus saves service status to dbtabase
func (driver *Driver) SaveStatus(state *services.ServiceStatus) error {
	session, err := mgo.Dial(driver.URI)
	if err != nil {
		return err
	}
	defer session.Close()

	colleciton := session.DB("").C(collectionName)
	err = colleciton.Insert(state)
	if err != nil {
		return err
	}

	return nil
}

// GetStatuses get last `count` serviced status collected entries
func (driver *Driver) GetStatuses(service *services.Service, count int) []*services.ServiceStatus {

	return nil
}

// InitDatabase initialize database in case it was not initialized already
func (driver *Driver) InitDatabase() error {
	session, err := mgo.Dial(driver.URI)
	if err != nil {
		return err
	}
	defer session.Close()

	var collectionInfo = mgo.CollectionInfo{
		Capped:  true,
		MaxDocs: 25000}

	err = session.DB("").C(collectionName).Create(&collectionInfo)
	if err != nil {
		return err
	}

	return nil
}
