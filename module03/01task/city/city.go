package city

import (
	"fmt"
	"gopackages/wordz"
)

// var delarr = wordz.Words
// var index = 5
// func remove(delarr []string, index int) []int {
// 	return append(delarr[:index],delarr[index+1]...)
// }

func City() {
	wordz.Words = []string{"Moscow", "London", "Ashgabat", "New York", "Tashkent"}
	city := wordz.Random()
	fmt.Println(city)
}
