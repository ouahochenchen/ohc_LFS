package grpc_req

import (
	"LFS/internal/dal/invoker/lls/grpc_connect"
	_go "LFS/protocol/grpc/go"
	"context"
)

type ReqLLS interface {
	GrpcReq(context context.Context, req *_go.LfsRequest) (*_go.LfsResponse, error)
}

type reqLLSImpl struct {
}

func NewReqLLS() ReqLLS {
	return &reqLLSImpl{}
}

func (r *reqLLSImpl) GrpcReq(ctx context.Context, req *_go.LfsRequest) (*_go.LfsResponse, error) {
	reps, err := grpc_connect.LlsClientVa.CreateLineOrder(ctx, req)
	if err != nil {
		return nil, err
	}
	return reps, nil
}
