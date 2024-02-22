package main

import "fmt"

func main() {
	s1()
	s2()
	s3()
}

func s1() {
	a := [13]string{1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "10", 11: "11", 12: "12"}
	s := a[4:7]
	s2 := a[6:9]
	fmt.Println(len(s), cap(s))
	fmt.Println(s)
	fmt.Println(s2)

	//
	for _, v := range s {
		for _, v2 := range s2 {
			if v2 == v {
				fmt.Printf("has same value: %s\n", v)
			}
		}
	}

	fmt.Println(s[:9])
	fmt.Println(s)
	s[0] = "9999"
	//fmt.Println(s)
	fmt.Println(a)
	s = append(s, "00", "01", "02")
	fmt.Println(s)
	fmt.Println(a)
}

func s2() {
	fmt.Println("\n\nin s2...")
	a := [4]int{1, 2, 3, 4}
	f := a[:]
	s := f[:]

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(a)
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}

func s3() {
	var s []int
	fmt.Printf("%t\n", s == nil) //true
	s = nil
	fmt.Printf("%t\n", s == nil) //true
	s = []int(nil)
	fmt.Printf("%t\n", s == nil) //true
	s = []int{}
	fmt.Printf("%t\n", s == nil) //false
}
