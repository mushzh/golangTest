package calculate

func Sum(num1, num2 int) int {
	return num1 + num2
}
func Sub(num1, num2 int) int {
	return num1 - num2
}

func calculate(a int, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
}

func calculate2(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}
