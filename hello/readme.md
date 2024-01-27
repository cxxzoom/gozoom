# GO

## 最基础
```
1. package: 包
2. base: 每个可执行程序有一个main函数
3. base: 优先级 init() > main
4. base: 首字母大写即外部可见
5. base: 首字母小写即包内可见
6. string: 字符串拼接 直接用  + 
7. base: 变量申明 var identifier type
8. base: 变量申明 identifier := value // 好像可以自动类型推断
9. base: 只定义变量不复制，各有各的初始值
10. const: 常量 ： const identi [type] = value
11. iota : 初始0，一直 ++，有值跳过，无值赋值; 嗯这里iota的值，怪得很
12. auto: 自动类型推断
13. base: 无三目运算符
14. 字符串是定长的
15. 数组是定长的
16. 数组申明方式 var a = []int{1,3,}  b := []int{1}
17. 数组申明长度但是不赋值，会以数组类型初始化
18. &a  表示取地址
19. 指针  var identity *type
20. 不能对指针随便赋值，只能放地址
21. 指针也是有地址的
22. 如果数组指定了类型或者长度，任意一项不匹配都不行
23. 闭包
24. base: 数组是定长的，切片是随意的
25. [size]type: 指针数组，存放指针的数组
26. pointer: 多重指针，即指向指针的指针。 var a **int // 二重指针
27. struct: s结构体指针  var c *Book = &book
28. pointer: 自动类型推断的指针 c := &book
29. func: 函数接收参数时，指定为 func c(c *Book){} // 这里要结构体类型
30. slice: 是数组的抽象
31. slice: 不指定大小的数组就是切片
32. slice: var inden []type
33. slice: var s []type = make([]type,len) 
34. slice: slice := make([]type,len)
35. make([]T, length, capacity) // 类型，长度，容量
36. slice: a := []int{1,2,3} s := a[start:end] // end > cap(a) 会报错
37. slice: s := []int{} var s []int ,分别用 s == nil 判断，前一个不为nil 
38. printf: %x 以16进制输出， %d以10进制输出
39. copy: copy(target, origin)（只针对slice），当target的cap小于origin，只copy target.cap个
40. range: for k,v = range f{} // f可以是array,slice,channel,map,
41. range: for _,v = range "你好"
42. map: map是引用类型，一个地方修改会影响到所有地方
43. map: 跟slice一样，会自动扩容
44. map: 始终建议用make来初始化
45. map: 杰哥： 一般不用多维，多维一般用struct
46. map: 多维数组看有没有初始化可以用 == nil来判断，为了程序的安全性
47. strconv: 类型转换
48. 类型转换的时候一定要处理error，不然我估计要崩溃
49. interface: 接口太自由了
50. interface: 任何其他类型实现了这些方法就是实现了这个接口。必须全部实现
51. goroutine: 所有goroutine共享一个地址空间
52. goroutine: 异步的，普通方法是同步的
53. goroutine: go func_name(...param)
53. channel： c := make(chan int)
54. channel： 在make的时候可以设置是否有缓冲区
55. channel： 如果无缓冲区，那么读写都是原子性的。意思会阻塞其他协程写
56. channel： 无缓冲就会导致竞争，这没问题吧？所以高并发怎么玩？
57. channel： 有缓冲区，如果没写满，就会一直写，不会竞争
58. channel： 有缓冲区，如果写满了，就要等有人来读，不然一直等，可能 deadlock
59. 逃逸： 逃逸是指生命周期从短生命周期提升到长生命周期。比如将func里面局部变量的地址赋给全局变量
60. 生命周期：生命周期长的变量是分配在堆上的；相反就是分配在栈上
```

### 并发
#### 关键字 go chan
#### 怎么玩的？
```go
package main
import "fmt"

func main(){
	// goroutine简单输出
	go test(1) // 1
	test(2) // 2
}

func test(a int){
	println(a)
}

```
这里需要注意的是，因为协程是异步的，如果 `2` 也用 `go test(2)`，就会导致main退出。然后就不执行

#### 无缓冲chan

#### [带缓冲channel死锁演示](./chanDL.go)


## 关键字
```` 
break	    default         func        interface   select
case	    defer           go          map         struct
chan	    else            goto        package     switch
const	    fallthrough	    if	        range       type
continue    for	            import      return      var
````

## 预定义标识符
```
append  bool    byte    cap     close   complex     complex64   complex128  uint16
copy	false   float32 float64 imag    int         int8        int16       uint32
int32   int64	iota	len     make	new     nil     panic	uint64
print	println real	recover string  true	uint	uint8	uintptr
```

## 执行/编译
``` 
1. go run file.go
2. go build file.go
3. ./file.exe
```