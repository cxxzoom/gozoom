package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	m1()
	m := map[string]int{"a": 0}
	mm := map[string]int{"b": 1}
	fmt.Printf("%t", mapEqual(&m, &mm))
	dedup()
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
