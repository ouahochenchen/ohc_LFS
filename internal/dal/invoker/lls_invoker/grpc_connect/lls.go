package grpc_connect

import (
	lls_pb "github.com/ouahochenchen/LLS/protocol/grpc/go"
	"google.golang.org/grpc"
	"log"
)

type LlsClient interface {
	getConnect(s string) (error, lls_pb.LfsServiceClient)
}

type llsClientConnectImpl struct {
}

func NewLlsClient() LlsClient {
	return &llsClientConnectImpl{}
}

func (lls *llsClientConnectImpl) getConnect(s string) (error, lls_pb.LfsServiceClient) {
	conn, err := grpc.Dial(s, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to grpc_connect: %v", err)
		return err, nil
	}
	client := lls_pb.NewLfsServiceClient(conn)
	return nil, client

}
