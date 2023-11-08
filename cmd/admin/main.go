package main

import (
	"LFS/apps/admin"
	_ "LFS/initialize"
	"LFS/internal/dal/invoker/lls_invoker/grpc_connect"
)

func main() {
	grpc_connect.InitResource("127.0.0.1:50051")
	admin.RouterInit()
}
