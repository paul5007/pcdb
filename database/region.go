package database

import (
	"fmt"
	"reflect"
)

// Region is a map used for storing data
type Region struct {
	Name string
	Data map[string]interface{}
}

// NewRegion creates a new Region for storage usage
// It is recommended to use PCDB
func NewRegion(name string) *Region {
	d := make(map[string]interface{})
	r := &Region{Name: name, Data: d}
	return r
}

// Add adds an object to map based on key
func (r *Region) Add(key string, value interface{}) error {
	if r.Data[key] != nil {
		return fmt.Errorf("Not allowed to overwrite data using Add. Use update to overwrite value")
	}
	for _, d := range r.Data {
		if reflect.TypeOf(d) != reflect.TypeOf(value) {
			return fmt.Errorf("Not allowed to store two different types of objects in a single region")
		}
		break
	}
	r.Data[key] = value
	return nil
}

// Remove removes an object from map based on key
func (r *Region) Remove(key string) error {
	r.Data = nil
	return nil
}

// Get returns the object stored in the region if it exists
func (r *Region) Get(key string) (interface{}, error) {
	val := r.Data[key]
	if val == nil {
		return nil, fmt.Errorf("Key: %v does not exist in region: %v", key, r.Name)
	}
	return val, nil
}

// Update updates a key with a new value
func (r *Region) Update(key string, value interface{}) error {
	for _, d := range r.Data {
		if reflect.TypeOf(d) != reflect.TypeOf(value) {
			return fmt.Errorf("Not allowed to store two different types of objects in a single region")
		}
		break
	}
	r.Data[key] = value
	return nil
}

// NumEntries returns the number of entries in Data
func (r *Region) NumEntries() int {
	return len(r.Data)
}
