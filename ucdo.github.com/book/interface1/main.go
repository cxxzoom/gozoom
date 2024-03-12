package main

import "fmt"

func main() {

}

func sqlExec(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x)
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return sqlQuoteString(x)
	default:
		panic(fmt.Sprintf("unexpected type %T : %v", x, x))
	}
}

// deal string...
func sqlQuoteString(x string) string {
	return x
}
