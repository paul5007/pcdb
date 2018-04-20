package domain

import (
	"fmt"
)

// Person holds information on a person
type Person struct {
	ID   string
	Name string
	Age  int
}

// Valid checks to make sure Person object is valid
func (p *Person) Valid() (bool, error) {
	var err error
	if p.ID == "" {
		err = fmt.Errorf("ID is nil ")
	}
	if p.Name == "" {
		err = fmt.Errorf("%v | Name is nil", err)
	}
	if p.Age <= 0 {
		err = fmt.Errorf("%v | Age: %v <= 0", err, p.Age)
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
