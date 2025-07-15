package loop

import "fmt"

func ForBasic() {
	fmt.Println("For Loop with basic initialization")
	for i := 0; i < 5; i++ {
		fmt.Printf("i : %d\n", i)
	}
}

func ForBasic2() {
	fmt.Println("For Loop with basic initialization and increment by 2")
	for i := 0; i < 7; i += 2 {
		fmt.Printf("i : %d\n", i)
	}
}

func ForSum() {
	sum := 0
	fmt.Println("For Loop to calculate the sum of first 5 numbers")
	for i := 1; i <= 5; i++ {
		sum += i
		fmt.Printf("Current sum: %d\n", i) // Uncomment to see the sum at each step
	}
	fmt.Printf("Sum of first 10 numbers: %d\n", sum) // This will print 15
}

func WhileLoop() {
	sum := 0
	fmt.Println("While Loop with condition")
	for sum <= 5 {
		sum += 1
		fmt.Printf("Current sum: %d\n", sum) // Uncomment to see the sum at each step
	}
}

func WhileLoop2() {
	sum := 0
	fmt.Println("While Loop with break condition")
	// loop for ever until break condition is met
	for {
		sum += 1
		if sum > 5 {
			break // Exit the loop when sum exceeds 10
		}
		fmt.Printf("Current sum: %d\n", sum) // Uncomment to see the sum at each step
	}
	fmt.Printf("Final sum: %d\n", sum) // This will print the final sum
}

func ForBasicWithInit() {
	i := 0
	fmt.Println("For Loop with initialization outside the loop")
	for ; i < 5; i++ {
		fmt.Printf("i : %d\n", i)
	}
	fmt.Printf("Final value of i: %d\n", i) // This will print 5
}

func ForMinus() {
	fmt.Println("For Loop with decrementing")
	for i := 5; i > 0; i-- {
		fmt.Printf("i : %d\n", i)
	}
}

func ForRange() {
	fmt.Println("For Loop with range over array and slice")
	arr := [5]int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	slice := []string{"Hello", "World", "!"}
	for i, v := range slice {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}

	for i, v := range slice {
		for j, k := range v {
			fmt.Printf("Outer Index: %d, Inner Index: %d, Value: %s\n", i, j, string(k))
		}
	}

	for _, v := range slice {
		fmt.Printf("Value: %s\n", v)
	}
}

func ForMap() {
	fmt.Println("For Loop with range over map")
	m2 := map[string]string{"name": "Alice", "city": "Wonderland"}
	for k, v := range m2 {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}
}

func ForStruct() {
	fmt.Println("For Loop with range over struct")
	type Person struct {
		Name string
		Age  int
	}

	persons := []Person{
		{"Alice", 30},
		{"Bob", 25},
	}

	for _, p := range persons {
		fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
	}

	// How to announce a struct with different ways
	persons2 := Person{"Charlie", 35}
	persons3 := Person{Name: "Dave", Age: 40}
	for _, p := range []Person{persons2, persons3} {
		fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
	}

}
