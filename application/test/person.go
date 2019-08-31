package test

import (
	"errors"
)

type person struct {
	name   string
	gender string
	age    int
}

func (p person) Validate() error {
	if p.name == "" {
		return errors.New("Name cannot be empty")
	}
	if p.gender != "Male" && p.gender != "Female" {
		return errors.New("Gender is either Male or Female")
	}
	if p.age < 0 {
		return errors.New("There is no such thing as negative age")
	}
	return nil
}
