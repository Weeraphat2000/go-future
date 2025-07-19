// go test ./...   สำหรับการทดสอบทั้งหมดในโปรเจกต์
// go test -v ./...   สำหรับการทดสอบแบบละเอียด เพื่อให้เห็นทุก case ทุก file
// go test -v .   สำหรับการทดสอบในแพ็คเกจปัจจุบัน

// ชื่อ file ต้องขึ้นต้นด้วย test_ หรือ _test.go
// ชื่อฟังก์ชันทดสอบต้องขึ้นต้นด้วย Test และมีพารามิเตอร์ *testing.T

// ใช้ t.Errorf() สำหรับรายงานข้อผิดพลาดในการทดสอบ
// ใช้ t.Fatal() หรือ t.Fatalf() สำหรับรายงานข้อผิดพลาดที่ทำให้การทดสอบหยุดทำงาน
// ใช้ t.Run() สำหรับการทดสอบย่อยภายในฟังก์ชันทดสอบหลัก
// ใช้ t.Skip() หรือ t.Skipf() สำหรับข้ามการทดสอบบางส่วน
// ใช้ t.Parallel() สำหรับการทดสอบแบบขนาน
// ใช้ t.Cleanup() สำหรับการทำความสะอาดหลังการทดสอบ
// ใช้ t.Log() หรือ t.Logf() สำหรับการบันทึกข้อมูลระหว่างการทดสอบ
// ใช้ t.Helper() เพื่อระบุว่าฟังก์ชันนี้เป็นฟังก์ชันช่วยเหลือในการทดสอบ
// ใช้ t.Name() เพื่อรับชื่อของฟังก์ชันทดสอบปัจจุบัน
// ใช้ t.Failed() เพื่อตรวจสอบว่าการทดสอบล้มเหลวหรือไม่
// ใช้ t.SkipNow() เพื่อข้ามการทดสอบทันที
// ใช้ t.Skipf() เพื่อข้ามการทดสอบพร้อมกับข้อความที่กำหนด
// ใช้ t.Setenv() เพื่อกำหนดตัวแปรสภาพแวดแวดล้อมสำหรับการทดสอบ
// ใช้ t.TempDir() เพื่อสร้างไดเรกทอรีชั่วคราวสำหรับการทดสอบ
// ใช้ t.TempFile() เพื่อสร้างไฟล์ชั่วคราวสำหรับการทดสอบ
// ใช้ t.Error() หรือ t.Errorf() สำหรับรายงานข้อผิดพลาดที่ไม่ทำให้การทดสอบหยุดทำงาน

package main

import (
	"fmt"
	"testing"
)

func TestAddForPackageMain(t *testing.T) {
	t.Run("Success add", func(t *testing.T) {
		result := add(2, 3)
		expected := 5

		if result != expected {
			t.Errorf("Add(2, 3) = %d; want %d", result, expected)
		} else {
			fmt.Println("TestAdd passed")
		}
	})

	t.Run("Negative add", func(t *testing.T) {
		result := add(-2, -3)
		expected := -5

		if result != expected {
			t.Errorf("Add(-2, -3) = %d; want %d", result, expected)
		} else {
			fmt.Println("TestAddNegative passed")
		}
	})
}
