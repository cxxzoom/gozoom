## 第一天

### 环境的搭建

直接在官网下载并安装就行了

### 工作区

### bin
存放编译之后的可执行文件
### pkg
包文件
### src
开发代码都放这
#### GOPATH
安装时会自动指定到安装目录下
这里自己开发，要改成工作区的地方，跟bin,pkg,src同级
#### windows下的
还需要把工作区的bin加到环境变量里面的path里面

### 编译
go build 
如果要编译到linux上需要执行：
SET CGO_ENABLE=0
SET GOOS=linux
SET GOARCH=amd64
要在编译成windows上的可执行文件
SET GOOS=windows

### 变量声明
var varName type  
短标签命名法不能在函数外部使用

### 赋值
已经声明的变量，不能再用短标签重复声明
不能重复声明变量

### const
```
cosnt(
    a  = 1
    b // 1
    c // 1
)
```
1. iota只能在常量里面使用
2. iota在consts里面，每新增一行变量申明就会加一
    ```
    const (
        a = iota // 0
        b = 100 // 100
        c = iota // 2
        d //3
    )
    const (
        a,j = iota + 1,iota + 2 // 0 + 1 , 0 + 2
        b,c = // 1 + 1 , 1 + 2
        d,e // 2 + 1, 2 + 2
    )
    ```
3. 遇到const会初始化为0
4. const 如果不写，就默认和上一行一样
   
### int
1. uint8 = 2^8  = byte 类型
2. int16 = C short
3. int64 = C long
4. int会根据机器的位数动态调整 32/64

### 在本地启一个godoc的文档服务
godoc -http=:8080  
然后通过127.0.0.1:8080

## 编码问题
中文是按照 UTF8存储的，占3个字节  
光是用fori遍历会出问题  


## for range
会自动处理，如果是utf8，则按utf8输出  
比如遍历 s:= "hello 中文"  
fori 遍历会乱码，但是for range不会


## 类型转换
go里面没有隐式转换，只有强制类型转换
### 修改一个字符串
必须要转换成[]byte或者[]rune类型才能操作  
看看代码  
```
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
```