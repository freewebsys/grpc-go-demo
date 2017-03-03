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
