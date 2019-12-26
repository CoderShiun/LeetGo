package find_Numbers_with_Even_Number_of_Digits

import (
	"strconv"
)

func FindNumbers(nums []int) int {
	var total int

	for _, v := range nums {
		len := len(strconv.Itoa(v))

		if len % 2 == 0 {
			total ++
		}
	}

	return total
}