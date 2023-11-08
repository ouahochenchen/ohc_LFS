package grpc_connect

import (
	lls_pb "github.com/ouahochenchen/LLS/protocol/grpc/go"
	"google.golang.org/grpc"
	"log"
)

type ResourceClient interface {
	getConnect(s string) (error, lls_pb.SiteServiceClient)
}

type resourceClientConnectImpl struct {
}

func NewResourceClient() ResourceClient {
	return &resourceClientConnectImpl{}
}

func (lls *resourceClientConnectImpl) getConnect(s string) (error, lls_pb.SiteServiceClient) {
	conn, err := grpc.Dial(s, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to grpc_connect: %v", err)
		return err, nil
	}
	client := lls_pb.NewSiteServiceClient(conn)
	return nil, client

}
