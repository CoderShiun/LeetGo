package groupTotals

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
Have the function GroupTotals(strArr) read in the strArr parameter containing key:value pairs where the key is a string
and the value is an integer. Your program should return a string with new key:value pairs
separated by a comma such that each key appears only once with the total values summed up.

For example: if strArr is ["B:-1", "A:1", "B:3", "A:5"] then your program should return the string A:6,B:2.

Your final output string should return the keys in alphabetical order.
Exclude keys that have a value of 0 after being summed up.
*/
func GroupTotals(strArr []string) string {
	res := make(map[string]int)

	for _, v := range strArr {
		sp := strings.Split(v, ":")
		str := sp[0]
		preVal := res[sp[0]]
		val, _ := strconv.Atoi(sp[1])
		res[str] = val + preVal
	}

	var result []string
	for k, v := range res {
		if v != 0 {
			result = append(result, fmt.Sprintf("%s:%d", k, v))
		}
	}

	sort.Strings(result)

	str := strings.Join(result, ",")

	return str
}