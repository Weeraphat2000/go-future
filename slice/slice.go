package slice

import "fmt"

func show(s ...string) {
	fmt.Println("show function called with slice:", s)
}

func show2(s []string) {
	fmt.Printf("show2 function called with slice: %#v\n", s)
}

// SliceExample demonstrates basic operations on slices in Go.
func SliceExample() {

	var a []string                                                                   // Declare a slice of strings
	fmt.Println("Initial slice:", a)                                                 // Should print nil
	fmt.Printf("Initial slice : %#v\n", a)                                           // Should print []string
	fmt.Printf("Initial slice type: %T\n", a)                                        // Should print []string
	fmt.Println("Length of initial slice:", len(a))                                  // Should print 0
	fmt.Println("a == nil => ", a == nil)                                            // Should print true
	fmt.Println("Type of a:", fmt.Sprintf("%T", a))                                  // Should print []string
	fmt.Println("Type of slice == []string => ", fmt.Sprintf("%T", a) == "[]string") // Should print []string
	// fmt.Println("a[0] => ", a[0])                                                    // This will panic because a is nil and has no elements

	// TODO:  Declare and initialize a slice ไม่กำหนดขนาด
	slice := []string{"Hello", "World", "!"}

	// Print the slice
	fmt.Println("Slice:", slice)

	// Access elements in the slice
	fmt.Println("First element:", slice[0])
	fmt.Println("Second element:", slice[1])
	fmt.Println("Third element:", slice[2])

	// Modify an element in the slice
	slice[1] = "Go"
	fmt.Println("Modified Slice:", slice)

	// TODO: append elements to the slice
	// Append an element to the slice
	slice = append(slice, "Programming")
	fmt.Printf("After appending: %#v\n", slice)

	// Delete an element from the slice
	// To delete an element, we create a new slice excluding the element to be deleted
	slice = append(slice[:1], slice[2:]...) // Deletes the second element ("Go")
	fmt.Printf("After deleting second element: %#v\n", slice)

	// Delete the last element
	slice = slice[:len(slice)-1] // Removes the last element ("Programming")
	fmt.Printf("After deleting last element: %#v\n", slice)

	// Delete the first element
	slice = slice[1:] // Removes the first element ("Hello")
	fmt.Printf("After deleting first element: %#v\n", slice)

	ss := append(slice, "New Element", "Another Element")
	fmt.Printf("After appending multiple elements(slice): %#v\n", slice)
	fmt.Printf("After appending multiple elements(ss): %#v\n", ss)

	slice = append(slice, "New Element", "Another Element")

	// Iterate over the slice using a for loop
	for i := 0; i < len(slice); i++ {
		fmt.Printf("Element at index %d: %s\n", i, slice[i])
	}

	// Iterate over the slice using range
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}
	a1 := []string{"Go", "Programming", "Language"}
	show(a1...)
	show2(a1)
	a2 := [3]string{"Go", "Programming", "Language"}
	show2(a2[:])                                           // Convert array to slice
	show2(append(a2[:], "New Element", "Another Element")) // Slice from index 1 to the end

	// TODO: copy by address
	a4 := []string{"1", "2", "3"}
	a5 := a4[:]       // Create a slice from a4
	show2(a4)         // {"1", "2", "3"}
	show2(a5)         // {"1", "2", "3"}
	a4[1] = "Updated" // Update the second element
	show2(a4)         // {"1", "Updated", "3"}
	show2(a5)         // {"1", "Updated", "3"} // a5 reflects the change because it points to the same underlying array

	// TODO: copy by value
	fmt.Println(" ")
	fmt.Println("Copy by value examples")
	a6 := []string{"1", "2", "3"}
	// TODO: make a new slice with the same length เป็นเหมือนการประกาศ slice ใหม่ โดยมีค่าเริ่มต้นเป็นค่าเริ่มต้นของชนิดข้อมูลนั้นๆมาให้ สามารถใช้ copy ได้ และ append ได้
	a7 := make([]string, len(a6)) // Create a new slice with the same length
	// a7[40] = "40" // This will panic because a7 has only 3 elements
	fmt.Printf("original slice a6: %#v\n", a6) // int[3]string{"", "", ""}
	fmt.Printf("new slice a7: %#v\n", a7)
	copy(a7, a6) // Copy elements from a6 to a7
	fmt.Printf("After copying, a6: %#v\n", a6)
	fmt.Printf("After copying, a7: %#v\n", a7)
	a6[1] = "Updated" // Update the second element in a6
	fmt.Printf("After updating a6 : %#v\n", a6)
	fmt.Printf("After updating a7 : %#v\n", a7)
	a6 = append(a6, "4444", "5555")                       // Append new elements to a6
	a7 = append(a7, []string{"4", "5", "6", "7", "8"}...) // Append a new element to a7
	fmt.Printf("After appending to a6: %#v\n", a6)
	fmt.Printf("After appending to a7: %#v\n", a7)

	// copy by value 2
	// TODO: This example demonstrates copying a slice by value.
	fmt.Println(" ")
	fmt.Println("Copy by value examples 2")
	b1 := []string{"1", "2", "3"}
	b2 := append([]string{}, b1...) // Create a new slice by copying b1
	fmt.Printf("original slice b1: %#v\n", b1)
	fmt.Printf("new slice b2: %#v\n", b2)
	b1[1] = "Updated" // Update the second element in b1
	fmt.Printf("After updating b1 : %#v\n", b1)
	fmt.Printf("After updating b2 : %#v\n", b2) // b2 remains unchanged
	b1 = append(b1, "4444", "5555")             // Append new elements to b1
	b2 = append(b2, "4", "5", "6", "7", "8")    // Append new elements to b2
	fmt.Printf("After appending to b1: %#v\n", b1)
	fmt.Printf("After appending to b2: %#v\n", b2)

	// copy by value 3
	a8 := []string{"1", "2", "3"}
	var a9 []string // Declare a new slice without initializing
	fmt.Printf("original slice a8: %#v\n", a8)
	fmt.Printf("new slice a9: %#v\n", a9)
	copy(a9, a8) // Attempt to copy elements from a8 to a9
	fmt.Printf("After copying, a8: %#v\n", a8)
	fmt.Printf("After copying, a9: %#v\n", a9) // a9
	a8[1] = "Updated"                          // Update the second element in a8
	fmt.Printf("After updating a8 : %#v\n", a8)
	fmt.Printf("After updating a9 : %#v\n", a9)

}
