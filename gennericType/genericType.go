package gennerictype

import (
	"fmt"
)

func AddInt(a, b int) int {
	return a + b
}

func AddFloat64(a, b float64) float64 {
	return a + b
}

func AddGeneric[T int | float64](a, b T) T {
	return a + b
}

func AddGenericMixed[T int | float64](a T, b int) T {
	return a + T(b)
}

func AddGenericIntFloat64String[T int | float64 | string](a T, b T) T {
	return a + b
}

func TwoGeneric[T, U any](a T, b U) T {
	fmt.Printf("a: %v, b: %v, type of a: %T, type of b: %T\n", a, b, a, b)

	switch v := any(a).(type) {
	case int:
		if bVal, ok := any(b).(int); ok {
			result := v + bVal
			return any(result).(T)
		}
	case float64:
		if bVal, ok := any(b).(float64); ok {
			result := v + bVal
			return any(result).(T)
		}
	default:
		fmt.Println("Unsupported types for addition")
	}
	var zero T
	return zero
}

type Number interface {
	int | float64
}

func min[T Number](a, b T) T {
	// aa := 29
	// bb := 30.66
	// fmt.Printf("aa: %d, bb: %f, aa > bb: %t\n", aa, bb, float64(aa) > bb) // TODO: แค่จะสื่อว่า int กับ float64 สามารถไม่เปรียบเทียบกันได้

	fmt.Printf("a: %v, b: %v, type of a: %T, type of b: %T\n", a, b, a, b)

	if a < b {
		return a
	}
	return b
}

func PrintAddResults() {
	fmt.Println("AddInt(2, 3):", AddInt(2, 3))
	fmt.Println("AddFloat64(2.5, 3.5):", AddFloat64(2.5, 3.5))
	fmt.Printf("AddGeneric[int](2, 3): %d, type: %T\n", AddGeneric(2, 3), AddGeneric(2, 3))
	fmt.Printf("AddGeneric[float64](2.5, 3.5): %f, type: %T\n", AddGeneric(2.5, 3.5), AddGeneric(2.5, 3.5))
	fmt.Printf("AddGeneric[float64](2.5, 3): %f, type: %T\n", AddGeneric(2.5, 3), AddGeneric(2.5, 3))

	fmt.Printf("AddGenericMixed[float64](2.5, 3): %f, type: %T\n", AddGenericMixed(2.5, 3), AddGenericMixed(2.5, 3))
	fmt.Printf("AddGenericMixed[int](2, 3): %d, type: %T\n", AddGenericMixed(2, 3), AddGenericMixed(2, 3))

	fmt.Printf("int + float64: %f, type: %T\n", AddGenericIntFloat64String(2, 3.5), AddGenericIntFloat64String(2, 3.5))
	fmt.Printf("float64 + int: %f, type: %T\n", AddGenericIntFloat64String(2.5, 3), AddGenericIntFloat64String(2.5, 3))
	// fmt.Printf("string + int: %s, type: %T\n", AddGenericIntFloat64String("Hello", 20), AddGenericIntFloat64String("Hello", 20)) // ไม่สามารถทำงานได้เพราะ string ไม่สามารถใช้ + กับ int ได้

	fmt.Println(" ")
	fmt.Println("TwoGeneric Results:")
	fmt.Printf("TwoGeneric[int, int](2, 3): %d, type: %T\n", TwoGeneric(2, 3), TwoGeneric(2, 3))
	fmt.Printf("TwoGeneric[float64, float64](2.5, 3.5): %f, type: %T\n", TwoGeneric(2.5, 3.5), TwoGeneric(2.5, 3.5))

	fmt.Printf("TwoGeneric[int, float64](2, 3.5): %d, type: %T\n", TwoGeneric(2, 3.5), TwoGeneric(2, 3.5))
	fmt.Printf("TwoGeneric[float64, int](2.5, 3): %f, type: %T\n", TwoGeneric(2.5, 3), TwoGeneric(2.5, 3))

	fmt.Println(" ")
	fmt.Println("Min Results:")
	fmt.Printf("min[int](2, 3): %d, type: %T\n", min(2, 3), min(2, 3))
	fmt.Printf("min[float64](2.5, 3.5): %f, type: %T\n", min(2.5, 3.5), min(2.5, 3.5))
	fmt.Printf("min[int](2, 3.5): %f, type: %T\n", min(2, 3.5), min(2, 3.5))
	fmt.Printf("min[float64](2.5, 3): %f, type: %T\n", min(2.5, 3), min(2.5, 3))

}
