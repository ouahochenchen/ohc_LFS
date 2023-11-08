package grpc_connect

import (
	resoucepb "github.com/ouahochenchen/LLS/protocol/grpc/go"
)

var ResourceClientVa resoucepb.SiteServiceClient

//var ReqResource grpc_req.ReqLLS

func InitResource(s string) {
	err, client := NewResourceClient().getConnect(s)
	if err != nil {
		return
	}
	ResourceClientVa = client

}
