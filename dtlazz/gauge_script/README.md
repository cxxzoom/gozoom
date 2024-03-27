## 你干嘛哈哈
对指定用户的指定量表进行一次随机答案的提交

## json文件
1. config.json 数据库配置文件
2. gauge.json 量表配置文件，指定哪些量表需要跑
3. query.json 指定提交的地址
4. user.json 指定哪些用户需要提交

## 编译
1. 看目标运行环境是windows还是linux
    1. windows：  SET GOOS=windows
    2. linux:     SET GOOS=linux
2. 目标运行环境的架构：32位还是64位，默认是64位
    1. 如果是32位： SET GOARCH=386
3. 设置完成之后，用 go env GOOS/GOARCH 查看是否配置成功
4. 这玩意儿配置要在 cmd配置，powershell不行；也不是不行，只是设置方式不是这样罢了
5. 编译
    1. windows： go build -o script_name.exe ./
    2. linux: go build -o script_name ./