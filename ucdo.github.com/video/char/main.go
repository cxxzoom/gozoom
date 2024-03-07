package main
import "fmt"
func main(){
	s := "golang"
	c := 'c' // ASCII 码下占一个字节=8位
	// 一个rune = 一个UTF8字符 比如可以存中文
	fmt.Println(s,c)

	c2 := "hello 中文"
	fmt.Println(c2,len(c2))
	for i := 0; i <len(c2); i++ {
		fmt.Println(i,c2[i])
	}

	// for range 是按rune类型遍历的
	for _, v := range c2 {
		fmt.Printf("%c\n",v)
	}

	// 修改字符串： 要先转换成[]rune或者[]byte类型才能操作
	// []byte强制类型转换
	s1 := "xxxx"
	// 这里其实也只是一个copy
	byteArr := []byte(s1)
	fmt.Println(byteArr)
	byteArr[0] = 'X'
	fmt.Println(byteArr)
	fmt.Println(string(byteArr))
	fmt.Println(s1)

	sx := "hello"
	sxx := []byte(sx)
	j := len(sxx) - 1
	for i := 0; i < len(sxx); i++ {
		if i == j || i > j {
			break
		}
		sxx[i],sxx[j] = sxx[j],sxx[i]
		j--
	}
	fmt.Println(string(sxx))

	for i := 0; i < 4; i++ {
		for j := 0; j < 8; j++ {
			if i == 2{
				break
			}
		}
		fmt.Println(i)
	}
}