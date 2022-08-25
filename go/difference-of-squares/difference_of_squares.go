package diffsquares

// SquareOfSum returns the square of the sum of the first n natural numbers.
func SquareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

func SumOfSquares(n int) int {
	// sum := 0
	// for i := 1; i <= n; i++ {
	// 	sum += i * i
	// }
	sum := (n * (n + 1)) >> 1
	return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
