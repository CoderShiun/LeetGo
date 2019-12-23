package subtract_the_product_and_sum_of_digits_of_an_integer

import (
	"fmt"
	"strconv"
)

func SubtractProductAndSum(n int) int {
	product := product(n)
	sum := sum(n)

	return product - sum
}

func product (n int) int {
	var result = 1

	for _, v := range strconv.Itoa(n) {

		num, _ := strconv.Atoi(string(v))
		result *= num
	}


	return result
}

func sum (n int) int {
	var result int

	for _, v := range strconv.Itoa(n) {
		num, _ := strconv.Atoi(string(v))
		result += num
	}

	fmt.Println(result)

	return result
}