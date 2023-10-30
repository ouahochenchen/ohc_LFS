package main

import (
	"LFS/apps/api"
	_ "LFS/initialize"
)

func main() {
	api.RouterInit()
}
