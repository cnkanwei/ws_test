# ws_test
A cli program to help test websocket server
一个简单的 测试websocket 的 命令行 客户端程序
## 启动参数
``` shell
  -origin string
        域名 (default "http://127.0.0.1:8080")
  -protocol string
        协议 (default "ws")
  -uri string
        uri (default "/echo")
```
## 功能
连接成功之后
* 会2秒发送一个 心跳包：{"type":"ping"} 
* 除了心跳包之外的消息会显示出来
* 可以在命令行输入数据，按回车发出消息。
