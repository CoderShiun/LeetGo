package Two_Sum

func TwoSum(nums []int, target int) []int {
	//s := len(nums) - 1
	var array []int
	//idx := 0
	s := len(nums)
	//block:
	for x := 0; x < s; x++ {
		for y, num := range nums {
			if x == y {
				continue
			}
			if nums[x]+num == target {
				array = append(array, x, y)
				return array
			} else {
				continue
			}
		}
	}
	return nil
}
