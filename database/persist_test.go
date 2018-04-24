package database

import (
	"testing"
)

func TestExportRegion(t *testing.T) {
	db := newMockDB()

	err := db.ExportRegion(testRegionName)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func TestImportRegion(t *testing.T) {
	db := newMockDB()

	err := db.ExportRegion(testRegionName)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	db = nil

	newDB := NewPCDB()
	newDB.AddRegion(testRegionName)

	err = newDB.ImportRegion(testRegionName, "./"+testRegionName+fileSuffix)
	if err != nil {
		t.Logf("Failed to import data into new DB | %v", err)
		t.FailNow()
	}
	r, err := newDB.GetRegion(testRegionName)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if 3 != r.NumEntries() {
		t.Log("Region should have imported 3 entries")
		t.Fail()
	}
}
