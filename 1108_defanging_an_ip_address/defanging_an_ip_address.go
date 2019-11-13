package defanging_an_ip_address

import (
	"fmt"
	"strings"
)

func DefangIPaddr(address string) string {
	slice := strings.Split(address, "")
	fmt.Println(slice,len(slice))
	for x, y := range slice {
		if y == "." {
			slice[x] = "[.]"
			fmt.Println(slice)
		}
	}
	result := strings.Join(slice, "")
	fmt.Println(result)

	return result
}