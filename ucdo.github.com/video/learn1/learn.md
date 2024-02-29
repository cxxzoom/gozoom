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

#### 第一天完结
才看到第5个视频