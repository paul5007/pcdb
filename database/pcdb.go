package database

import (
	"fmt"
	"regexp"
)

// PCDB will act as a stateful list of regions
type PCDB struct {
	RegionList []Region
}

var (
	allowedKeyValues = regexp.MustCompile("[^a-zA-Z$]")
)

// NewPCDB creates a new PCDB with a region initialized
func NewPCDB() *PCDB {
	regionList := []Region{}
	db := &PCDB{RegionList: regionList}
	return db
}

// AddRegion creates a new region in the PCDB
func (db *PCDB) AddRegion(name string) error {
	if allowedKeyValues.MatchString(name) {
		return fmt.Errorf("Illegal characters used in name: %v", name)
	}
	data := make(map[string]interface{})
	r := Region{Data: data, Name: name}
	size := db.NumRegions()
	db.RegionList = append(db.RegionList, r)
	if len(db.RegionList) <= size {
		return fmt.Errorf("Failed to create new region named: %v", name)
	}
	return nil
}

// RemoveRegion removes a region from PCDB
func (db *PCDB) RemoveRegion(name string) error {
	if allowedKeyValues.MatchString(name) {
		return fmt.Errorf("Illegal characters used in name: %v", name)
	}
	regions := db.RegionList
	size := db.NumRegions()
	for i, r := range regions {
		if r.Name == name {
			db.RegionList = append(regions[:i], regions[i+1:]...)
			break
		}
	}
	if len(db.RegionList) >= size {
		return fmt.Errorf("Failed to remove region named: %v", name)
	}
	return nil
}

// GetRegion returns a pointer to a Region object
func (db *PCDB) GetRegion(name string) (*Region, error) {
	if allowedKeyValues.MatchString(name) {
		return nil, fmt.Errorf("Illegal characters used in name: %v", name)
	}
	regions := db.RegionList
	for i, r := range regions {
		if r.Name == name {
			return &db.RegionList[i], nil
		}
	}
	return nil, fmt.Errorf("Failed to find region named: %v", name)
}

// NumRegions returns the number of regions in the mockDB
func (db *PCDB) NumRegions() int {
	return len(db.RegionList)
}
