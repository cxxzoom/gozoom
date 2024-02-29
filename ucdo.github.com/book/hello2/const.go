package main

import (
	"fmt"
	"time"
)

func main() {
	T()
}

func T() {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %v", noDelay, noDelay)
	s := 1
	fmt.Printf("%v", s)

	const (
		a = 1
		b
		c
		d = 2
		e
	)

	fmt.Println("\n", a, b, c, d, e)

	var f float64 = 3 + 0i
	f = 2
	f = 1e123
	f = 'a'
	fmt.Println(f)
}

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Flags uint

const (
	FlagUp           Flags = 1 << iota
	FlagBroadcast          // supports broadcast access capability
	FlagLoopback           // is a loopback interface
	FlagPointToPoint       // belongs to a point-to-point link
	FlagMulticast          // supports multicast access capability
)

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776             (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424    (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)

const (
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB
	EB = 1000 * PB
	ZB = 1000 * EB
	YB = 1000 * ZB
)
