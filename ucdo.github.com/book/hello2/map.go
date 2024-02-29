package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	//m1()
	//m := map[string]int{"a": 0}
	//mm := map[string]int{"b": 1}
	//fmt.Printf("%t", mapEqual(&m, &mm))
	//dedup()
	charCount()
}

func m1() {
	m := map[string]int{
		"xxx1": 1,
		"xxx2": 2,
		"xxx3": 3,
		"xxx4": 4,
	}

	for i, i2 := range m {
		fmt.Println(i, " => ", i2)
	}

}

func mapEqual(m *map[string]int, mm *map[string]int) bool {
	if len(*m) != len(*mm) {
		return false
	}

	for s, i := range *m {
		if v, ok := (*mm)[s]; !ok || v != i {
			return false
		}
	}

	return true
}

func dedup() {
	s := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !s[line] {
			s[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

func charCount() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int // array
	invalid := 0
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		counts[r]++
		utflen[r]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
