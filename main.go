// all file in go must have a package
// go mod init github.com/weeraphat2000/myProject
// go mod tidy คือ เหมือน pnpm install ใน Node.js
// go run main.go
// go get ...
package main

import (
	"fmt"

	"math"
	"myapp/arrays"
	changevalue "myapp/changeValue"
	"myapp/constants"
	deferss "myapp/defer"
	errorss "myapp/error"
	function1 "myapp/function"
	interfaces "myapp/interface"
	"myapp/json"
	"myapp/loop"
	"myapp/maps"
	methods1 "myapp/method"
	"myapp/movies"
	"myapp/pointer"
	"myapp/slice"
	structs1 "myapp/struct"
)

// TODO: init function จะถูกเรียกอัตโนมัติเมื่อโปรแกรมเริ่มทำงาน
// TODO: import package อะไรก่อนก็จะเรียกใช้งาน init function ของ package นั้นก่อน
func init() {
	// init function is called before main function
	fmt.Println("init main is called")
}

var a int = 10 // global variable
// b: = 15 // can not use short variable declaration outside function

var c int = 30 // global variable

// function declaration
func add(a int, b int) int {
	return a + b
}

func addAndsubAndEqual(a int, b int) (int, int, bool) {
	return a - b, a + b, a == b
}

// function parameter for array
func functionForArray3Lenght(sk [3]string) {
	fmt.Printf("Array in show function: %#v\n", sk)
}

func functionForSlice(sk []string) {
	fmt.Printf("Slice in show function: %#v\n", sk)
}

func main() {
	// fmt it means format
	// Println is print line
	fmt.Println("Hello, World!")

	// variable declaration
	var name string = "John Doe"
	fmt.Println("Original Name:", name)
	name = "Weeraphat Yian"
	fmt.Println("Updated Name:", name)
	var x string
	x = "Hello, Go!"
	fmt.Println("Variable x:", x)
	msg1 := "Welcome to"
	msg2 := "Go programming"
	fmt.Println(msg1 + msg2)

	fmt.Println("Global variable a:", a)
	a = 20
	fmt.Println("Updated Global variable a:", a)

	// short variable declaration
	// short variable can not be used outside the function
	age := 25
	fmt.Println("Original Age:", age)
	age = 30
	fmt.Println("Updated Age:", age, "years old")
	fmt.Println("%", 25%7) // 25 modulus 7 = 4

	// integer variable
	var floatVar float64 = 3.14
	fmt.Println("Float Variable:", floatVar)

	// boolean variable
	var boolVar bool = true
	fmt.Println("Boolean Variable:", boolVar)

	// multiple variable declaration
	aa, bb, cc := 5, "Hun", 3.14
	fmt.Println("Values of aa and bb and cc:", aa, bb, cc)
	var dd, ee int = 10, 20
	fmt.Println("Values of dd and ee:", dd, ee)

	// workshop
	fmt.Println(" ")
	fmt.Println("Workshop")
	// \n for new line
	// \t for tab
	// \r for carriage return
	// \v for vertical tab
	// \b for backspace
	// \f for form feed
	// \a for alert (bell)
	// \x for hexadecimal escape sequence
	// \u for unicode escape sequence
	fmt.Print("name: Hun \nage:\t20 \nheight: 170 cm \nweight: 60 kg \n")

	fmt.Println(" ")
	fmt.Println("Workshop with Printf")
	fmt.Printf("Name: %s\nAge: %d\nHeight: %.2f cm\nWeight: %.2f kg\n true%t", "Hun", 20, 170.0, 60.00900, true)

	// raw string
	fmt.Println(" ")
	fmt.Println("Workshop with Raw String ``")
	fmt.Println(`Name: Hun
Age: 20
Height: 170 cm
Weight: 60 kg`)

	fmt.Printf("%+v\n",
		struct {
			Name string
			Age  int
		}{
			Name: "Hun",
			Age:  20,
		})

	// runes
	fmt.Println(" ")
	fmt.Println("Workshop with Runes")
	var r rune = 'A'                       // rune is an alias for int32
	fmt.Println("Rune value:", r)          // A is 65 in ASCII
	r2 := '✅'                              // another rune
	fmt.Println("Another Rune value:", r2) // ✅ is a check mark emoji
	fmt.Printf("Rune %c has Unicode code point U+%04X\n", r2, r2)

	// %#v for everything
	fmt.Println(" ")
	fmt.Println(`Workshop with %#v for printing any value`)
	fmt.Printf(
		" String: %#v\n int: %#v\n float64: %#v\n bool: %#v\n rune: %#v\n", "Hello, World!",
		42, 3.14, true, 'A')
	// %T for type
	fmt.Println(" ")
	fmt.Println("Workshop with %T for printing type of value")
	fmt.Printf("Type of string: %T\n", "Hello, World!")
	fmt.Printf("Type of int: %T\n", 42)
	fmt.Printf("Type of float64: %T\n", 3.14)
	fmt.Printf("Type of bool: %T\n", true)

	// zero value
	// default value of variable in Go (zero value)
	fmt.Println(" ")
	fmt.Println("Zero Value in Go")
	var zeroInt int
	var zeroFloat float64
	var zeroString string
	var zeroBool bool
	fmt.Println("Zero value of int:", zeroInt)           // 0
	fmt.Println("Zero value of float64:", zeroFloat)     // 0.0
	fmt.Println("Zero value of string:", zeroString)     // "" (empty string) without quotes
	fmt.Println("Zero value of bool:", zeroBool)         // false
	fmt.Printf("Zero value of string: %q\n", zeroString) // prints empty string with quotes

	// >= <= > <= operators
	fmt.Println(" ")
	fmt.Println("Comparison Operators >= <= > <=")
	var x1 int = 10
	var y1 int = 20
	fmt.Println("x1 >= y1:", x1 >= y1) // false
	fmt.Println("x1 <= y1:", x1 <= y1) // true
	fmt.Println("x1 > y1:", x1 > y1)   // false
	fmt.Println("x1 < y1:", x1 < y1)   // true

	// if statement
	fmt.Println(" ")
	fmt.Println("If")
	var number int = 10
	if number >= 0 {
		fmt.Printf("The number %d is positive.\n", number)
	} else {
		fmt.Printf("The number %d is not positive.\n", number)
	}

	// if-else statement
	fmt.Println(" ")
	fmt.Println("If-Else")
	var age2 int = 16
	if age2 >= 18 {
		fmt.Println("You are an adult.")
	} else if age2 >= 13 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are a child.")
	}

	// if-else with short variable declaration
	fmt.Println(" ")
	fmt.Println("If-Else with Short Variable Declaration")
	if v12 := 10; v12 > 50 {
		fmt.Println("v12 more", v12)
	} else {
		fmt.Println("v12 less", v12)
	}
	// fmt.Println("v12", v12) // v12 is not accessible here, it will cause an error

	// workshop with if-else
	fmt.Println(" ")
	fmt.Println("Workshop with If-Else and Short Variable Declaration")
	if score := 85; score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 && score < 90 {
		fmt.Println("Grade: B")
	} else if score >= 70 && score < 80 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: D")
	}

	// && || operators
	// TODO: operator && , ||
	fmt.Println(" ")
	fmt.Println("Logical Operators && and ||")
	var t bool = true
	var f bool = false
	fmt.Println("&& Operator: ", t && f) // true if both are true
	fmt.Println("|| Operator: ", t || f) // true if at least

	// math package
	fmt.Println(" ")
	fmt.Println("Math Package")
	v := math.Pow(2, 3) // 2 raised to the power of 3
	fmt.Println("2 raised to the power of 3 is:", v)

	// switch case statement
	fmt.Println(" ")
	fmt.Println("Switch Case Statement")
	switch day := 3; day {
	case 1:
		fmt.Println("Monday")
		// break // break is optional, it will break the switch case
		// TODO: if you want to continue to the next case, you can use fallthrough
		fallthrough // fallthrough means it will continue to the next case
	case 2:
		fmt.Println("Tuesday ")
	case 3:
		fmt.Println("Wednesday")
	case 4, 5: // multiple cases can be combined
		fmt.Println("Thursday or Friday")
	default:
		fmt.Println("Invalid day")
	}
	day := "Friday"
	fmt.Println("day is:", day)

	// switch case with string
	fmt.Println(" ")
	fmt.Println("Switch Case")
	num := 1
	switch {
	case num > 1 && num < 5:
		fmt.Println("num is 1")
		fallthrough // fallthrough means it will continue to the next case
	case num == 2:
		fmt.Println("num is 2")
	default:
		fmt.Println("num is neither 1 nor 2")
	}

	// function declaration
	fmt.Println(" ")
	fmt.Println("Function Declaration")
	fmt.Println("The sum of 10 and 20 is:", add(10, 20))

	// function multiple return values
	fmt.Println(" ")
	fmt.Println("Function with Multiple Return Values")
	// TODO: function returns multiple values and not use some values
	a1, _, c1 := addAndsubAndEqual(10, 20)
	fmt.Println("The difference is:", a1)
	fmt.Println("Are they equal?", c1)

	// function from another package
	fmt.Println(" ")
	fmt.Println("Function from Another Package")
	fmt.Println("The sum is:", function1.FunctionAddTwoNumbers(10, 20)) // package.function for calling function from another package
	A1, B1, C1 := function1.FunctionAddAndSubAndEqual(10, 20)
	fmt.Println("The difference is:", A1)
	fmt.Println("The sum is:", B1)
	fmt.Println("Are they equal?", C1)
	D1, E1, F1 := function1.FunctionAddMinusEqual(10, 20)
	fmt.Println("The sum from function2 is:", D1)
	fmt.Println("The difference from function2 is:", E1)
	fmt.Println("Are they equal from function2?", F1)
	function1.IntPlusFloat(10, 20.5)
	G1, H1 := function1.Split(10, 20)
	fmt.Println("The sum from Split function is:", G1)
	fmt.Println("The sum from Split function as float is:", H1)

	// function with variable number of arguments
	fmt.Println(" ")
	fmt.Println("Function with Variable Number of Arguments")
	add2 := func(a int, b int) int {
		return a + b
	}
	fmt.Println("The sum from anonymous function is:", add2(10, 20))
	fmt.Printf("%T\n", add2) // %T for type

	// signature of function
	fmt.Println(" ")
	fmt.Println("Function Signature")
	// TODO: function declaration with signature and variable declaration
	// TODO: function ก็เหมือน ตัวแปร ตัวหนึ่ง
	var add3 func(int, int) float64 = func(a int, b int) float64 {
		return float64(a + b)
	}
	fmt.Printf("%T\n", add3) // %T for type
	function1.IntPlusFloat(10, 20.5)

	// parameter function
	fmt.Println(" ")
	fmt.Println("Function with Function as Parameter")
	// TODO: function that takes another function as parameter and send arguments true signature
	fmt.Println("compute function with function as parameter", function1.Compute(func(a, b int) int { return a + b }, 10, 20))

	// function return function
	fmt.Println(" ")
	fmt.Println("Function Return Function")

	// TODO: function that returns another function and can access the state
	addFunc, subFunc := function1.Adder2()    // returns two functions
	fmt.Println("Add 10:", addFunc(10))       // Add 10
	fmt.Println("Add 5:", addFunc(5))         // Add 15
	fmt.Println("Subtract 10:", subFunc(100)) // Subtract -85
	fmt.Println("Subtract 5:", subFunc(5))    // Subtract -90

	fmt.Println(" ")
	fmt.Println("Array")
	fmt.Println(arrays.Arr)                                 // Accessing exported variable from arrays package
	fmt.Printf("%#v\n", arrays.ArrString)                   // Accessing another exported variable from arrays package, and other value it zero value if not initialized
	fmt.Printf("%+v\n", arrays.ArrString)                   // Accessing another exported variable from arrays package, and other value it zero value if not initialized
	fmt.Printf("%#v\n", arrays.ArrInt)                      // Accessing another exported variable from arrays package, and other value it zero value if not initialized
	fmt.Printf("%+v\n", arrays.ArrFloat)                    // Accessing another exported variable from arrays package, and other value it zero value if not initialized
	fmt.Printf("%#v\n", arrays.ArrBool)                     // Accessing another exported variable from arrays package, and other value it zero value if not initialized
	fmt.Println("Length of Arr:", len(arrays.Arr))          // Length of array
	fmt.Printf("%#v\n", arrays.ArrString[1])                // Print the array with %v format
	fmt.Printf("%#v\n", arrays.ArrStringAndInistialized)    // Accessing another exported variable from arrays package, and other value it initialized
	fmt.Printf("%#v\n", arrays.ArrStringAndInistialized[0]) // "Hello"
	fmt.Println(arrays.ArrStringAndInistialized[0])         // Hello
	arrays.ArrStringAndInistialized[0] = "Weeraphat Yian"   // Update the value of the first element
	fmt.Println(arrays.ArrStringAndInistialized[0])         // Weeraphat Yian
	// length of array
	fmt.Println("Length of ArrStringAndInistialized:", len(arrays.ArrStringAndInistialized)) // Length of array
	sh := [3]string{"Hello", "World", "!"}                                                   // Declare an array with 3 elements
	functionForArray3Lenght(sh)                                                              // Call the show function with the array
	// TODO: ประกาศตัวแปร array ที่มีขนาดตามค่าที่ใส่เข้าไป เดี๋ยว go ก็จะรู้เองว่าขนาดเท่าไหร่
	arrayDotDotDot := [...]string{"Hello", "World", "..."}          // Declare an array with initialized values(length)
	fmt.Printf("Array with ...: %#v\n", arrays.ArrDot)              // Print the array with ... format
	fmt.Printf("Array string with ...: %#v\n", arrays.ArrDotString) // Print the array with ... format
	functionForArray3Lenght(arrayDotDotDot)                         // Call the show function with the array
	// TODO: ถ้าจะใช้ append ต้องแปลง array เป็น slice ก่อน ด้วยการใช้ sh[:]
	arrayAppend := append(sh[1:], "New Element")            // ถ้าจะให้ append ต้องแปลงเป็น slice ก่อน คือ sh[:] จะได้ slice
	fmt.Printf("Array after appending: %#v\n", arrayAppend) // ["World" "!" "New Element"]

	// slice
	fmt.Println(" ")
	fmt.Println("Slice")
	var slicessss []string // nil slice, not initialized
	// TODO: append is used to add elements to a slice, it will create a new slice if the original slice is nil or not initialized and return a new slice for new or old variable
	slicessss2 := append(slicessss, "New Element")                  // Append a new element to the slice and assign back
	fmt.Printf("len for initialized slice : %#v\n", len(slicessss)) // nil slice, not initialized
	fmt.Printf("Slice before appending: %#v\n", slicessss)          // nil slice, not initialized
	fmt.Printf("Slice after appending: %#v\n", slicessss2)          // ["New Element"]
	functionForSlice(slicessss2)                                    // Call the function with the slice
	array := [3]string{"Hello", "World", "!"}                       // Declare a slice with 3 elements
	functionForSlice(array[:])                                      // can not functionForSlice(array) // because array is not a slice, but we can convert it to a slice using array[:]

	//
	fmt.Println(" ")
	fmt.Println("Workshop: Loop")
	loop.ForBasic()         // Call the Loop function from loop package
	loop.ForBasic2()        // Call the Loop function from loop package
	loop.ForBasicWithInit() // Call the Loop function from loop package
	loop.ForSum()           // Call the Loop function from loop package
	loop.ForMinus()         // Call the Loop function from loop package
	loop.ForRange()         // Call the Loop function from loop package
	loop.ForMap()           // Call the Loop function from loop package
	loop.ForStruct()        // Call the Loop function from loop package
	loop.WhileLoop()        // Call the Loop function from loop package
	loop.WhileLoop2()       // Call the Loop function from loop package

	//
	fmt.Println(" ")
	fmt.Println("Slice")
	slice.SliceExample()
	slice.SliceAdvancedExample() // Call the SliceAdvancedExample function from slice package

	//
	fmt.Println(" ")
	fmt.Println("Struct")
	structs1.StructExample() // Call the StructExample function from structs package

	//
	fmt.Println(" ")
	fmt.Println("Methods")
	methods1.Methods() // Call the Methods function from methods package
	//

	//
	fmt.Println(" ")
	fmt.Println("Pointer")
	pointer.PointerExample()  // Call the PointerExample function from pointer package
	pointer.PointerExample2() // Call the PointerExample2 function from pointer package
	pointer.PointerExample3() // Call the PointerExample3 function from pointer package
	pointer.PointerExample4() // Call the PointerExample4 function from pointer package
	pointer.WorkshopPointer() // Call the WorkshopPointer function from pointer package

	//
	fmt.Println(" ")
	fmt.Println("Interface")
	interfaces.InterfacesExample()
	interfaces.InterfacesExample2()

	//
	fmt.Println(" ")
	fmt.Println("Error")
	errorss.ErrorExample()

	//
	fmt.Println(" ")
	fmt.Println("Map")
	maps.MapExample()  // Call the MapExample function from maps package
	maps.WorkshopMap() // Call the WorkshopMap function from maps package

	//
	fmt.Println(" ")
	fmt.Println("Change Value")
	changevalue.ChangeValue() // Call the ChangeValue function from changevalue package

	//
	//
	//
	fmt.Println(" ")
	fmt.Println("Example Project")
	fmt.Println(movies.ReviewMovie("Inception", 5))  // Call the ReviewMovie function from movies package
	fmt.Println(movies.FindMovieById("1"))           // Call the FindMovieById function from movies package
	fmt.Println(movies.CheckTicketAvailability("1")) // Call the CheckTicketAvailability function from movies package

	//
	fmt.Println(" ")
	fmt.Println("JSON")
	json.JsonExample()
	json.JsonWorkShop()

	//
	fmt.Println(" ")
	fmt.Println("Constants")
	constants.ConstantsExample() // Call the ConstantsExample function from constants package
	constants.VaradicExample()   // Call the VaradicExample function from constants package

	//
	fmt.Println(" ")
	fmt.Println("Defer")
	deferss.DeferExample()

	//
	fmt.Println(" ")
	str := "Hello123World456"
	sum := 0

	// Iterate over each character in the string
	for index, char := range str {
		// Check if the character is a digit
		if char >= '0' && char <= '9' {
			// Convert the character to its integer value and add it to the sum
			sum += int(char - '0')
		}
		fmt.Printf("Character at index %d: %c\n", index, char)
	}

	// Print the result
	fmt.Println("The sum of digits in the string is:", sum)
}
