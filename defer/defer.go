package deferss

import "fmt"

func lastFunction() {
	// This function will be deferred
	fmt.Println("This is the last function to execute.")
	fmt.Println(". . .")
}

func DeferExample() {

	// TODO: defer คือ การเลื่อนการทำงานของฟังก์ชันไปจนกว่าจะถึงจุดที่ฟังก์ชันนั้นถูกเรียกใช้
	defer lastFunction() // This will be executed last
	defer fmt.Println("This will be printed second.")
	defer fmt.Println("This will be printed first.")

	fmt.Println("This will be printed first.")
	fmt.Println("Defer statements are executed in LIFO order.")

	for i := range 3 {
		defer fmt.Println("Deferred call", i)
	}

	for i := range 3 {
		defer func(i int) {
			fmt.Println("Deferred call with argument", i)
		}(i)
	}

}
