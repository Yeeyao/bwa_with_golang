# [build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang)

- 这里内容有点像 Unknown 的教程

## Chapter 3

### Go http 有两个核心的功能：Conn ServeMux

### 客户端每次请求都会创建一个 Conn，这个 Conn 保存了该次请求的信息，然后再传递给相应的 handler，后者便可以读取到相应的 header 信息

## Chapter4 表单输入

- 从 r *http.Request 中获取信息

- 请求方法通过 r.Method 来获取

- 需要使用 r.ParserForm() 来解析 Form 才能对表单数据进行操作

- html/template 有函数可以帮助转义，防止 XSS 攻击

- 防止多次递交 在表单中添加一个带有唯一值的隐藏字段

## Chapter5 Database

- Go 为开发数据库驱动定义了一些标准接口，开发者可以根据定义的接口来开发对应的数据库驱动

- sql.Register 用来注册数据库驱动。注册需要实现 init 函数。

- 其中内部通过一个 map 来存储用户定义的相应驱动。因此可以同时注册多个数据库驱动

- Driver 定义了 Open(name string) method，返回一个 Conn 接口。后者只能用来进行一次 goroutine 操作

- Stmt 与 Conn 关联，只能用于一个 goroutine，是一种准备好的状态

- driver.Tx 为事务处理，一般是提交或者回滚两种

- driver.Execer 为一个 Conn 可选的实现接口

- driver.Result 是操作返回的结果接口定义

- driver.Value 可以容纳任何数据的空接口，数据是驱动能够操作的 Value。同时 ValueConverter 接口定义了把普通值转化为 Value 的接口

- PostgreSQL 对象-关系数据库服务器，BSD-许可证。

