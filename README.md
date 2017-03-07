# 关于grpc

http://www.grpc.io/docs/quickstart/go.html



# 生成代码

cd src/helloworld
protoc -I ./ helloworld.proto --go_out=plugins=grpc:.
将proto文件生成到当前目录下面。


# run server & run client

go run server/main.go
go run client/main.go

# 其他写在博客上面：

http://blog.csdn.net/freewebsys/article/details/59483427

# 增加了consul的测试：

https://github.com/freewebsys/grpc-go-demo/blob/master/src/consul/consul_test.go
测试的时候注册了3个tomcat，第3个在检测的时候是个错误的地址。
consul修改端口启动：

```
consul agent -dev -ui -server -node=consul-dev -client=10.0.2.15 -dns-port=53 -domain=freeweb.consul 
```
-dns-port=53 是将consul伪装成一个dns服务器。
-domain=freeweb.consul 是设置自己域。