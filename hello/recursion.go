package main

func main() {
	ss := recursion(2, 2)
	println(ss)
}

func recursion(i int, n int) int {
	if n > 0 {
		ret := i * recursion(i, n-1)
		return ret
	}

	return 1
}
