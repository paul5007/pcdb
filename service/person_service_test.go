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
