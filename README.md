# 关于grpc

http://www.grpc.io/docs/quickstart/go.html



# 生成代码

cd src
protoc -I ./ helloworld.proto --go_out=plugins=grpc:helloworld
将proto文件生成到当前目录下面。


# run server

go run server/main.go