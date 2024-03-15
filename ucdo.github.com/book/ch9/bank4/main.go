package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var x, y int
	go func() {
		x = 1
		fmt.Println("y:", y)
	}()

	go func() {
		y = 1
		fmt.Println("x:", x)
	}()

	time.Sleep(3)
}

var icons map[string]string

func loadIcons() {
	icons = map[string]string{
		"spades.png":   Icon1("spades.png"),
		"hearts.png":   Icon1("hearts.png"),
		"diamonds.png": Icon1("diamonds.png"),
		"clubs.png":    Icon1("clubs.png"),
	}
}

func Icon1(s string) string {
	return s
}

var loadOnce sync.Once

func Icon(name string) string {

	loadOnce.Do(loadIcons)
	return icons[name]
}
