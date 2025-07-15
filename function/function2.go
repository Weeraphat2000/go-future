package function1

// function name starts with lowercase letter for this package only
func functionAddNumberInFunction2(a int, b int) int {
	return a + b
}

func functionMinusNumberInFunction2(a int, b int) int {
	return a - b
}

func functionEqualNumberInFunction2(a int, b int) bool {
	return a == b
}

func intPlusFloat(a int, b float64) float64 {
	return float64(a) + b
}
