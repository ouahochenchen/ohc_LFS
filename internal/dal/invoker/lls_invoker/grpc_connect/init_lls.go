package grpc_connect

import (
	lls_pb "github.com/ouahochenchen/LLS/protocol/grpc/go"
)

var LlsClientVa lls_pb.LfsServiceClient

func InitLLS(s string) {
	err, client := NewLlsClient().getConnect(s)
	if err != nil {
		return
	}
	LlsClientVa = client

}
