package main

import (
	"LeetGo/1108_defanging_an_ip_address"
	"LeetGo/1_Two_Sum"
	"fmt"
)

func main() {
	//1
	nums := []int{3, 2, 4}
	target := 6
	ans := Two_Sum.TwoSum(nums, target)
	fmt.Println(ans)

	//1108
	defanging_an_ip_address.DefangIPaddr("127.12.0.1")
}
