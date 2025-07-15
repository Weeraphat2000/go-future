package movies

import (
	"fmt"
	"strconv"
)

// TODO: ถ้า package ถูกเรียกใช้งาน จะทำ init() ถูกเรียกใช้งานก่อน
func init() {
	fmt.Println("init movies is called")
}

// TODO: ถ้าชื่อขึ้นต้นด้วยตัวพิมพ์ใหญ่ จะทำให้สามารถเรียกใช้ได้จาก package อื่น ๆ ถ้าขึ้นต้นด้วยตัวพิมพ์เล็ก จะทำให้สามารถเรียกใช้ได้เฉพาะภายใน package นี้เท่านั้น
func ReviewMovie(title string, rating int) string {
	if rating < 1 || rating > 5 {
		return "Invalid rating. Please provide a rating between 1 and 5."
	}
	return "You reviewed the movie '" + title + "' with a rating of " + strconv.Itoa(rating) + " stars."
}

func FindMovieById(id string) string {
	switch id {
	case "1":
		return "The Shawshank Redemption"
	case "2":
		return "The Godfather"
	case "3":
		return "The Dark Knight"
	default:
		return "Movie not found with ID: " + id
	}
}
