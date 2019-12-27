# Common

>   个人 `GO` 项目中，使用到的通用基础类库。


### 使用示例


#### 助手工具类

```go
package main

import (
	"github.com/douyasi/common/helper"
)

func main() {
    // 字符串 md5 值
    helper.StrMD5("hello")
    // 路径是否存在
    helper.PathExists("/temp/temp.log")
    // 生成固定长度的随机字符串
    helper.RandomStr(6)
}
```

#### 日志

位于 `log/logger.go` 文件，基于 `logrus` ，支持同时向控制台与文件输出日志。

```go
package main
import "github.com/douyasi/common/log"

func init() {
    log.InitLog("./go.log")
}


func main() {
    info := "hello world"
    log.Logger().Info("info:", info)
    err := "something wrong, please check your params"
    log.Logger().Error("error:", err)
}
```

#### 字符串变换

位于 `str/case.go` 文件，具体查看 [README](str/README.md) 。

#### 统一结构的 `JSON` 响应

位于 `resp/json_result.go` 文件，结合一些框架，可以实现统一结构的接口响应。


```json
{
  "code": 20000,
  "message": "success",
  "data": {
  },
  "timestamp": "1577350555"
}
```

