package grpc_connect

import (
	"LFS/internal/dal/invoker/lls/grpc_req"
	_go "LFS/protocol/grpc/go"
)

var LlsClientVa _go.LfsServiceClient
var ReqLLS grpc_req.ReqLLS

func Init(s string) {
	err, client := NewLlsClient().getConnect(s)
	if err != nil {
		return
	}
	LlsClientVa = client
	r := grpc_req.NewReqLLS()
	ReqLLS = r
}
