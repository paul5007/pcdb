package database

import (
	"testing"
)

func TestNewPCDB(t *testing.T) {
	d := NewPCDB()
	if d == nil {
		t.Fail()
	}
}

func TestNumRegions(t *testing.T) {
	d := NewPCDB()
	if d == nil {
		t.Log("Database not initialized")
		t.FailNow()
	}
	if 0 != d.NumRegions() {
		t.Fail()
	}
	if len(d.RegionList) != d.NumRegions() {
		t.Fail()
	}
}

func TestAddRegion(t *testing.T) {
	d := NewPCDB()
	err := d.AddRegion("TestAdd")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if 1 != d.NumRegions() {
		t.Log("Incorrect number of regions:", d.NumRegions())
		t.Log(d.RegionList)
		t.Fail()
	}
}

func TestTwoAddRegionsSameName(t *testing.T) {
	d := NewPCDB()
	err := d.AddRegion("TestAdd")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if 1 != d.NumRegions() {
		t.Log("Incorrect number of regions:", d.NumRegions())
		t.Log(d.RegionList)
		t.Fail()
	}
	err = d.AddRegion("TestAdd")
	if err == nil {
		t.Log("Should fail to create a region with same name")
		t.Fail()
	}
	if 1 != d.NumRegions() {
		t.Log("Incorrect number of regions:", d.NumRegions())
		t.Log(d.RegionList)
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
