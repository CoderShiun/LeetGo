package subtract_the_product_and_sum_of_digits_of_an_integer

import (
	"strconv"
)

func SubtractProductAndSum(n int) int {
	var product = 1
	var sum int

	for _, v := range strconv.Itoa(n) {
		num, _ := strconv.Atoi(string(v))
		product *= num
		sum += num
	}

	return product - sum
}