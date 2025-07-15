package maps

import (
	"fmt"
	"strings"
)

func MapExample() {
	// Example of using maps in Go
	// Maps are key-value pairs, similar to dictionaries in Python or objects in JavaScript.

	// TODO: zero value of map
	var m map[string]int // Create an empty map with string keys and int values
	fmt.Printf("m: %#v\n", m)

	// TODO: make(map[string]int) == map[string]int{}

	// Create a map
	// TODO: map[KeyType]ValueType is a map with the specified key and value types
	a := map[string]int{"apple": 5, "banana": 10}
	fmt.Printf("a: %#v\n", a)

	// TODO: make(map[KeyType]ValueType) returns a map with the specified key and value types
	myMap := make(map[string]int)
	myMap2 := make(map[string]any)

	// TODO: zero value of any type is nil

	// Add key-value pairs to the map
	myMap["apple"] = 5
	myMap["banana"] = 10
	myMap["orange"] = 15
	// myMap["grape"] = "asdf" // This will cause a compile-time error because the value type is int

	myMap2["name"] = "John Doe"
	myMap2["age"] = 30
	myMap2["isStudent"] = false
	myMap2["grades"] = []int{90, 85, 88}

	// Access values by key
	fmt.Println("Apple count:", myMap["apple"])
	fmt.Println("Banana count:", myMap["banana"])

	fmt.Printf("Apple: %d\n", myMap["apple"])
	// TODO: if have key, it will return value, if don't have key
	delete(myMap, "apple") // Remove the key "apple" from myMap
	// ลบ 2 ครั้งก็ไม่เป็นอะไร
	delete(myMap, "apple") // Remove the key "apple" from myMap
	// TODO: if don't have key, it will return zero value of value type
	fmt.Printf("Apple after delete: %d\n", myMap["apple"]) // Will print 0 because "apple" key is deleted

	fmt.Println("isStudent before delete:", myMap2["isStudent"])
	// TODO: delete(map, key) removes the key-value pair from the map
	delete(myMap2, "isStudent") // Remove the key "isStudent" from myMap2
	// TODO: if have key, it will remove key and if don't have key, it will not do anything
	delete(myMap2, "isStudent") // Remove the key "isStudent" from myMap2
	// TODO: if don't have key, it will return nil (value of any type)
	fmt.Println("isStudent after delete:", myMap2["isStudent"])

	fmt.Printf("Type of myMap2: %T\n", myMap2)
	fmt.Printf("myMap2: %#v\n", myMap2)

	//
	// TODO: ถ้ามี key ก็จะได้ value ถ้าไม่มี key จะได้ nil (value of any type)
	value := myMap2["gradessssssssss"]
	fmt.Printf("gradessssssssss : %v\n", value)

	//
	// TODO: Check if a key exists in the map
	value2, exists2 := myMap2["gradessss"]
	fmt.Printf("value2 : %v, exists2: %v\n", value2, exists2)

	//
	// Check if a key exists
	if value, exists := myMap["orange"]; exists {
		println("Orange count:", value)
	} else {
		println("Orange not found")
	}

	// Iterate over the map
	for key, value := range myMap {
		println(key, ":", value)
	}
}

func WordCount(s string) map[string]int {
	// Create a map to store word counts
	wordCount := make(map[string]int)

	// Split the string into words
	// TODO: strings.Fields(s) splits the string by whitespace and returns a slice of words
	words := strings.Fields(s)
	fmt.Printf("Words: %#v", words)
	// Iterate over the words and count them
	for _, word := range words {
		// Convert word to lowercase for case-insensitive counting
		// word = strings.ToLower(word)

		// Increment the count for the word in the map
		wordCount[word] += 1
		// wordCount[word]++
	}
	return wordCount
	// The map will contain the count of each word in the string

	// words := strings.Split(s, " ")
	// fmt.Printf("Words: %#v", words)
	// for _, word := range words {
	// 	// word = strings.ToLower(word) // Convert word to lowercase for case-insensitive counting
	// 	if value, exists := wordCount[word]; exists {
	// 		wordCount[word] = value + 1
	// 	} else {
	// 		wordCount[word] = 1
	// 	}
	// }
	// return wordCount
}

func WorkshopMap() {

	//
	fmt.Println(" ")
	fmt.Println("Workshop Map")
	s := "If it looks      like a duck, swims like a duck, and quacks like a duck, then it probably is a duck."
	wordCount := WordCount(s)
	fmt.Println("Word Count:", wordCount)
}
