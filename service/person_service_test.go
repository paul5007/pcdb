package service

import (
	"testing"

	"github.com/paul5007/pcdb/database"
	"github.com/paul5007/pcdb/domain"
)

func startDB() *database.PCDB {
	db := database.NewPCDB()
	db.AddRegion("Person")
	return db
}

func TestAddPerson(t *testing.T) {
	db := startDB()
	ps := PersonService{}
	p := domain.Person{ID: "1", Name: "Paul", Age: 25}
	err := ps.AddPerson(db, p)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestAddPersonNoIDProvided(t *testing.T) {
	db := startDB()
	ps := PersonService{}
	p := domain.Person{Name: "Paul", Age: 25}
	err := ps.AddPerson(db, p)
	if err != nil {
		t.Log(err)
	}
	list, _ := ps.ListPersons(db)
	if 1 != len(list) {
		t.Logf("Invalid number of Persons in database. Should be 1 but is %v", list)
		t.Fail()
	}
}

func TestGetPerson(t *testing.T) {
	db := startDB()
	ps := PersonService{}
	p := domain.Person{ID: "1", Name: "Paul", Age: 25}
	err := ps.AddPerson(db, p)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	per, err := ps.GetPerson(db, "1")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if *per != p {
		t.Logf("Person: %v does not equal Person: %v", p, per)
		t.FailNow()
	}
}
