# [地址](https://astaxie.gitbooks.io/build-web-application-with-golang/content/zh/03.4.html)

## Chapter 3

- Go http 有两个核心的功能：Conn ServeMux

- 客户端每次请求都会创建一个 Conn，这个 Conn 保存了该次请求的信息，然后再传递给相应的 handler，后者便可以读取到相应的 header 信息

