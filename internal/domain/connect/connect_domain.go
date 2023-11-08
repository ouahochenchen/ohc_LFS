package connect

import (
	"LFS/internal/dal/invoker/lls_invoker/grpc_connect"
	"LFS/internal/dal/repository/ls_connect_repo"
	"LFS/internal/infrastructure/err_code"
	"LFS/protocol/admin"
	"context"
	_go "github.com/ouahochenchen/LLS/protocol/grpc/go"
)

type ConnectDomain interface {
	Create(*admin.CreateConnectRequest) (*admin.CreateConnectResponse, error)
}
type connectDomainImpl struct {
	connectService ls_connect_repo.ConnectRepo
}

func NewConnectDomain(connectRepo ls_connect_repo.ConnectRepo) ConnectDomain {
	return &connectDomainImpl{
		connectService: connectRepo,
	}
}
func (c *connectDomainImpl) Create(req *admin.CreateConnectRequest) (*admin.CreateConnectResponse, error) {
	tab := ls_connect_repo.LaneSiteConnectConfigurationTab{
		ResourceId:     req.ResourceId,
		NextResourceId: req.NextResourceId,
		ResourceType:   req.ResourceType,
		NextType:       req.NextType,
	}
	reqGrpc := &_go.ExistSiteLineRequest{
		ResourceId:   req.ResourceId,
		NextId:       req.NextResourceId,
		ResourceType: req.ResourceType,
		NextType:     req.NextType,
	}
	resource, err := grpc_connect.ResourceClientVa.IsExistResource(context.Background(), reqGrpc)
	if resource.IsExist == false {
		return nil, &err_code.MyError{Msg: "所选内容含不存在的点线"}
	}
	if err != nil {
		return nil, err
	}
	createId, err := c.connectService.Create(&tab)
	resp := admin.CreateConnectResponse{
		Id:             &createId,
		ResourceId:     &tab.ResourceId,
		NextResourceId: &tab.NextResourceId,
	}
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
