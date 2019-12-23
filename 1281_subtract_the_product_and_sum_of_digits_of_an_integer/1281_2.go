package subtract_the_product_and_sum_of_digits_of_an_integer

import (
	"strconv"
	"strings"
)

func SubtractProductAndSum2(n int) int {
	var product = 1
	var sum int

	num := strconv.Itoa(n)
	numarr := strings.Split(num, "")

	for _, v := range numarr {
		newN, _ := strconv.Atoi(v)
		product *= newN
		sum += newN
	}

	return product - sum
}