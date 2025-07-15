package structs1

import (
	"fmt"
)

// Define a struct type
type Person struct {
	Name string
	Age  int
}

type Person2 struct {
	Name string
	Age  *int
}

// StructExample demonstrates basic operations on structs in Go.
func StructExample() {

	// Create an instance of the struct
	p1 := Person{Name: "Alice", Age: 30}
	fmt.Println("Person 1:", p1)

	// Access fields of the struct
	fmt.Println("Name:", p1.Name)
	fmt.Println("Age:", p1.Age)

	// Modify fields of the struct
	p1.Age = 31
	fmt.Println("Updated Age:", p1.Age)

	// Create another instance of the struct
	p2 := Person{"Bob", 25}
	fmt.Println("Person 2:", p2)

	// Compare structs
	if p1.Name == p2.Name {
		fmt.Println("Both persons have the same name.")
	} else {
		fmt.Println("Persons have different names.")
	}

	pp1 := Person{Name: "Alice", Age: 30}
	pp2 := Person{"Alice", 30}
	pp3 := Person{Name: "Alice"}
	pp4 := Person{}
	fmt.Printf("pp1 %#v\n", pp1) // Will print "pp1 {Alice 30}"
	fmt.Printf("pp2 %#v\n", pp2) // Will print "pp2 {Alice 30}"
	fmt.Printf("pp3 %#v\n", pp3) // Will print "pp3 {Alice 0}" since Age is not initialized (zero value)
	fmt.Printf("pp4 %#v\n", pp4) // Will print "pp4 { 0}" since both Name and Age

	// TODO: ถ้าไม่ได้กำหนดค่าให้กับ field ของ struct จะเป็นค่า zero value (ไม่ได้ใช้ pointer)
	c3 := Person{Name: "Charlie"}
	fmt.Println("Person 3:", c3) // Will print "Person 3: {Charlie 0}" since Age is not initialized

	// TODO: Copy by value
	// When you assign a struct to another variable, it creates a copy of the struct.
	c := p1.Name
	fmt.Println("Variable c:", c) // Will print "Alice"
	p1.Name = "Alice2"
	fmt.Println("Updated Person 1 Name:", p1.Name) // Will print "Alice2
	fmt.Println("Variable c after update:", c)     // Will still print "Alice" since c is a copy of the value

	// Use a slice of structs
	persons := []Person{p1, p2}
	fmt.Printf("Persons slice: %#v\n", persons)

	// TODO: Use pointer to struct ถ้าใช้ pointer จะทำให้ไม่ได้ใส่ค่า zero value เข้าไปใน field ที่เป็น pointer ของ struct
	persons2 := Person2{Name: "Charlie"}
	fmt.Printf("Person2: %#v\n", persons2)          // {Name: "Charlie", Age: <nil>}
	fmt.Printf("Person2 Name: %s\n", persons2.Name) // Will print "Charlie" since Name is initialized
	fmt.Printf("Person2 Age: %v\n", persons2.Age)   // Will print <nil> since Age is not initialized
	fmt.Println(persons2.Age == nil)                // Will print true since Age is not initialized

}
