# [build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang)

- 这里内容有点像 Unknown 的教程

## Chapter 3

- Go http 有两个核心的功能：Conn ServeMux

- 客户端每次请求都会创建一个 Conn，这个 Conn 保存了该次请求的信息，然后再传递给相应的 handler，后者便可以读取到相应的 header 信息

## Chapter4 表单输入

- 从 r *http.Request 中获取信息

- 请求方法通过 r.Method 来获取

- 需要使用 r.ParserForm() 来解析 Form 才能对表单数据进行操作

- html/template 有函数可以帮助转义，防止 XSS 攻击

- 防止多次递交 在表单中添加一个带有唯一值的隐藏字段
