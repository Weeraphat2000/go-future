package constants

import "fmt"

type day int

func ConstantsExample() {
	// TODO: Example of constants in Go
	const Pi = 3.14
	const Greeting = "Hello, World!"

	// Printing constants
	println("Pi:", Pi)
	println("Greeting:", Greeting)

	// Pi = 3.14159 // This will cause a compile-time error because Pi is a constant and cannot be reassigned

	// Constants can be used in expressions
	const Radius = 5
	area := Pi * Radius * Radius
	println("Area of circle with radius", Radius, "is", area)

	// TODO: ปรพกาศตัวแปรคงที่แบบกลุ่ม
	const (
		Red   = "ff0000"
		Green = "00ff00"
		Blue  = "0000ff"
	)
	fmt.Println("Colors:", Red, Green, Blue)

	const (
		// TODO: iota คือ จะเริ่มต้นที่ 0 และเพิ่มขึ้นทีละ 1
		Sunday day = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println("Days of the week:", Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
	fmt.Printf("Type of Saturday: %T\n", Saturday) // constants.day คือ package.type

}

func varadicfunction(args ...int) {
	// TODO: args ก็จะกลายเป็น slice ของ int
	// This function takes a variable number of int arguments
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func VaradicExample() {
	// Example of a variadic function
	fmt.Println(" ")
	fmt.Println("Variadic")

	varadicfunction(1, 2, 3, 4, 5) // Passing multiple arguments
	varadicfunction(10, 20)        // Passing fewer arguments
	varadicfunction()              // Passing no arguments
	fmt.Println("End of Variadic Function Example")

	fmt.Println(" ")
	a := []int{1, 2, 3, 4, 5}
	varadicfunction(a...)             // Using a slice with variadic function, the ... operator unpacks
	c := append([]int{1, 1, 1}, a...) // Append b to a, the ... operator unpacks b
	fmt.Println("Slice after append:", c)
}
