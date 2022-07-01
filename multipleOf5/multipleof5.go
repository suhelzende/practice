package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// N = 1598 => result = 198

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	newNum := 0
	found := 0
	multiplier := 1
	for num != 0 {
		rem := num % 10
		if (rem != -5 && rem != 5) || found == 1 {
			newNum = (rem * multiplier) + newNum
			multiplier *= 10
		}

		if rem == 5 || rem == -5 {
			found = 1
		}
		num /= 10
	}
	fmt.Println(newNum)
}

// 15859 =>  1859   1589
