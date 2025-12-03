package util

import (
	"fmt"
	"math/big"
)

func Run_factorial() {
	// Example of big.NewInt usage
	num1 := big.NewInt(5)  // Creates big.Int with value 5
	num2 := big.NewInt(10) // Creates big.Int with value 10

	result := new(big.Int) // Creates a new big.Int with value 0
	result.Mul(num1, num2) // Multiplies num1 and num2

	fmt.Printf("num1: %v (type: %T)\n", num1, num1)
	fmt.Printf("num2: %v (type: %T)\n", num2, num2)
	fmt.Printf("5 x 10 = %v\n\n", result)

	// Now let's calculate some factorials
	fmt.Printf("Factorial 85: %v\n", bigFactorial(85))
	fmt.Printf("Factorial 86: %v\n", bigFactorial(86))
}

func bigFactorial(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	result := big.NewInt(n)
	return result.Mul(result, bigFactorial(n-1))
}
