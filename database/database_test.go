package database

const (
	testRegionName     = "Test"
	testListRegionName = "TestList"
)

type mockObj struct {
	ID  string
	Val int
}

type mockObjList struct {
	ID   string
	List []mockObj
}

func newMockDB() *PCDB {
	db := NewPCDB()
	if db == nil {
		panic("database did not initialize")
	}
	err := db.AddRegion(testRegionName)
	if err != nil {
		panic(err)
	}

	// initialize test region with data
	r, err := db.GetRegion(testRegionName)
	if err != nil {
		panic(err)
	}

	key1, key2, key3 := "1", "2", "3"
	val1, val2, val3 := int(3), int(7), int(11)
	o1 := mockObj{key1, val1}
	o2 := mockObj{key2, val2}
	o3 := mockObj{key3, val3}
	err = r.Add(key1, o1)
	if err != nil {
		panic(err)
	}
	err = r.Add(key2, o2)
	if err != nil {
		panic(err)
	}
	err = r.Add(key3, o3)
	if err != nil {
		panic(err)
	}
	if 3 != r.NumEntries() {
		panic("Failed to add 3 test entries")
	}

	// initialize testListRegion with data
	err = db.AddRegion(testListRegionName)
	if err != nil {
		panic(err)
	}

	r, err = db.GetRegion(testListRegionName)
	if err != nil {
		panic(err)
	}

	ol := []mockObj{o1, o2, o3}
	mol1 := &mockObjList{ID: key1, List: ol}
	err = r.Add(key1, mol1)
	if err != nil {
		panic(err)
	}
	if 1 != r.NumEntries() {
		panic("Failed to add 3 test entries")
	}

	return db
}
