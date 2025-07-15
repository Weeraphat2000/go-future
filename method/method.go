package method1

import (
	"fmt"
	"strconv"
)

type MethodExample struct {
	Name string
	Age  int
}

type Movie struct {
	Title  string
	Year   int
	Rating float64
	Genre  []string
}

func (m Movie) Info() string {
	genres := ""
	for i, genre := range m.Genre {
		if i > 0 {
			genres += ", "
		}
		genres += genre
	}
	return fmt.Sprintf("Title: %s, Year: %d, Rating: %.1f, Genre: %s", m.Title, m.Year, m.Rating, genres)
}

// create methods Greet for returning a greeting message when called method Greet
// TODO: จะทำให้ struct MethodExample มี method Greet ที่สามารถเรียกใช้ได้
// copy by value
// TODO: Greet คือ method ของ struct MethodExample (receiver)
func (m MethodExample) Greet(name string) string {
	m.Name = name
	return "Hello, my name is " + m.Name + " and I am " + strconv.Itoa(m.Age) + " years old."
}

// Method to update the age of the person
// TODO: use pointer receiver for performance optimization
// change by reference
// TODO: UpdateAge คือ method ของ struct MethodExample (receiver)
func (m *MethodExample) UpdateAge(newAge int) {
	m.Age = newAge
}

// TODO: GetDetails คือ method ของ struct MethodExample (receiver)
func (m MethodExample) GetDetails(address string, gender string) string {
	return "Name: " + m.Name + ", Age: " + strconv.Itoa(m.Age) + ", Address: " + address + "gender: " + gender
}

func Methods() {

	person := MethodExample{Name: "John", Age: 30}
	fmt.Println(person.Greet("Johnnnnnnnnnnnnnnnnnnnnnnn")) // Hello, my name is Johnnnnnnnnnnnnnnnnnnnnnnn and I am 30 years old.
	fmt.Println("Name :", person.Name)                      // Name : John
	person.UpdateAge(300)
	fmt.Println(person.GetDetails("123 Main St", "Boy")) // Name: John, Age: 300, Address: 123 Main Stgender: Boy
	fmt.Println("Methods completed successfully.")

	fmt.Println(" ")
	fmt.Println("Movie Details:")
	movie := Movie{Title: "Inception", Year: 2010, Rating: 8.8, Genre: []string{"Sci-Fi", "Action"}}
	fmt.Println(movie.Info())
	movie.Rating = 9.0 // Update rating
	fmt.Println("Updated Movie Details:", movie.Info())
}
