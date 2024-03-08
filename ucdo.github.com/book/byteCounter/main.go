package main

import (
	"bufio"
	"fmt"
	"strings"
)

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	*b += ByteCounter(len(p)) // type convert
	return len(p), nil
}

func main() {
	var b ByteCounter
	b.Write([]byte("hello"))
	fmt.Println(b)
	b = 0
	name := "xxx"
	fmt.Fprintf(&b, "hello,%s", name)
	fmt.Println(b)
	fmt.Println("=========")
	fmt.Println(*(b.ByteCounterPer("xx11 z")))

}

func (b *ByteCounter) ByteCounterPer(s string) *map[string]int {
	// 先定义 bufio.NewScanner
	c := make(map[string]int)
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if c[scanner.Text()] == 0 {
			c[scanner.Text()] = 1
			continue
		}
		c[scanner.Text()]++
	}

	return &c
}
