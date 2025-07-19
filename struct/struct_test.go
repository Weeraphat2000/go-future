package structs1

import (
	"fmt"
	"testing"
)

func TestStructExample(t *testing.T) {
	// Create an instance of the struct
	p1 := Person{Name: "Alice", Age: 30}
	fmt.Println("Person 1:", p1)

	// Access fields of the struct
	if p1.Name != "Alice" || p1.Age != 30 {
		t.Errorf("Expected Person 1 to be {Alice 30}, got {%s %d}", p1.Name, p1.Age)
	}

}
