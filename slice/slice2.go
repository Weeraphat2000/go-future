package slice

import "fmt"

func show1(tage string, s []string) {
	fmt.Printf("%s: %#v\n", tage, s)
}

func SliceAdvancedExample() {

	//
	fmt.Println(" ")
	fmt.Println("Slice Advanced Example")
	skills := []string{"Go", "Python", "JavaScript", "C++", "Java"}

	// TODO: copy by address some address
	s1 := skills[1:] // Slice from index 1 to 2 (exclusive of 3)
	s2 := skills[:2] // Slice from the start to index 3 (exclusive of 3)

	show1("s1", s1) // {"Python", "JavaScript", "C++", "Java"}
	show1("s2", s2) // {"Go", "Python"}

	skills[1] = "Weeraphat"                // Update an element in the slice
	show1("skills after update 1", skills) // {"Go", "Weeraphat", "JavaScript", "C++", "Java"}
	show1("s1 after update 1", s1)         // {"Weeraphat", "JavaScript", "C++", "Java"}
	show1("s2 after update 1", s2)         // {"Go", "Weeraphat"}

	// TODO: append to the new slice เลยทำให้ s1, s2 ไม่เปลี่ยนแปลง
	skills = append(skills, "Ruby", "PHP") // Append multiple elements to the slice
	show1("skills after append", skills)   // {"Go", "Python", "JavaScript", "C++", "Java", "Ruby", "PHP"}
	show1("s1 after append", s1)           // {"Python", "JavaScript", "C++", "Java"}
	show1("s2 after append", s2)           // {"Go", "Python"}
	skills[1] = "Golang"                   // Update an element in the slice
	show1("skills after update 2", skills) // {"Go", "Golang", "JavaScript", "C++", "Java", "Ruby", "PHP"}
	show1("s1 after update 2", s1)         // {"Weeraphat", "JavaScript", "C++", "Java"}
	show1("s2 after update 2", s2)         // {"Go", "Weeraphat"}
}
