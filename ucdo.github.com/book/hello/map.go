package main

import "fmt"

func main() {
	m := map[int]string{}
	f := make(map[int]string)
	if m == nil {
		println("m is nil")
	} else {
		println("m is not nil")
	}

	if f == nil {
		println("f is nil")
	} else {
		println("f is not nil")
	}

	for i2, s := range m {
		fmt.Printf("%d => %s", i2, s)
	}

	for i2, s := range f {
		fmt.Printf("%d => %s", i2, s)
	}

	v, ok := m[1]
	println(v, ok)
	m[22] = "hello"
	println(m[22])
	for i2, s := range m {
		fmt.Printf("%d %s", i2, s)
	}

	delete(m, 23)
	println(m[22])

	println(m[234])
	key := 123
	maps := map[int]string{}
	println(maps[key])

	map1 := make(map[string]int)
	println(map1["string"])

	map2 := make(map[int]map[string]int)
	map2[1] = make(map[string]int)
	map2[1]["tool"] = 1
	map2[1] = make(map[string]int)
	println(map2[2] == nil)

	s := t1()
	println(s)
}

func t1() int {
	return 1
}
