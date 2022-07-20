package main

import "fmt"

func AddPositiveNumber(num1 int, num2 int, num3 int) int {

	if num1 <= 0 || num2 <= 0 || num3 <= 0 {
		return -1
	}

	fmt.Printf("sum of %d + %d + %d=  %d", num1, num2, num3, num1+num2+num3)
	return num1 + num2 + num3
}

func main() {

	AddPositiveNumber(2, 3, 4)

}
