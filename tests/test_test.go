// go test ./tests สำหรับการทดสอบ เฉพาะฟังก์ชันในแพ็คเกจ tests
// go test ./... สำหรับการทดสอบทั้งหมดในโปรเจกต์
// go test -v ./tests สำหรับการทดสอบแบบละเอียด เพื่อให้เห็นทุก case เลย

package tests

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	} else {
		fmt.Println("TestAdd passed")
	}
}

func TestAddNegative(t *testing.T) {
	result := Add(-2, -3)
	expected := -5

	if result != expected {
		t.Errorf("Add(-2, -3) = %d; want %d", result, expected)
	} else {
		fmt.Println("TestAddNegative passed")
	}
}
