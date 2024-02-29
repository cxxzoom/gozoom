package main

import "fmt"

type F float64 // 这里相当于是定义了一种类型，就像C 的 typedef一样
type C float64

const (
	AbsoluteZeroC C = -273.15 // 绝对零度
	FreezingC     C = 0       // 结冰点温度
	BoilingC      C = 100     // 沸水温度 // 沸点
)

func F2C(f F) C {
	return C((f - 32) * 5 / 9)
}

func C2F(c C) F {
	return F(c*9/5 + 32)
}

func (c C) String() string {
	return fmt.Sprintf("%g°C", c)
}

func main() {
	c := F2C(212.0)
	fmt.Println(c.String()) // "100°C"
}
