package test

import (
	"testing"
)

func TestPerson_Validate(t *testing.T){
	t.Run("Empty Name", func(t *testing.T) {
		np := person{name: "",gender:"M",age:24}
		if err := np.Validate(); err != nil {
			t.Fatalf(`Error msg: %v`, err)
		}
	})

	t.Run("Empty Gender", func(t *testing.T) {
		np := person{name: "John",gender:"L",age:24}
		if err := np.Validate(); err != nil {
			t.Fatalf(`Error msg: %v`, err)
		}
	})

	t.Run("Negative Age", func(t *testing.T) {
		np := person{name: "John",gender:"Male",age:-1}
		if err := np.Validate(); err != nil {
			t.Fatalf(`Error msg: %v`, err)
		}
	})
}