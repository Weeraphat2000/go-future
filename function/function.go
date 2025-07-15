package function1

import (
	"fmt"
)

// if start with uppercase letter, it is exported and can be accessed from other packages

// function name starts with uppercase letter for public access
func FunctionAddTwoNumbers(a int, b int) int {
	return a + b
}

// function name starts with lowercase letter for this package only
func add(a int, b int) int {
	return a + b
}

func minus(a int, b int) int {
	return a - b
}

func equal(a int, b int) bool {
	return a == b
}

// function name starts with uppercase letter for public access
func FunctionAddAndSubAndEqual(a int, b int) (int, int, bool) {
	return minus(a, b), add(a, b), equal(a, b)
}

func FunctionAddMinusEqual(a, b int) (int, int, bool) {
	return functionAddNumberInFunction2(a, b), functionMinusNumberInFunction2(a, b), functionEqualNumberInFunction2(a, b)
}

func IntPlusFloat(a int, b float64) {
	fmt.Println("The sum of int and float is:", intPlusFloat(a, b))
}

// ประกาศตัวแปรที่จะคืนค่าจากฟังก์ชันแล้ว แล้วเมื่อเจอคำสั่ง return จะคืนค่าล่าสุดกลับไป
func Split(a, b int) (x int, y float64) {
	x = a + b
	y = float64(a) + float64(b)
	return
}

// function แบบนี้จะต้องส่ง function signature เข้ามาเป็นพารามิเตอร์ และให้ถูกต้องด้วย
func Compute(fu func(int, int) int, a, b int) int {
	return fu(a, b)
}

// higher-order function
// function ที่รับฟังก์ชันเป็นพารามิเตอร์ หรือคืนค่าฟังก์ชันเป็นผลลัพธ์
// และสามารถเก็บ สถานะ(state) ได้
func Adder1() (add func(int) int, subtract func(int) int) {
	// state variable
	sum := 0
	add = func(x int) int {
		sum += x
		return sum
	}
	subtract = func(x int) int {
		sum -= x
		return sum
	}
	return
}

// function แบบนี้จะ return signature ของฟังก์ชันที่ต้องการคืนค่าไปด้วย
func Adder2() (func(int) int, func(int) int) {
	sum := 0
	add := func(x int) int {
		sum += x
		return sum
	}
	subtract := func(x int) int {
		sum -= x
		return sum
	}
	return add, subtract
}

func Emote(ratings float64) string {
	if ratings >= 4.5 {
		return "😊"
	} else if ratings >= 3.5 {
		return "🙂"
	} else if ratings >= 2.5 {
		return "😐"
	} else if ratings >= 1.5 {
		return "🙁"
	}
	return "😞"
}
