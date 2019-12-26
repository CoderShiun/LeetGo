package subtract_the_product_and_sum_of_digits_of_an_integer

import (
	"strconv"
)

func SubtractProductAndSum(n int) int {
	var product = 1
	var sum int

	strN := strconv.Itoa(n)
	for _, v := range strN {
		num, _ := strconv.Atoi(string(v))
		product *= num
		sum += num
	}

	return product - sum
}