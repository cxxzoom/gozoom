package main
import "fmt"

func main(){
	fmt.Println("hello world")
	var (
		a string
		b int
	)

	a = "xxx"
	b = 7
	var xxx string = "123"
	var xxx1 = "xxx1"
	fmt.Println(a,b,xxx,xxx1)
	xxxmm,b,err := funcCall()
	fmt.Println(xxxmm,b,err)
}

func funcCall()(string,int,error){
	return "xxx",1,nil
}