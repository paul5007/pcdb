package service

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/paul5007/pcdb/database"
	"github.com/paul5007/pcdb/domain"
)

const (
	regionName = "Person"
)

var (
	idCounter int
)

// PersonService allows manipulation of Persons within the database and service layer
type PersonService struct{}

// AddPerson adds a new person to the database
func (ps *PersonService) AddPerson(db database.PaulCacheDatabase, p domain.Person) error {
	r, err := db.GetRegion(regionName)
	if err != nil {
		return fmt.Errorf("Failed to find region: %v | %v", regionName, err)
	}
	// if ID for Person is not provided then generate one
	if p.ID == "" {
		p.ID = generateID()
	}
	err = r.Add(p.ID, p)
	if err != nil {
		return fmt.Errorf("Failed to add Person: %v to region: %v | %v", p, regionName, err)
	}
	return nil
}

// RemovePerson removes a person from the database
func (ps *PersonService) RemovePerson(db database.PaulCacheDatabase, id string) error {
	r, err := db.GetRegion(regionName)
	if err != nil {
		return fmt.Errorf("Failed to find region: %v | %v", regionName, err)
	}
	err = r.Remove(id)
	if err != nil {
		return fmt.Errorf("Failed to remove Person with ID: %v from region: %v | %v", id, regionName, err)
	}
	return nil
}

// GetPerson returns a person from the database
func (ps *PersonService) GetPerson(db database.PaulCacheDatabase, id string) (*domain.Person, error) {
	r, err := db.GetRegion(regionName)
	if err != nil {
		return nil, fmt.Errorf("Failed to find region: %v | %v", regionName, err)
	}
	data, err := r.Get(id)
	if err != nil {
		return nil, fmt.Errorf("Failed to get Person with ID: %v from region: %v | %v", id, regionName, err)
	}
	if p, ok := data.(domain.Person); ok {
		return &p, nil
	}
	return nil, fmt.Errorf("Data returned was not of domain Person. Returned %v", data)
}

// ListPersons will return all persons from database
func (ps *PersonService) ListPersons(db database.PaulCacheDatabase) ([]domain.Person, error) {
	r, err := db.GetRegion(regionName)
	if err != nil {
		return nil, fmt.Errorf("Failed to find region: %v | %v", regionName, err)
	}
	list := r.GetAll()
	var personList []domain.Person
	for _, p := range list {
		person, ok := p.(domain.Person)
		if ok {
			personList = append(personList, person)
		}
	}
	return personList, nil
}

func generateID() string {
	time := time.Now()
	rand.Seed(time.Unix())

	rnd := rand.Intn(1000000)

	id := "d" + strconv.FormatInt(int64(rnd), 10)
	return id
}
