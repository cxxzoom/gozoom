package main

import (
	"fmt"
)

func main() {
	// fmt.Print(sum(1, 2, 3, 4, 5))
	s := "xxx"
	s = join2(s, "11", "22", "33", "44")
	fmt.Print(s)

}

func sum(val ...int) int {
	sum := 0
	for _, v := range val {
		sum += v
	}

	return sum
}

func max1(val ...int) int {
	if len(val) < 1 {
		return 0
	}

	max := val[0]
	for v := range val {
		if v > max {
			max = v
		}
	}

	return max
}

func join2(s string, j ...string) string {
	for _, v := range j {
		s += v
	}

	return s
}
