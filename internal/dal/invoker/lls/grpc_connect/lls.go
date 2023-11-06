package grpc_connect

import (
	_go "LFS/protocol/grpc/go"
	"google.golang.org/grpc"
	"log"
)

type LlsClient interface {
	getConnect(s string) (error, _go.LfsServiceClient)
}

type llsClientConnectImpl struct {
}

func NewLlsClient() LlsClient {
	return &llsClientConnectImpl{}
}

func (lls *llsClientConnectImpl) getConnect(s string) (error, _go.LfsServiceClient) {
	conn, err := grpc.Dial(s, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to grpc_connect: %v", err)
		return err, nil
	}

	client := _go.NewLfsServiceClient(conn)
	return nil, client

}
