package service

import (
	"fmt"

	"github.com/paul5007/pcdb/database"
	"github.com/paul5007/pcdb/domain"
)

const (
	regionName = "Person"
)

// PersonService allows manipulation of Persons within the database and service layer
type PersonService struct{}

// AddPerson adds a new person to the database
func (ps *PersonService) AddPerson(db database.PaulCacheDatabase, p domain.Person) error {
	r, err := db.GetRegion(regionName)
	if err != nil {
		return fmt.Errorf("Failed to find region: %v | %v", regionName, err)
	}

	err = r.Add(p.ID, p)
	if err != nil {
		return fmt.Errorf("Failed to add Person: %v to region: %v | %v", p, regionName, err)
	}
	return nil
}

// RemovePerson removes a person from the database
func (ps *PersonService) RemovePerson(db database.PaulCacheDatabase, p domain.Person) error {
	r, err := db.GetRegion(regionName)
	if err != nil {
		return fmt.Errorf("Failed to find region: %v | %v", regionName, err)
	}
	err = r.Remove(p.ID)
	if err != nil {
		return fmt.Errorf("Failed to remove Person: %v from region: %v | %v", p, regionName, err)
	}
	return nil
}
