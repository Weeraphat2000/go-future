package json

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Address  string // TODO: ถ้าไม่ใช้ tag json จะใช้ชื่อ field ตรง ๆ ใน JSON (ตัวพิมพ์เล็กหรืือพิมพ์ใหญ่ได้หมด)
	IsActive bool   `json:"is_active"` // TODO: ถ้าใช้ tag json ชื่อ field จะถูกเปลี่ยนเป็น is_active ใน JSON (ตัวพิมพ์เล็กหรืือพิมพ์ใหญ่ได้หมด)
}

type Course struct {
	Name  string // TODO: ถ้าขึ้นต้นด้วยตัวพิมพ์ใหญ่ จะทำให้เป็น public field
	level string // TODO: ถ้าขึ้นต้นด้วยตัวพิมพ์เล็ก package json จะไม่สามารถเข้าถึงได้
	Views int    `json:"viewssssssss"` // TODO: ถ้าใช้ tag json จะทำให้สามารถกำหนดชื่อ field ใน JSON ได้
}

func JsonExample() {
	// Example of a simple JSON structure
	course := Course{Name: "Go Programming", level: "Beginner", Views: 1000}

	// TODO: json.Marshal คือ การแปลงข้อมูลจาก struct ไปเป็น JSON string
	jsonData, err := json.Marshal(course)
	fmt.Printf("Type: %T\n", jsonData)       // []uint8
	fmt.Printf("byte slice: %v\n", jsonData) // [123 34 78 97 109 101 34 58 34 71 111 32 80 114 111 103 114 97 109 109 105 110 103 34 44 34 76 101 118 101 108 34 58 34 66 101 103 105 110 110 101 114, ...]
	fmt.Printf("String: %s\n", jsonData)     // {"Name":"Go Programming","Level":"Beginner","Views":1000}
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
	} else {
		// TODO: string(jsonData) คือ การแปลง byte slice กลับไปเป็น string
		fmt.Printf("JSON data: %#v\n", string(jsonData))
		fmt.Printf("JSON data: %s\n", jsonData)
	}

	var courseData Course
	// TODO: json.Unmarshal คือ การแปลงข้อมูลจาก JSON string กลับไปเป็น struct
	err = json.Unmarshal(jsonData, &courseData)
	if err != nil {
		fmt.Println("Error sconverting from JSON:", err)
	} else {
		fmt.Printf("Converted back to struct: %#v\n", courseData)
		fmt.Printf("Course Name: %s, Level: %s, Views: %d\n", courseData.Name, courseData.level, courseData.Views)
	}

	//
	//
	var personJson = `{"name": "John Doe", "AGE": 3000, "AddreSS": "123 Main St", "is_actiVE": true}`
	var personData Person
	// TODO: json.Unmarshal คือ การแปลงข้อมูลจาก JSON string กลับไปเป็น struct
	err = json.Unmarshal([]byte(personJson), &personData)
	if err != nil {
		fmt.Println("Error converting from JSON:", err)
	} else {
		fmt.Printf("Converted Person: %#v\n", personData)
		fmt.Printf("Name: %s, Age: %d\n", personData.Name, personData.Age)
	}

	// TODO: json.Marshal คือ การแปลงข้อมูลจาก struct ไปเป็น JSON string
	result, err := json.Marshal(personData)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
	} else {
		fmt.Printf("JSON data: %s\n", result)
		fmt.Printf("JSON data: %#v\n", string(result))
	}
}

type Movie struct {
	Title       string  `json:"titleeeee"`
	Year        int     `json:"year"`
	Rating      float64 `json:"rating"`
	Genre       string  `json:"genre"`
	IsSuperHero bool    `json:"is_super_hero"`
}

func JsonWorkShop() {

	fmt.Println(" ")
	fmt.Println("JSON Workshop")

	body := `[{"id": "tt01", "titleeeee": "The Dark Knight", "year": 2008, "rating": 9.0, "genre": "Action", "is_super_hero": true}, {"id": "tt02", "titleeeee": "Inception", "YEAR": 2010, "rating": 8.8, "genre": "Sci-Fi", "is_super_hero": false}]`

	var movies []Movie
	// TODO: json.Unmarshal คือ การแปลงข้อมูลจาก JSON string กลับไปเป็น struct
	err := json.Unmarshal([]byte(body), &movies)
	if err != nil {
		fmt.Println("Error converting from JSON:", err)
	} else {
		for _, m := range movies {
			fmt.Printf("Title: %s, Year: %d, Rating: %.1f, Genre: %s, Is Super Hero: %t\n", m.Title, m.Year, m.Rating, m.Genre, m.IsSuperHero)
		}
	}

}
