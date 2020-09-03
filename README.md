# MyGin

## Gin介绍

- Gin是一个Golang的框架

## Gin安装

```
官方英文文档 https://gin-gonic.com/docs/quickstart/
官方中文文档 https://gin-gonic.com/zh-cn/docs/introduction/
```

```
源码 https://github.com/gin-gonic/gin
```

```
下载安装
$ go get -u -v github.com/gin-gonic/gin
```

下载之前需要设置代理

```
开启Go Module 
$ go env -w GO111MODULE=on
```

```
开启代理
$ go env -w GOPROXY=https://goproxy.cn,direct
```

```
查看环境变量修改
$ go env
```

```
初始化模块
$ go mod init xuqiulin.com/mygin
```

## 错误

\# runtime/cgo exec: "gcc": executable file not found in %PATH%

安装mingw64不仅能编译64位程序，也能编译32位程序

找到MinGW-W64-install.exe

- https://sourceforge.net/projects/mingw-w64/files/mingw-w64/mingw-w64-release/

- https://osdn.net/projects/mingw/downloads/68260/mingw-get-setup.exe/

安装教程

- https://jingyan.baidu.com/article/0320e2c11564ca1b87507b8f.html





















