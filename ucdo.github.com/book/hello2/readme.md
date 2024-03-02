## package

``` 
1. 类似于其他语言的模块或者库的概念
2. 封装、模块化、代码重用
3. 单独编译
4. 每个包都对应一个独立的名字空间
5. 首字母大写即可导出
6. 每个包在解决依赖的前提下，以导入声明的顺序初始化，每个包只会被初始化一次
```

## 作用域

``` 
1. 声明语句的作用域是指源代码中 可以有效使用这个变量名的范围
2. 不要把作用域和生命周期混为一谈
3. 它属于编译是的属性
4. 当然，有全局的，也有局部的
5. if 的作用域就在if语句块内，在外面就没了
6. tmp2.go 演示了if的作用域
7. 
```

## 生命周期

``` 
1. 一个变量的声明周期指的是程序在运行时变量存在的有效时间段
2. 在声明周期内可以被其他程序引用
3. 运行时概念

```

## 编译器

``` 
1. 当编译器遇到一个名字应用时，它会对其定义进行查找，查找的过程从最内层的词法域向全局的作用域进行。
2. 
```

## 基础类型

1. 0开头的8进制
2. 0x或者0X开头的16进制

### 整型 int

```
1. int8 int16 int32 int64
2. uint8 uint16 uint32 uint64
3. byte 等价于 uint8 
4. rune 等价于 uint32
5. uintptr 这个是跟C交互会用到
6. int8 = 2^(8-1) ~ 2^(8-1)-1 
7. uint8 = 0 ~ 2^8 - 1 = 0 ~ 255
```

### 浮点型 float

``` 
1. float32 float64
2. 在math包里
```

### 复数 complex

### 布尔型 bool

### 字符串 string

``` 
1. 一个不可改变得字节序列
2. s[i] i >= len(s) : panic
3. 非ASCII字符的UTF8编码，会占用多个字节，比如 s := "国家" len(s) // 6
4. s1 := s[0,3] // 国  s2 := [1,4] //  ���
5. 4个重要包： bytes,strings,strconv,unicode
6. bytes.Buffer
7. strconv.Itoa // int to string
8. s[i:j] 字符串切片生成新的字符串
```

### 字节 []byte

``` 
1. 字节切片生成新的 []byte
```

### 常量

``` 
1. 在编译期完成
2. 
```

## 复合类型

### 数组

``` 
1. 定长
2. 如果只指定长度，不进行操作，那么默认为0值
3. 定义的是时候要指定长度： a := [2]type{} 或者 s := [...]type{1,2,3}
4. a := [3]type{1,2,3} 清空数组  a = [3]type{}
```

### slice

``` 
1. 类似数组，但是变长
2. 一个slice由指针，长度，以及容量构成
3. 指针指向第一个slice元素对应的底层数组元素的地址（但是不一定就是数组的第一个元素，比如我从3开始切片）
4. 切片操作 s[i:j] 切片将有用 j-i个元素 0 <= i <= j <= cap(s)
5. 容量： a := [13]int{} s := a[4,7] fmt.Println(len(s),cap(s)) // 3 9
6. 真神奇： 继续使用上面的s， fmt.Println(s[:9]) // 这里将输出 a[4:]里面对应的内容
7. 如果切片操作超出cap(s)的上限将导致一个panic异常
8. s[0] = 123 // 这里修改了，然后 数组 a 对应的值也会被修改，具体位置要看slice的第一个元素的指针位置
9. 向函数传递slice意味着，底层元素也可能会被修改
10. 不能像 array 一样直接比较， 好像有个bytes.Equal，但是只能比较 []byte 类型的，其他类型需要自己写
11. var s []int  s == nil // true 
12. s = []int(nil) s== nil // true
13. s = nil s == nil // true 
14. s = []int{}  s == nil // false
15. 11-14没什么用
16. 要判断slice是否为空，只需要 len(s) == 0 来判断
17. make([]T,len,cap) []T{} make的优点是可以
18. slice appendint1 //老版本 ： 必须先检测slice底层数组是否有足够的容量来保存新添加的元素
19. slice appentint1 版本： 如果有足够空间的话，直接扩展slice（依然在原有的底层数组之上），将新添加的y元素复制到新扩展的空间，并返回slice。因此，输入的x和输出的z共享相同的底层数组。
20. 如果没有足够的增长空间的话，appendInt函数则会先分配一个足够大的slice用于保存新的结果，先将输入的x复制到新的空间，然后添加y元素。结果z和输入的x引用的将是不同的底层数组。
21. 内置的append函数更加复杂；所以不清楚新的slice和原始slice是否引用的是相同的底层数组空间。
22. 同样也不能确定在原先的slice上操作会不会影响新的slice
23. ... 变长参数slice
24. 有坑！ 要是操作slice，并返回slice的切片的时候，一定要注意切片切的地方；要么就直接make一个，然后append进去再返回
```

### map

``` 
1. 无序kv集合，检索、更新、删除都是常数的时间复杂度
2. 哈希表的引用
3. 创建 
    1. age := make(map[string]int) 
    2. age := map[string]int{
            "name1":1,
            "name2":2,
        }
    3. age := map[string]int{}
4. 跟php的array类似？ 但是 age["keyyyy"] 如果key不存在，不会报错
5. 禁止对map取地址，取也会编译报错；因为map可能会随着元素增加而重新分配更大的存储空间，从而导致之前的地址失效
6. 遍历的顺序是随机的
7. 好像，好像不能直接对map进行排序
8. v,ok := age[key]; !ok{/*值不存在*/}
9. map 之间也不能进行比较，除非nil 或者你就要手动写代码去比较
10. map的key居然可以是结构体，吐了
11. map的key用struct也无所谓，因为他又不是引用类型，改了也无所谓
```

跟PHParray的区别

``` 
1. age["keyyyy"] 如果key不存在，不会报错
2. 每次遍历的顺序，php是一致的；而go的map不确定
3. 好像不能直接排序，要把map的key放在slice里面，用 sort.Strings(s)对slice里面的值进行排序
4. 当前结构体可以嵌套其他结构体
5. 两个匿名的结构体成员吧不能有同名的成员，不然会报错
6. 在包外部，不可导出的成员，也不能通过匿名方式访问和初始化
7. 
```

## 结构体 struct

``` 
1. 可以是复合类型
2. 里面可以引用struct本身
3. 初始化可以记住struct申明的顺序结构，然后挨个初始化
4. 初始化也可以通过申明的名字来初始化
5. 未初始化的成员，会被默认初始化成对应类型的零值
6. 不能初始化其他包里面可导出struct的不可导出成员
7. type s struct{
        a int
    }
    ...
    pp := &s{a:111} // 初始化并把地址赋予给pp
    // 下面是同样的操作
    pp := new(s)
    *pp = s{a:777}
    ...
```

### json

``` 
1. 发送和接收结构化数据的解析，类似的还有xml，protobuf
2. 标准库 encoding/json
3. 只有可导出的才会被编码
4. type f struct{
        a string 
        B string `json:"BNickname"`
    } 
5. json.Marshal(T) // 默认格式
6. json.MarshalIndent(T,"","    ") // 格式化
7. 结构体成员的json里面人如果带了 omitempty ，当改成员为空时，则不导出
```

## 函数 func

``` 
1. 指针，slice(切片)、map、function、channel等类型，实参可能会由于函数的间接引用被修改。
2. 任何进行I/O操作的函数都会面临出现错误的可能，只有没有经验的程序员才会相信读写操作不会失败
3. 居然也可以像PHP一样，把方法名赋予变量，然后通过 v() 调用
```

## 匿名函数

``` 
1. 捕获迭代变量：这里很容易出问题
2. 看个1 的例子
    var rmdirs []func()
    for _, d := range tempDirs() {
        dir := d // NOTE: necessary!
        os.MkdirAll(dir, 0755) // creates parent directories too
        rmdirs = append(rmdirs, func() {
            os.RemoveAll(dir)
        })
    }
    // ...do some work…
    for _, rmdir := range rmdirs {
        rmdir() // clean up
    }

    代码的第三行为什么要这么操作？ 因为后面的匿名函数记录的是
    运行时变量的内存地址，如果直接赋值d的话，d就是最后一次迭代的值了，
    那么所有的删除操作都变成对 最后一个文件的删除操作了

3. 为了解决2上的问题，通常引入一个同名的局部变量
    for _, dir := range tempDirs() {
        dir := dir // declares inner dir, initialized to outer dir
        // ...
    }

    go programming language 写到
4. 不仅仅是for range有这个问题，fori，go，defer都有这个问题
5. go或者defer会等循环结束之后再执行，所以4
6. 函数内的匿名函数可以访问包括返回值在内的所有变量


```

## 可变参数 -> 参数的数量不确定

```
1. sum(val ...int) 这里可以传多个int类型的参数
2. 调用者会创建一个匿名数组来装接收值
3. 怎么向sum里面传 int类型的slice？
4. s := []int{1,2,3,4} sum(s...)
```

## defer

```
1. 好用的defer
2. 一个日志记录的例子
    func bigSlowOperation(){
        defer trace("some msg")()
    }

    func trace(msg string) func(){
        timeStart := time.Now()
        log.Printf("%s is start...\n",msg)
        return func(){
            log.Printf("%s is ending, cost time %s",msg,time.Since(timeStart))
        }
    }
3. defer语句中的函数会在return语句更新返回值变量后再执行
4. 特别注意： 在fori，forr中defer会在最后才执行
5. 
```

## panic

```
1. 运行时错误会导致panic，比如数组越界，空指针引用
2. 发生panic时，会导致程序中断-》并输出堆栈跟踪信息-》
   通常，我们不需要再次运行程序去定位问题，日志信息已经提供了足够的诊断依据（一定什么事都要记录日志）
3. 由于panic会引起程序的崩溃，因此panic一般用于严重错误，如程序内部的逻辑不一致
4. 对于大部分漏洞，我们应该使用Go提供的错误机制，而不是panic，尽量避免程序的崩溃
5. 
```

## recover -> 从panic中恢复正常

```
1. 不加区分的恢复所有的panic异常，不是可取的做法
2. 你不应该试图去恢复其他包引起的panic
3. 你也不应该恢复一个由他人开发的函数引起的panic
4. 安全的做法是有选择性的recover
5. 换句话说，只恢复应该被恢复的panic异常
6. 恢复的异常占比应该尽可能低
7. >< 自己测试了，如果recover会导致一些问题
```

## 方法 （OOP）
``` 
1. 封装和组合
2. 
```


## 错误

``` 
1. 虽然Go有各种异常机制，但这些机制仅被使用在处理那些未被预料到的错误，
   即bug，而不是那些在健壮程序中应该被避免的程序错误
2. Go中大部分函数的代码结构几乎相同，首先是一系列的初始检查，
    防止错误发生，之后是函数的实际逻辑。   
3. 
```

## 类型转换

### 隐式转换

```
var f float64 = 3 + 0i
f = 2
f = 1e123
f = 'a'
上面的代码相当于
var f float64 = float64(3 + 0i)
f = float64(2)
f = float64(1e123)
f = float64('a')
```

## 格式化输出

```
1. %o %#o   分别以八进制以及带符号的八进制输出
2. %x %#x   十六进制以及带符号的十六进制
3. %c %q    分别输出单字符以及带引号的单字符
4. %b %#b   以二进制得形式输出
5. %T       打印类型
6. %v       暂时还不是很确定
7. %t       打印bool类型
```