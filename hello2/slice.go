package main

import "fmt"

func main() {
	s1()
	s2()
	s3()
	s4()
	s := []string{"one", "", "three", "", "tt"}
	//fmt.Println("\n", noneempty(s))
	//f := noneempty(s)
	//fmt.Println(s)
	//fmt.Println(f)

	//fmt.Println(noneempty2(s))

	f := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//rewriteReverse(&f)
	//fmt.Println(f)
	Rotate(&f, 2)
	fmt.Println(f)
	s = []string{"1", "", "2", "", "3", "7", "4", "5"}
	insituNoneepmty(&s)
	fmt.Println(s)
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

func s4() {
	var runes []rune

	for _, v := range "hello,你好" {
		runes = append(runes, v)
	}

	fmt.Printf("%q", runes)
}

func appendInt(x []int, y ...int) []int {
	var z []int
	lens := len(x) + len(y)
	if lens < cap(x) {
		z = x[:lens]
	} else {
		// new lens,new cap
		caps := lens
		if caps < 2*lens {
			caps = 2 * lens
		}
		z = make([]int, lens, caps)
		copy(z, x)
	}

	copy(z[len(x):], y)
	return z
}

func noneempty(s []string) []string {
	i := 0
	for _, v := range s {
		if v != "" {
			s[i] = v
			i++
		}
	}
	return s[:i]
}

func noneempty2(s []string) []string {
	x := make([]string, 0, len(s))
	for _, s5 := range s {
		if s5 != "" {
			x = append(x, s5)
		}
	}

	return x
}

func rewriteReverse(p *[]int) {
	for i, j := 0, len(*p)-1; i < j; i, j = i+1, j-1 {
		(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
	}
}

func Rotate(p *[]int, place int) {
	if place == 0 {
		return
	}

	left := make([]int, 0, place)
	right := make([]int, 0, len((*p))-place)

	for i, v := range *p {
		if i < place {
			right = append(right, v)
		} else {
			left = append(left, v)
		}
	}

	*p = append(left, right...)
}

func insituNoneepmty(p *[]string) {
	if len(*p) < 0 {
		return
	}

	i, j := 0, 1

	for {
		if j >= len(*p) {
			return
		}

		if (*p)[j] != "" && (*p)[i] == "" {
			(*p)[i] = (*p)[j]
			i++
		}

		if (*p)[i] == "" {
			i++
		}

		j++
	}
}
