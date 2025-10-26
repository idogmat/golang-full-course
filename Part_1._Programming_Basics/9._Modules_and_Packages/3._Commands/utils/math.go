package utils

func Multiply(a, b int) int {
	InfoLog("Multiply")
	return a * b
}

func Add(a, b int) int {
	InfoLog("Add")
	return a + b
}

func Subtract(a, b int) int {
	InfoLog("Subtract")
	return a - b
}

func Divide(a, b int) int {
	InfoLog("Divide")
	if b == 0 {
		return 0
	}
	return a / b
}