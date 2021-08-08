package validation

import (
	"fmt"
	"testing"
)

func TestCheckValidation(t *testing.T) {
	fmt.Println("Inside test validation")
	p := &Product{
		Name:  "Product1",
		Price: 44.54,
		SKU:   "ass-fdffd-fdfdfdf",
	}
	fmt.Println(p)
	err := p.Validate()
	if err != nil {
		t.Fatal("Validation failed", err)
	}

}
