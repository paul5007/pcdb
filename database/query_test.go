package database

import "testing"

func TestRunQuery(t *testing.T) {
	db := newMockDB()

	q := Query{Statement: "select * from Test"}
	r, err := db.RunQuery(q)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	rList, err := r.List()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if rList == nil {
		t.Fail()
	}
}

func TestRunQueryWithWhere(t *testing.T) {
	db := newMockDB()

	q := Query{Statement: "select * from Test where Val = 3"}
	r, err := db.RunQuery(q)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	rList, err := r.List()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if 1 != len(rList) {
		t.Logf("Wrong number of results: %v", len(rList))
		for index, each := range rList {
			t.Logf("Result %v: %+v", index+1, each)
		}
		t.Fail()
	}
}
