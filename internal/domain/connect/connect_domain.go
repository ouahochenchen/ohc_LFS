package connect

import (
	"LFS/internal/dal/repository/ls_connect_repo"
	"LFS/protocol/admin"
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
