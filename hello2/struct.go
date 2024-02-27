package main

import "fmt"

type employee struct {
	ID       int
	Name     string
	Salary   int
	Position string
}

var e employee

func main() {
	println(e.ID)
	f := &e
	f.ID = 7777
	id := &f.ID
	*id = 9999
	position := &f.Position
	*position = "sdfksjdjfk"
	println(e.ID)
	println(e.Position)
	fmt.Println(getId(&e).ID)
	getId(&e).ID = 9
	fmt.Println(e.ID)
}

func getId(id *employee) *employee {
	id.ID = 0001
	return id
}
