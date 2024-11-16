package main

import (
	"fmt"
	"strconv"
	"strings"
)

func complexNumberMultiply(num1, num2 string) string {
	parseComplex := func(num string) (int, int) {
		parts := strings.Split(num[:len(num)-1], "+")
		real, _ := strconv.Atoi(parts[0])
		imaginary, _ := strconv.Atoi(parts[1])
		return real, imaginary
	}

	real1, imag1 := parseComplex(num1)
	real2, imag2 := parseComplex(num2)

	realResult := real1*real2 - imag1*imag2
	imagResult := real1*imag2 + imag1*real2

	return fmt.Sprintf("%d+%di", realResult, imagResult)
}

func main() {
	var num1, num2 string
	fmt.Println("Enter the first complex number (in the form real+imaginaryi):")
	fmt.Scan(&num1)
	fmt.Println("Enter the second complex number (in the form real+imaginaryi):")
	fmt.Scan(&num2)

	result := complexNumberMultiply(num1, num2)
	fmt.Printf("The result of the multiplication is: %s\n", result)
}
