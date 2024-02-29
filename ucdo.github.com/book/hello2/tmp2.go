package main

import (
	"fmt"
	"os"
)

func main() {

}

func test(name string) string {
	f, err := os.Open(name)
	if err != nil {
		return err.Error()
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			// log...
		}
	}(f)

	//
	return ""
}

func test2(name string) {
	if _, err := os.Open(name); err != nil { // 这里本本来是f的，但是没使用，所以先注释了
		fmt.Println(err.Error())
	}

	// f.Close() 会报错

}
