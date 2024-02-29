package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "国家"
	s1 := s[1:4]

	fmt.Println(len(s), s1)

	t := "123"
	xx := t
	t = "0000" + t
	t += "0000"

	println(t)
	println(xx)

	fmt.Println(basename("abc"))
	fmt.Println(basename("a/b/c"))
	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename("a/b/c.d.go"))

	fmt.Println(addSp("12345"))
	fmt.Println(int2String([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(i2s([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(addSp2("123456.0", false))
	fmt.Printf("%s\n", addSp2("123456.0000", true))
	fmt.Println(isReverse("1xxx", "xxx1"))

	fmt.Println(itoa())
}

func basename(s string) string {
	// 为什么要逆序？因为是找最后一个，只需要找到最后一个/的位置并截取
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// 这里也是，因为可能会有多个.,比如a.b.c.d.e.go
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func addSp(s string) string {
	n := len(s)
	if n < 3 {
		return s
	}
	return addSp(s[:n-3]) + "," + s[n-3:]
}

func addSp2(s string, unsign bool) string {
	var buf bytes.Buffer
	left := s
	last := strings.LastIndex(s, ".")
	if last > 0 {
		left = s[:last]
	}

	// 这里要写成 len / 3 > 0?
	lens := len(left)
	ss := ""
	for {
		tmp := ""
		if lens-3 > 0 {
			// 这里就用3分隔
			tmp = left[lens-3:]
			lens = max(0, lens-3)
			if lens > 0 {
				tmp = "," + tmp
			}
		} else {
			tmp = left[:lens]
			lens = 0
		}

		ss = tmp + ss
		if lens == 0 {
			break
		}
	}

	if last != -1 {
		ss += s[last:]
	}

	if unsign {
		ss = "-" + ss
	}

	fmt.Fprintf(&buf, "%c", ss)

	ss = buf.String()
	return ss
}

func int2String(val []int) string {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range val {
		if i > 0 {
			buf.WriteString(", ")
		}

		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteString("]")
	return buf.String()
}

func i2s(i []int) string {
	var s string
	s += "["
	for i2, i3 := range i {
		if i2 > 0 {
			s += ", "
		}

		s += strconv.Itoa(i3)
	}
	s += "]"

	return s
}

// 判断s1 和 s2是否互为倒叙
func isReverse(s1 string, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	j := len(s2) - 1
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[j] {
			return false
		}
		j--
	}

	return true
}

func itoa() string {
	x := 123
	fmt.Println(fmt.Sprintf("%d", x))
	fmt.Println(strconv.Itoa(x))
	fmt.Println(fmt.Sprintf("%b %#b", x, x))
	return ""
}

func testS() {
	s := "1234567890123"
	ttt := s[4:7]
	fmt.Println(len(ttt))
}
