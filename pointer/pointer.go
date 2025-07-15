package pointer

import "fmt"

func PointerExample() {
	fmt.Println("Pointer Example")
	// Create a variable
	var x int = 10
	var x1 int = 10

	// Create a pointer to the variable
	// TODO: ตัวแปร p เก็บ address ของ x
	var p *int = &x
	// pp := &x // pp is a pointer to x, it holds the address of x

	// TODO: ตัวแปร p1 เก็บ value ของ x1
	var p1 int = x1

	fmt.Printf("Type of p: %T\n", p) // Output: Type of p: *int (pointer to int)
	// Print the value of the variable and the pointer
	println("Value of x:", x) // Output: Value of x: 10
	// TODO: dereference pointer p to get the value
	println("Value of *p:", *p) // Output: Value of *p: 10
	// TODO: print address of x (ตัวแปรธรรมดา)
	println("Address of x:", &x) // Output: Address of x: <memory address>
	// TODO: print address of p (pointer หรือตอนประกาศมันเก็บ address)
	println("Address stored in p:", p) // Output: Address stored in p: <memory address>

	println("Value of x1:", x1) // Output: Value of x1: 10
	println("Value of p1:", p1) // Output: Value of p1: 10

	x = 20                                   // Change the value of x
	x1 = 20                                  // Change the value of x1
	println("Value of x after change:", x)   // Output: Value of x after change: 20
	println("Value of *p after change:", *p) // Output: Value of *p after change: 20

	println("Value of x1 after change:", x1) // Output: Value of x1 after change: 20
	println("Value of p1 after change:", p1) // Output: Value of p1 after change: 10

	// TODO: change value of x using pointer p
	// TODO: เอา value 30 มาเก็บใน address p
	*p = 30                                                  // Change the value of x using the pointer
	println("Value of x after change through pointer:", x)   // Output: Value of x after change through pointer: 30
	println("Value of *p after change through pointer:", *p) // Output: Value of *p after change through pointer: 30

	// TODO: เอา value ใน address p มาเก็บใน v
	v := *p
	println("Value of v after change through pointer:", v)   // Output: Value of v after change through pointer: 30
	v = 40                                                   // Change the value of v (this does not change x)
	println("Value of v after change:", v)                   // Output: Value of v after change : 40
	println("Value of x after change through pointer:", x)   // Output: Value of x after change through pointer: 30
	println("Value of *p after change through pointer:", *p) // Output: Value of *p after change through pointer: 30
}

// TODO: ให้ส่ง address เข้ามาใน parameter p
func changeValue(p *int) {
	// Change the value of the variable using the pointer
	// TODO: ถ้าเป็นตัวแปรธรรมดา จะต้องมี *
	*p = 100                                           // Change the value at the address stored in p
	fmt.Println("function changeValue value of p:", p) // Output: value of p: <memory address>
	// address จริงๆ ของ p => &p
	fmt.Println("function changeValue address of p:", &p) // Output: address of p: <memory address>
}

// TODO: pass by value
func changeValue2(p int) {
	// Change the value of the variable using the pointer
	p = 1000000                                            // This does not change the original variable
	fmt.Println("function changeValue2 address of p:", &p) // Output: address of p: <memory address>
}

func PointerExample2() {
	fmt.Println(" ")
	fmt.Println("Pointer Example 2")

	// Create a variable
	// TODO: ตัวแปร x เก็บค่า 10
	// TODO: &x คือ address ของตัวแปร x
	var x int = 10

	// Create a pointer to the variable
	// TODO: ตัวแปร p เก็บ address ของ x
	// TODO: &p คือ address ของตัวแปร p
	var p *int = &x

	// Print the value of the variable and the pointer
	println("Value of x:", x) // Output: Value of x: 10
	// TODO: *p คือ dereference pointer p เพื่อให้ได้ value ของ x
	println("Value of *p:", *p) // Output: Value of *p: 10
	// TODO: &x คือ address ของ x
	println("Address of x:", &x) // Output: Address of x: <memory address>
	// TODO: p คือ address ของ x
	println("Address stored in p:", p) // Output: Address stored in p: <memory address>
	// TODO: &p คือ address ของ p
	println("Address of p:", &p) // Output: Address of p: <memory address>

	// Change the value of x using the pointer
	changeValue(p)

	println("Value of x after change:", x)   // Output: Value of x after change: 100
	println("Value of *p after change:", *p) // Output: Value of *p after change: 100

	changeValue2(x)
	println("Value of x after change2:", x)   // Output: Value of x after change2: 100
	println("Value of *p after change2:", *p) // Output: Value of *p after change2: 100

}

func printValueAndAddress(x *int) {
	// Print the value and address of the variable
	fmt.Println("Value of x:", x)    // Dereference pointer to get value
	fmt.Println("Address of x:", &x) // Print address stored in pointer
	fmt.Printf("Type of x: %T\n", x) // Print type of pointer => *int
}

func PointerExample3() {
	fmt.Println(" ")
	fmt.Println("Pointer Example 3")

	x1 := 10
	p1 := &x1                            // p1 is a pointer to x1, it holds the address of x1
	println("Value of x1:", x1)          // Output: Value of x1
	println("Value of *p1:", *p1)        // Output: Value of *p1: 10
	println("Address of x1:", &x1)       // Output: Address of x1: <memory address>
	println("Address stored in p1:", p1) // Output: Address stored in p1: <memory address>
	println("Address of p1:", &p1)       // Output: Address of p1: <memory address>
	printValueAndAddress(p1)             // Call the function to print value and address
	p2 := &p1
	println("Value of p2:", p2)        // Output: Value of p2: <memory address>
	println("Address of p2:", &p2)     // Output: Address of p2: <memory address>
	println("Value of *p2:", *p2)      // Output: Value of *p2: <memory address>
	println("Value of **p2:", **p2)    // Output: Value of **p2: 10
	fmt.Printf("Type of p2: %T\n", p2) // Output: Type of p2: **int (pointer to pointer to int)
}

type Person struct {
	Name string
	Age  int
}

func ChangeName1(p Person, name string) {
	// Change the name of the person
	p.Name = name
	fmt.Println("Inside ChangeName1:", p.Name) // Output: Inside ChangeName1: Alice
}

func ChangeName2(p *Person, name string) {
	// Change the name of the person using pointer
	// TODO: ถ้าเป็น struct ไม่ต้องมี *
	p.Name = name
	// (*p).Name = name
	fmt.Println("Inside ChangeName2:", p.Name) // Output: Inside ChangeName2: Alice
}

func PointerExample4() {
	fmt.Println(" ")
	fmt.Println("Pointer Example 4")

	// Create a variable of type Person
	person := Person{Name: "John", Age: 30}

	fmt.Println("Before ChangeName1:", person.Name) // Output: Before ChangeName1: John
	ChangeName1(person, "Alice")                    // Call ChangeName1 with person (pass by value)
	fmt.Println("After ChangeName1:", person.Name)  // Output: After ChangeName1: John

	ChangeName2(&person, "Alice")                  // Call ChangeName2 with address of person (pass by pointer)
	fmt.Println("After ChangeName2:", person.Name) // Output: After ChangeName2: Alice

	person2 := &Person{Name: "Bob", Age: 25}
	fmt.Println("Person2:", person2)                         // Output: Person2: &{Bob 25}
	fmt.Println("Before ChangeName2:", person2.Name)         // Output: Before ChangeName2: Bob
	ChangeName2(person2, "Charlie")                          // Call ChangeName2 with address of person2 (pass by pointer)
	fmt.Println("After ChangeName2:", person2.Name)          // Output: After ChangeName2: Charlie
	fmt.Println("Address of person2:", person2)              // Output: Address of person2: <memory address>
	fmt.Println("Address of person2.Name:", &person2.Name)   // Output: Address of person2.Name: <memory address>
	fmt.Println("Address of person2.Age:", &person2.Age)     // Output: Address of person2.Age: <memory address>
	fmt.Printf("Type of person2: %T\n", person2)             // Output: Type of person2: *pointer.Person (pointer to Person struct)
	fmt.Printf("Type of person2.Name: %T\n", person2.Name)   // Output: Type of person2.Name: string
	fmt.Printf("Type of person2.Age: %T\n", person2.Age)     // Output: Type of person2.Age: int
	fmt.Printf("Type of &person2: %T\n", &person2)           // Output: Type of &person2: **pointer.Person (pointer to pointer to Person struct)
	fmt.Printf("Type of &person2.Name: %T\n", &person2.Name) // Output: Type of &person2.Name: *string (pointer to string)

	// TODO: new keyword to create a new Pointer to Person
	c := new(Person)                                   // Create a new Person using new
	fmt.Printf("c : %#v\n", c)                         // Output: c: &{ 0} (pointer to Person struct with zero values)
	ChangeName2(c, "David")                            // Call ChangeName1 with c (pass by valu)
	fmt.Println("After ChangeName2 with new:", c.Name) // Output: After ChangeName2 with new: David
	fmt.Printf("c : %#v\n", c)                         // Output: c: &{David 0} (pointer to Person struct with updated name)

	var zeroValueOfPointer *int
	// TODO: zero value of pointer is nil
	fmt.Println("Zero value of pointer:", zeroValueOfPointer) // Output: Zero value of pointer: <nil>
	if zeroValueOfPointer == nil {
		fmt.Println("Pointer is nil") // Output: Pointer is nil
	} else {
		fmt.Println("Pointer is not nil")
	}
}

type movie struct {
	Title        string
	Year         int
	Rating       float64
	Vote         []float64
	IsSupperHero bool
}

func (m *movie) addVote(vote float64) {
	// Add a vote to the movie
	m.Vote = append(m.Vote, vote)
	fmt.Printf("Added vote: %.1f to movie: %s\n", vote, m.Title)
}

// TODO: vote ...float64 คือ variadic parameter
func (m movie) addVote2(vote ...float64) []float64 {
	// Add a vote to the movie (without pointer receiver)
	fmt.Printf("Added vote: %.1f to movie: %s\n", vote, m.Title)
	// TODO: m.Vote เป็น slice ของ float64 ก็เลยต้องใช้ ... เพื่อให้สามารถ append หลายค่าได้
	return append(m.Vote, vote...)
}

func (m movie) addVote3(voteList []float64, vote ...float64) []float64 {
	// Add a vote to the movie (without pointer receiver)
	fmt.Printf("Added vote: %.1f to movie: %s\n", vote, m.Title)
	return append(voteList, vote...) // Append the votes to the provided slice
}

func (m *movie) calculateRating() float64 {
	// Calculate the average rating of the movie
	if len(m.Vote) == 0 {
		fmt.Println("No votes available to calculate rating.")
		return 0.0
	}

	sum := 0.0
	for _, v := range m.Vote {
		sum += v
	}
	return sum / float64(len(m.Vote))
}

func (m *movie) ChangeTitle(newTitle string) {
	// Change the title of the movie
	m.Title = newTitle
	fmt.Printf("Changed movie title to: %s\n", m.Title)
}

func WorkshopPointer() {
	fmt.Println(" ")
	fmt.Println("Workshop Pointer")

	movie1 := movie{
		Title:        "The Dark Knight",
		Year:         2008,
		Rating:       9.0,
		Vote:         []float64{},
		IsSupperHero: true}
	fmt.Printf("before Movie: %s, Year: %d, Rating: %.1f, Votes: %v, IsSupperHero: %t\n", movie1.Title, movie1.Year, movie1.Rating, movie1.Vote, movie1.IsSupperHero)
	movie1.addVote(9.0) // Add a vote to the movie
	movie1.addVote(9.0) // Add a vote to the movie
	movie1.addVote(9.2) // Add a vote to the movie
	fmt.Printf("after Movie: %s, Year: %d, Rating: %.1f, Votes: %v, IsSupperHero: %t\n", movie1.Title, movie1.Year, movie1.Rating, movie1.Vote, movie1.IsSupperHero)

	rating := movie1.calculateRating() // Calculate the average rating of the movie
	fmt.Printf("Average rating of movie %s: %.1f\n", movie1.Title, rating)

	movie1.ChangeTitle("The Dark Knight Rises") // Change the title of the movie
	fmt.Printf("after ChangeTitle Movie: %s, Year: %d, Rating: %.1f, Votes: %v, IsSupperHero: %t\n", movie1.Title, movie1.Year, movie1.Rating, movie1.Vote, movie1.IsSupperHero)

	movie1.Vote = movie1.addVote2(9.5, 1, 1, 1, 1) // Add a vote to the movie (without pointer receiver)
	fmt.Printf("after addVote2 Movie: %s, Year: %d, Rating: %.1f, Votes: %v, IsSupperHero: %t\n", movie1.Title, movie1.Year, movie1.Rating, movie1.Vote, movie1.IsSupperHero)

	movie1.Vote = movie1.addVote3([]float64{1, 1}, 5, 1, 1, 1, 1) // Add a vote to the movie (without pointer receiver)
	fmt.Printf("after addVote3 Movie: %s, Year: %d, Rating: %.1f, Votes: %v, IsSupperHero: %t\n", movie1.Title, movie1.Year, movie1.Rating, movie1.Vote, movie1.IsSupperHero)
}
