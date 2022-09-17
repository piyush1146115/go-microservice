package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "Pk",
		Price: 100,
		SKU:   "abc-def-abc",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
