package interfaces

import (
	"fmt"
)

func show(msg string) {
	fmt.Println(msg)
}

func showAny(msg any) {
	fmt.Println(msg)
}

func plusAny(a, b any) {

	// TODO: value, ok := variable.(type) คือ เช็คว่า variable เป็น type นี้ไหม เป็นสิ่งที่ควรทำ
	// TODO: value := variable.(type) คือ ต้องชัวร์ว่า variable เป็น type นี้จริง ๆ ถ้าไม่ใช่จะ panic
	a1, ok1 := a.(int) // Attempt to assert a as an int
	b1, ok2 := b.(int) // Attempt to assert b as an int

	fmt.Printf("a: %v, b: %v, a1: %v, b1: %v, a1+b1: %v \n", a, b, a1, b1, a1+b1)
	if ok1 || ok2 {
		fmt.Println("a+b:", a1+b1)
	}
}

func showSwitchCase(value any) {

	// TODO: value.(type) ใช้ได้เฉพาะใน switch case เท่านั้น
	// a, b := value.(type) // ใช้ได้เฉพาะใน switch case เท่านั้น

	switch v := value.(type) {
	case string:
		fmt.Println("string", v)
	case int:
		fmt.Println("int", v)
	default:
		fmt.Println("any")
	}
}

func InterfacesExample() {
	fmt.Println("InterfacesExample 1")
	var i interface{}                         // Declare an empty interface variable
	i = 42                                    // Assign an integer to the interface
	fmt.Printf("Value: %v, Type: %T\n", i, i) // Print the value and type of the interface
	i = "Hello, World!"                       // Assign a string to the interface
	fmt.Printf("Value: %v, Type: %T\n", i, i) // Print the value and type of the interface

	// TODO: empty interface == any คือ ทั้ง 2 อันนี้ type เป็นอะไรก็ได้
	// TODO: interface {} == any
	var v any                                 // Declare a variable of type 'any' (which is an alias for interface{})
	v = 3.14                                  // Assign a float64 to the variable
	fmt.Printf("Value: %v, Type: %T\n", v, v) // Print the value and type of the variable
	v = true                                  // Assign a boolean to the variable
	fmt.Printf("Value: %v, Type: %T\n", v, v) // Print the value and type of the variable

	// Use type assertion to convert v to string before passing to show
	var msg any = "This is a message"
	// TODO: valiable.(type) คือ เช็คว่าตัวแปรนี้เป็น type นี้ใช่ไหม
	str, ok := msg.(string)                  // Attempt to assert msg as a string
	fmt.Printf("str: %v, ok: %v\n", str, ok) // Print the result of the assertion
	if ok {
		show(str) // Call show with the string value
	} else {
		fmt.Println("msg is not a string")
	}

	s, ok := msg.(int)                   // Attempt to assert msg as an int
	fmt.Printf("s: %v, ok: %v\n", s, ok) // s: 0, ok: false ถ้าไม่ใช่ type ก็จะได้ค่า zero value ของ type นั้น ๆ และ ok จะเป็น false

	showAny(msg) // Call showAny with the msg variable
	msg = 100    // Change msg to an integer
	showAny(msg) // Call showAny with the new integer value

	// TODO: any รับค่าได้ทุก type แต่ type any ส่งค่าไปได้แค่ type any
	showAny(42)              // Call showAny with an integer literal
	showAny("Hello, World!") // Call showAny with a string literal
	showAny(3.14)            // Call showAny with a float64 literal
	showAny(true)            // Call showAny with a boolean literal

	plusAny(10, 20)           // 30
	plusAny("10", "20")       // Call plusAny with two strings (will not add them since they are not integers)
	plusAny(10, "20")         // 10
	plusAny(10.5, 20.5)       // Call plusAny with two float64s (will not add them since they are not integers)
	plusAny("Hello", "World") // Call plusAny with two strings (will not add them since they are not integers
	plusAny(true, false)      // Call plusAny with two booleans (will not add them since they are not integers)
	plusAny(10, true)         // 10

	showSwitchCase("Hello") // Call showSwitchCase with a string
	showSwitchCase(42)      // Call showSwitchCase with an integer
	showSwitchCase(3.14)    // Call showSwitchCase with a float64
}

// TODO: ถ้าอยากเป็น interface promotion ต้องมี method ที่ชื่อว่า discount() ไม่รับ parameter ใดๆ และ return type เป็น int (มี method มากกว่า interface นี้กำหนดไว้ก็ได้ แต่ต้องมี method ที่ interface กำหนดไว้ด้วย)
type promotion interface {
	discount() int
}

type Stringer interface {
	String() string // String method to return a string representation of the type
}

// TODO: premoter interface ต้องมี method ที่ชื่อว่า discount() และ String() ซึ่งเป็นไปตาม interface promotion และ Stringer (เหมือนการรวม interface ทั้งสองเข้ามาไว้ใน interface premoter)
type premoter interface {
	promotion // Embedding the promotion interface
	Stringer  // Embedding the Stringer interface
}

func printPromotion(p premoter) {
	fmt.Printf("Promotion: %s, Discount: %d\n", p.String(), p.discount())
}

type product struct {
	name  string
	price int
}

// TODO: product มี method ที่ชื่อว่า discount() ซึ่งเป็นไปตาม interface promotion
func (p product) discount() int {
	return p.price / 10 // Example discount logic: 10% off
}

func (p product) String() string {
	return fmt.Sprintf("Product: %s, Price: %d", p.name, p.price)
}

func sales(p promotion) {
	fmt.Printf("discount: %d\n", p.discount()) // Call the discount method on the promotion interface
}

func InterfacesExample2() {

	fmt.Println(" ")
	fmt.Println("Interfaces Example 2")

	p := product{name: "Laptop", price: 1000}                                              // Create a product instance
	var prom promotion = p                                                                 // Assign the product to the promotion interface
	fmt.Printf("Product: %s, Price: %d, Discount: %d\n", p.name, p.price, prom.discount()) // Print product details and discount
	// Output: Product: Laptop, Price: 1000, Discount: 100
	fmt.Println("p.discount():", p.discount())     // Call the discount method directly on the product instance
	fmt.Println("p.String():", p.String())         // Call the String method on the product instance
	p.name = "Gaming Laptop"                       // Change the product name
	fmt.Println("Updated p.String():", p.String()) // Print the updated product details

	var i1 promotion = product{name: "Phone", price: 500} // Assign a product to the promotion interface
	fmt.Printf("Product: %s, Price: %d, Discount: %d\n", i1.(product).name, i1.(product).price, i1.discount())

	// TODO: sales function expects a promotion interface, so we can pass any type that implements the discount() method
	p1 := product{name: "Tablet", price: 300} // Create another product instance
	// p1 มันต้อง discount() method เพื่อให้เป็น promotion interface
	// p1.discount()
	sales(p1)                              // Call the sales function with the product instance
	fmt.Println("sales(p1):", p1.String()) // Print the product details after calling sales
	printPromotion(p1)                     // Call printPromotion with the product instance, which implements both promotion and Stringer interfaces

}
