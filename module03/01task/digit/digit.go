package digit

import (
	"fmt"
	"gopackages/wordz"
)

func init() {
	fmt.Println("This is random number")
}

func Digit() {
	digit := wordz.Random()
	fmt.Println(digit)
}
