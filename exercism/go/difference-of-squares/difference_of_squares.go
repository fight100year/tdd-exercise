package diffsquares

// SquareOfSum calc the square of sum
func SquareOfSum(count int) int {
	sum := 0
	for i := 1; i <= count; i++ {
		sum += i
	}

	return sum * sum
}

// SumOfSquares calc the sum of squares
func SumOfSquares(count int) int {
	result := 0
	for i := 1; i <= count; i++ {
		result += i * i
	}

	return result
}

// Difference is diff of sum and square
func Difference(count int) int {
	return SquareOfSum(count) - SumOfSquares(count)
}
