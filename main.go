package main

import (
	"LeetGo/groupTotals"
	"fmt"
)

func main() {
	//1
	/*nums := []int{3, 2, 4}
	target := 6
	ans := Two_Sum.TwoSum(nums, target)
	fmt.Println(ans)*/

	//1108
	//defanging_an_ip_address.DefangIPaddr("127.12.0.1")

	/*arr := []int{12,345,2,6,7896}
	fmt.Println(arr)
	find_Numbers_with_Even_Number_of_Digits.FindNumbers(arr)*/

	/*n := 234
	subtract_the_product_and_sum_of_digits_of_an_integer.SubtractProductAndSum(n)*/

	/*s := strconv.FormatInt(5, 2)
	fmt.Println(s)*/

	/*arr := []int{2, 2, 2, 2, 4, 1}
	result := sumMulti.SumMultiplier(arr)
	fmt.Println(result)*/

	strArr := []string{"X:-1", "Y:1", "X:-4", "B:3", "X:5"}
	fmt.Println(groupTotals.GroupTotals(strArr))
}
