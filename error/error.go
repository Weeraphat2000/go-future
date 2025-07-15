package errorss

import (
	"errors"
	"fmt"
)

// TODO: zero value of error type is nil
func doSomething() error {
	// Simulate an operation that might fail
	// In a real application, this could be a database operation, file I/O, etc.
	return nil // or return an error if something goes wrong
}

type CustomError struct{}

// TODO: ถ้าจะสร้าง error ใหม่ ให้สร้าง struct ที่ method Error() string ด้วย
func (e CustomError) Error() string {
	return "this is a custom error"
}

func DivideWithCustomError(a int, b int) (int, error) {
	// This function divides two integers and returns a custom error if division by zero is attempted.
	if b == 0 {
		return 0, CustomError{}
	}
	return a / b, nil
}

func Divide(a int, b int) (int, error) {
	// This function divides two integers and returns an error if division by zero is attempted.
	if b == 0 {

		// TODO: return error with message ใช้ errors.New() หรือ fmt.Errorf() ก็ได้
		return 0, errors.New("division by zero is not allowed")
		// return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}

func ErrorExample() {
	// Example of error handling in Go
	// This function demonstrates how to handle errors using the built-in error type.
	// It simulates a function that might return an error and shows how to check for it.

	err := doSomething()
	if err != nil {
		// Handle the error
		println("An error occurred:", err.Error())
	} else {
		println("Operation completed successfully")
	}

	if result, err := Divide(10, 0); err != nil {
		// Handle the division by zero error
		println("Error:", err.Error())
	} else {
		println("Result of division:", result)
	}

	if result, err := DivideWithCustomError(10, 0); err != nil {
		// Handle the custom error
		println("Custom Error:", err.Error())
	} else {
		println("Result of division with custom error:", result)
	}

	test := CustomError{}
	fmt.Printf("CustomError: %#v\n", test.Error())
}
