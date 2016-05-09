package db

import (
	"github.com/kalambet/mission-control/services"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var collectionName = "bmcio"
var bmcioIndexName = "bmcioIndex"

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
func (driver *Driver) GetStatuses(service *services.Service, count int) ([]services.ServiceStatus, error) {
	session, err := mgo.Dial(driver.URI)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	var result = make([]services.ServiceStatus, count)
	collection := session.DB("").C(collectionName)
	collection.Find(bson.M{"name": service.Name}).Sort("-updatetime").Limit(count).All(&result)

	return result, nil
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

	collection := session.DB("").C(collectionName)
	err = collection.Create(&collectionInfo)
	if err != nil {
		return err
	}

	// Now we need to create an index for the collection
	var bmcioIndex = mgo.Index{
		Key:        []string{"name", "updatetime"},
		Background: true,
		Name:       bmcioIndexName,
		Unique:     false}

	err = collection.EnsureIndex(bmcioIndex)
	if err != nil {
		return err
	}

	return nil
}
