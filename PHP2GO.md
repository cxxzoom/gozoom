# PHP2GO
## 基础
### 变量



## interface
1. 相同之处：
    ```
    关键字都是interface；
    都是约定一些待实现的方法
    都需要全部实现接口里面的方才算实现了这个接口
    ```
2. 不同之处：
    ```
   php需要implements，属于是显示实现；
   go是隐式实现，只要你实现了这个interface里面的方法，就行
    ```
## json操作
php
```php
<?php
$file = "\path\to\file.ext";
$json = json_decode(file_get_contents($file));
```
go

````go
package main

import (
   "encoding/json"
   "fmt"
   "os"
)

type Jctx struct {
   Name string `json:"name"` // 前面的一定要大写，才是可导出的，后面的json:"name"对应json文件里面的字段
   Age  int    `json:"age"`
}

func main() {
   filePath := "/path/to/file.ext"
   ctx, err := os.ReadFile(filePath)
   if err != nil {
      fmt.Println(err)
   }

   data := Jctx{}

   err = json.Unmarshal(ctx, &data)
   if err != nil {
	   fmt.Println(err)
   }
}
````