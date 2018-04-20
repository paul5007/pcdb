package database

import (
	"testing"
)

func TestNewPCDB(t *testing.T) {
	db := NewPCDB()
	if db == nil {
		t.Fail()
	}
}

func TestNumRegions(t *testing.T) {
	db := NewPCDB()
	if db == nil {
		t.Log("Database not initialized")
		t.FailNow()
	}
	if 0 != db.NumRegions() {
		t.Fail()
	}
	if len(db.RegionList) != db.NumRegions() {
		t.Fail()
	}
}

func TestAddRegion(t *testing.T) {
	db := NewPCDB()
	err := db.AddRegion("TestAdd")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if 1 != db.NumRegions() {
		t.Log("Incorrect number of regions:", db.NumRegions())
		t.Log(db.RegionList)
		t.Fail()
	}
}

func TestAddRegionBadName(t *testing.T) {
	db := NewPCDB()
	err := db.AddRegion("TestBadAdd!@#$")
	if err == nil {
		t.Log(err)
		t.FailNow()
	}
	err = db.AddRegion("TestBa!@#dAdd")
	if err == nil {
		t.Log(err)
		t.FailNow()
	}
	if 0 != db.NumRegions() {
		t.Log("Incorrect number of regions:", db.NumRegions())
		t.Log(db.RegionList)
		t.Fail()
	}
}

func TestRemoveRegion(t *testing.T) {
	db := NewPCDB()
	r := "TestDelete"
	err := db.AddRegion(r)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if 1 != db.NumRegions() {
		t.Log("Incorrect number of regions:", db.NumRegions())
		t.Log(db.RegionList)
		t.FailNow()
	}

	err = db.RemoveRegion(r)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if 0 != db.NumRegions() {
		t.Log("Incorrect number of regions:", db.NumRegions())
		t.Log(db.RegionList)
		t.Fail()
	}
}

func TestRemoveRegionBadName(t *testing.T) {
	db := NewPCDB()

	err := db.RemoveRegion("Test@!#$Delete")
	if err == nil {
		t.Fail()
	}

	err = db.RemoveRegion("!%$^TestDelete!@#")
	if err == nil {
		t.Fail()
	}
}

func TestGetRegion(t *testing.T) {
	db := NewPCDB()
	err := db.AddRegion("TestGet")
	if err != nil {
		t.Fail()
	}

	r, err := db.GetRegion("TestGet")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if r.Name != "TestGet" {
		t.Fail()
	}
}
