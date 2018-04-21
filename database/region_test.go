package database

import (
	"reflect"
	"testing"
)

func TestNewRegion(t *testing.T) {
	r := NewRegion("Test")
	if r.Name != "Test" {
		t.Fail()
	}
	if r.Data == nil {
		t.Fail()
	}
}

func TestRegionNumEntries(t *testing.T) {
	r := NewRegion("Test")
	if 0 != r.NumEntries() {
		t.Fail()
	}
	if len(r.Data) != r.NumEntries() {
		t.Fail()
	}
}

func TestRegionAddObject(t *testing.T) {
	key := "key"
	mock := int(10)

	r := NewRegion("Test")
	err := r.Add(key, mock)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if r.Data[key] == nil {
		t.Fail()
	}
}

func TestRegionAddTwoSimilarObjects(t *testing.T) {
	key1, key2 := "key1", "key2"
	mock1 := int(10)
	mock2 := int(25)

	r := NewRegion("Test")
	err := r.Add(key1, mock1)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if r.Data[key1] == nil {
		t.Fail()
	}

	err = r.Add(key2, mock2)
	if err != nil {
		t.FailNow()
	}
	if r.Data[key2] == nil {
		t.Fail()
	}
}

func TestRegionAddTwoSimilarObjectsToSameKey(t *testing.T) {
	// fail if try to add two values to a single key
	key := "key"
	mock1 := int(10)
	mock2 := int(25)

	r := NewRegion("Test")
	err := r.Add(key, mock1)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if r.Data[key] == nil {
		t.Fail()
	}

	err = r.Add(key, mock2)
	if err == nil {
		t.FailNow()
	}
	// value should still be first value added without overwrite
	if r.Data[key] != mock1 {
		t.Fail()
	}
}

func TestRegionAddTwoDifferentObjects(t *testing.T) {
	// fail if there are two different data types in a region
	key1, key2 := "key1", "key2"
	mock1 := int(10)
	mock2 := "25"

	r := NewRegion("Test")
	err := r.Add(key1, mock1)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if r.Data[key1] == nil {
		t.Fail()
	}

	err = r.Add(key2, mock2)
	if err == nil {
		t.Log("Different data type should be reject")
		t.Fail()
	}
	if r.Data[key2] != nil {
		t.Fail()
	}
}

func TestRegionRemoveObject(t *testing.T) {
	key := "key"
	mockObj := int(10)

	r := NewRegion("Test")
	err := r.Add(key, mockObj)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	err = r.Remove(key)
	if err != nil {
		t.FailNow()
	}

	if r.Data == nil {
		t.Fail()
	}

	if r.Data[key] != nil {
		t.Fail()
	}
}

func TestRegionGetObject(t *testing.T) {
	key := "key"
	mockObj := int(10)

	r := NewRegion("Test")
	err := r.Add(key, mockObj)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	val, err := r.Get(key)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if reflect.TypeOf(val).Kind() != reflect.TypeOf(mockObj).Kind() {
		t.Log("Type of stored and retrieved object does not match")
		t.Fail()
	}
	if val != int(10) {
		t.Fail()
	}
}

func TestRegionUpdate(t *testing.T) {
	// fail if try to add two values to a single key
	key := "key"
	mock1 := int(10)
	mock2 := int(25)

	r := NewRegion("Test")
	err := r.Add(key, mock1)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if r.Data[key] == nil {
		t.Fail()
	}

	err = r.Update(key, mock2)
	if err != nil {
		t.FailNow()
	}
	if r.Data[key] != mock2 {
		t.Fail()
	}
}

func TestGetAll(t *testing.T) {
	// fail if there are two different data types in a region
	key1, key2 := "key1", "key2"
	mock1 := int(10)
	mock2 := int(25)

	r := NewRegion("Test")
	err := r.Add(key1, mock1)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	err = r.Add(key2, mock2)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	all := r.GetAll()
	if 2 != len(all) {
		t.Logf("Did not return both entries. Returned: %v", all)
		t.Fail()
	}

}
