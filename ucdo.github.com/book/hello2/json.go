package main

import (
	"encoding/json"
	"fmt"
)

type movie struct {
	name   string
	author []string
	Year   string `json:"released,omitempty"`
	Color  string `json:"mCool"`
}

func main() {
	j1()
}

func j1() {
	j := []movie{
		{name: "星际牛仔", author: []string{"111", "?????"}, Year: "199x", Color: "red"},
		{name: "星际牛仔1", author: []string{"111", "?????"}, Year: "199x", Color: "red"},
		{name: "星际牛仔2", author: []string{"111", "?????"}, Year: "199x", Color: "red"},
		{name: "星际牛仔3", author: []string{"111", "?????"}, Year: "199x", Color: "red"},
		{name: "星际牛仔4", author: []string{"111", "?????"}, Year: "199x", Color: "red"},
		{name: "星际牛仔5", author: []string{"111", "?????"}, Year: "199x", Color: "red"},
		{name: "星际牛仔6", author: []string{"111", "?????"}, Year: "199x", Color: "red"},
	}

	r, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", r)

	data, err := json.MarshalIndent(j, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data)

	data1, err := json.MarshalIndent([]movie{
		{name: "xxx", Year: "xxx", Color: "x"},
		{name: "xxx", Color: "x"},
	}, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data1)

	j2 := make([]movie, 0)
	err = json.Unmarshal(data1, &j2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", j2)
}
