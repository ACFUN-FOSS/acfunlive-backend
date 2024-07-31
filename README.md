# acfunlive-backend

AcFun 直播通用后端

### 安装

```
go get -u github.com/ACFUN-FOSS/acfunlive-backend
```

### 运行参数

`-port 端口`：设置后端 WebSocket 端口，默认是 15368

`-tcp`：使用 TCP 连接弹幕服务器

`-debug`：打印调试信息，默认不打印后端发送和接收的消息

`-logall`：打印所有调试信息，会打印后端发送和接收的消息

`-logfile 日志文件`：将日志输出到指定文件

`-logmax 日志文件大小`：单个日志文件最大的大小（按字节算），默认是 20MB，超出部分会保存为备份日志

`-logversions 备份日志数量`：备份日志的数量，默认是 2，备份日志是在日志文件后面加上“.1”、“.2”等后缀名

### 文档

[后端 WebSocket 接口](https://github.com/ACFUN-FOSS/acfunlive-backend/blob/main/doc/%E5%90%8E%E7%AB%AFWebSocket%E6%8E%A5%E5%8F%A3.md)
