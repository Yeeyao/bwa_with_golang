# [build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang)

- 这里内容有点像 Unknown 的教程

- 当然，这里对相关的概念解释得更加清楚

- 纯 Go 代码可以使用 delve 来调试

- 压力测试 需要使用命令 go test -test.bench=".*" "" 内部表示需要运行的压力测试函数

- gotest package 来测试

- 后面的部分没有特别详细阅读

- 日志 seelog logrus

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

## Chapter 6 Session

- cookie 以及 session 服务器使用一种类似于 hash table 的结构来保存信息，每个网站都将被分配一个唯一的 sessionID

- 两者都用来客服 http 协议无状态的缺陷

- cookie 是本地计算机保存一些用户操作的历史信息，并在用户再次访问该站点时，

    - 浏览器通过 HTTP 协议将本地 cookie 内容发送给服务器
    
    - cookie 生命周期分为两种：会话 cookie 以及持久 cookie

- session 是在服务器保存用户操作的历史信息，服务器用 session id 来标识 session。

## 6.2 使用

- session 基本原理是由服务器为每个会话维护一份信息数据，客户端和服务端依靠一个全局唯一的标识来访问这份数据

- 发送 session 标识符: cookie 和 URL 重写 

- Cookie 服务端通过设置 set-cookie 头将 session 标识符传送到客户端，好组合每次请求都会带上

- URL 重写 返回给用户的页面里所有 URL 后面追加 session 标识符

# Chapter 7 文本处理

## XML 
- xml Unmarshal 函数来解析

- 同时，我们定义 struct tag 来辅助反射。首先读取 struct tag，如果没有，将读取对应的字段名

- 解析的时候，tag，字段名，XML 元素都是大小写敏感的。需要一一对应

```go
// 因为 struct 和 XML 有类似的树接口特征，所以我们定义类似的 struct 类型，然后将数据解析成 struct 对应的对象
func Unmarshal(data []byte, v interface{}) error
```

- XML 提供 Marshal 和 MarshalIndent 两个函数来生成 XML 文件

## Json

- json 格式，使用的函数名字与 xml 的类似

- 同时，寻找策略首先找可导出的 struct 字段，之后找符合字段名的导出字段，最后找大小写不敏感的导出字段

- 这里寻找字段的策略使得 JSON 解析的时候只会解析能找得到的字段，找不到的字段会被忽略

- JSON 包存储 JSON 对象的格式 

- 下面的格式配合断言就可以访问未知格式的 JSON 文件内容

- [一个推荐的库](https://github.com/bitly/go-simplejson)

```go
map[string]interface{}
```

- 需要注意，之后 struct 中的导出类型才会被输出，需要小写的字段名则用 struct tag 来实现

- "-" "omitempty" ",string" 等 struct tag 的含义

## 正则表达式

- regexp 包。Go 实现 RE2 标准

- 可以用 strings 包则尽量使用它

- Match MatchReader MatchString 仅仅是后面的参数不同 

- Compile 会解析正则表达式是否合法，如果合法，则会返回一个 Regexp，之后可以用它在任意字符串上执行操作

- 搜索函数以及 Expand 实现命令行的例子

## 模板处理

- MVC 设计模式 Model 处理数据，View 展现结果，Controller 控制用户的请求

- Go 语言中 {{}} 里面的都是 pipelines

- 模板变量 可以在 with range if 中声明局部变量，作用域在 {{end}} 前

- 模板函数 

```go
// 这里通过一个唯一的名字与对应的 Go 函数关联
// 需要在 parse 前就将函数注入
type FuncMap map[string]interface{}
```

- Must 函数检测模板是否正确

## 小结

- 文本处理工具的介绍，XML、JSON、正则表达式以及模板。前两个是数据交互工具，正则表达式可以处理文本，模板可以展现数据

# Chapter8 Socket 编程

## Socket 

### TCP Socket

- net package 有一个类型 TCPConn 作为客户端和服务器交互通道

- net package 中 DialTCP 函数建立一个 TCP 连接，返回一个 TCPConn 类型的对象，客户端和服务器都拥有 TCPConn 对象来进行数据交换

```go
// 建立连接
func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err os.Error)

// 设置超时
func DialTimeout(net, addr string, timeout time.Duration) (Conn, error)

// 定时发送包，类似心跳包功能
func (c *TCPConn) SetKeepAlive(keepalive bool) os.Error
```

### UDP Socket

## WebSocket

- 采用了特殊的包头，使得浏览器和服务器只需要做一个握手动作，就可以在两者之间建立一条连接通道

## REST

- Resources Representation State Transfer

### RESTful 架构 

- 每个 URL 代表一种资源

- 客户端和服务器间，传递这种资源的某种表现层

- 客户端通过四个 HTTP 动词，对服务器字段进行操作，实现表现层状态转化

- Web 应用要满足 REST 最重要的原则是:客户端和服务器之间的交互在请求之间是无状态的,即从客户端到服务器的每个请求都必须包含理解请求所必需的信息。

- 另一个 REST 原则是系统分层。

## RPC

- GO 支持三种级别的 PRC：TCP HTTP JSONPRC

