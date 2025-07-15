package changevalue

import (
	"fmt"
	"strconv"
)

func stringToInt(s string) (int, error) {
	// TODO: strconv.Atoi is used to convert a string to an int. if can not convert, it returns an error.
	// TODO: strconv.Atoi => string to int
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("error converting string to int: %w", err)
	}
	return i, nil
}

func intToString(i int) string {
	// TODO: strconv.Itoa is used to convert an int to a string. if can not convert, it returns an error.
	// TODO: strconv.Atoi is used to convert a string to an int. if can not convert, it returns an error.
	// TODO: strconv.Itoa => int to string
	i, err := strconv.Atoi(strconv.Itoa(i))
	if err != nil {
		return fmt.Sprintf("Error converting int to string: %s", err)
	}
	return strconv.Itoa(i)
}

func intToString2(i int) string {
	return strconv.Itoa(i)
}

func intToString3(i any) string {
	// test:= i.(type) // can not use .(type) outside a type switch

	switch v := i.(type) {
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case string:
		// strconv.Atoi => string to int
		_, err := strconv.Atoi(v)
		if err == nil {
			return v // If it's a valid integer string, return it as is
		}
		return fmt.Sprintf("error: %s", err)
	default:
		return fmt.Sprintf("Unsupported type: %T", v)
	}
}

func floatToInt(f float64) int {
	return int(f)
}

func intToFloat(i int) float64 {
	return float64(i)
}

func floatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func ChangeValue() {

	result1, error1 := stringToInt("123")
	fmt.Println("Converted string to int:", result1, "Type:", fmt.Sprintf("%T", result1), "error1:", error1)

	result2, error2 := stringToInt("abc")
	fmt.Println("Converted string to int:", result2, "Type:", fmt.Sprintf("%T", result2), "error2:", error2)

	result3 := intToString(456)
	fmt.Println("Converted int to string:", result3, "Type:", fmt.Sprintf("%T", result3))

	result3a := intToString2(456)
	fmt.Println("Converted int to string (alternative):", result3a, "Type:", fmt.Sprintf("%T", result3a))

	result3b := intToString3(456)
	fmt.Println("Converted int to string (generic):", result3b, "Type:", fmt.Sprintf("%T", result3b))
	result3c := intToString3(789.12)
	fmt.Println("Converted float to string (generic):", result3c, "Type:", fmt.Sprintf("%T", result3c))
	result3d := intToString3("test")
	fmt.Println("Converted string to string (generic):", result3d, "Type:", fmt.Sprintf("%T", result3d))
	result3e := intToString3("123")
	fmt.Println("Converted string to string (generic):", result3e, "Type:", fmt.Sprintf("%T", result3e))

	result4 := floatToInt(789.12)
	fmt.Println("Converted float to int:", result4, "Type:", fmt.Sprintf("%T", result4))

	result5 := intToFloat(456)
	fmt.Println("Converted int to float:", result5, "Type:", fmt.Sprintf("%T", result5))

	result6 := floatToString(789.12)
	fmt.Println("Converted float to string:", result6, "Type:", fmt.Sprintf("%T", result6))
}
