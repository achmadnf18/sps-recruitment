package application

import (
	"fmt"
	"strconv"
)

type person struct {
	name   string
	gender string
	age    int
}

func NewPerson() person{
	return person{
		name:   "",
		gender: "",
		age:    0,
	}
}

func (p person) Name(s string) person {
	p.name = s
	return p
}

func (p person) Gender(s string) person {
	p.gender = s
	return p
}

func (p person) Age(i int) person {
	p.age = i
	return p
}

func (p person) ToString() string {
	return p.name + `, ` + p.gender + `, ` + strconv.Itoa(p.age)
}

func GetPerson() {
	jon := NewPerson().Name("Jon Snow").Gender("Male").Age(24).ToString()
	fmt.Println(jon)
}

