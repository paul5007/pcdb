package database

import (
	"fmt"
	"reflect"
	"strings"
)

// Query contains information for a specific query
type Query struct {
	Statement string
}

// SelectResults is returned by a query will all data gathered
type SelectResults struct {
	resultList []interface{}
}

// RunQuery will execute a query statement
func (db *PCDB) RunQuery(q Query) (*SelectResults, error) {
	statement := q.Statement
	pieces := strings.Split(statement, " ")

	if strings.ToUpper(pieces[0]) != strings.ToUpper("select") {
		return nil, fmt.Errorf("Only SELECT statements are supported")
	}
	if strings.ToUpper(pieces[1]) != strings.ToUpper("*") {
		return nil, fmt.Errorf("Only * scope is supported")
	}
	if strings.ToUpper(pieces[2]) != strings.ToUpper("from") {
		return nil, fmt.Errorf("Only FROM clause is supported")
	}
	regionName := pieces[3]
	r, err := db.GetRegion(regionName)
	if err != nil {
		return nil, fmt.Errorf("Region does not exist. Region names are case sensitive")
	}
	resultSet := &SelectResults{resultList: r.GetAll()}
	if len(pieces) <= 4 {
		return resultSet, nil
	}

	clause := pieces[4]
	if strings.ToUpper(clause) != strings.ToUpper("where") {
		return nil, fmt.Errorf("Only WHERE clause is supported")
	}
	rest := pieces[5:]
	if rest[1] != "=" {
		return nil, fmt.Errorf("Only Equals is supported")
	}
	if field := rest[0]; field != "" {
		if value := rest[2]; value != "" {
			result, err := parsedResults(resultSet, field, value)
			if err != nil {
				return nil, fmt.Errorf("Failed to parse results. %v", err)
			}
			return result, nil
		}
	}
	return nil, fmt.Errorf("Failed to complete query")

}

// List returns all the results as a list
func (sr *SelectResults) List() ([]interface{}, error) {
	return sr.resultList, nil
}

func parsedResults(resultSet *SelectResults, field string, value string) (*SelectResults, error) {
	dataList := reflect.ValueOf(resultSet.resultList)

	if dataList.Kind() == reflect.Slice {
		var newResults []interface{}
		for i := 0; i < dataList.Len(); i++ {
			d := dataList.Index(i).Elem()
			kind := d.Kind()
			_ = kind
			x := d.FieldByName(field)
			xS := fmt.Sprintf("%v", x)
			if xS == value {
				newResults = append(newResults, d)
			}
		}
		return &SelectResults{resultList: newResults}, nil
	}
	return nil, fmt.Errorf("Internal Error: failed to parse results")
}
