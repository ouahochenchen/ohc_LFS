package connect

import (
	"LFS/internal/dal/repositry/ls_connect_repo"
	"LFS/protocol/admin"
)

var repo = ls_connect_repo.NewConnectRepo()

type ConnectDomain interface {
	Create(*admin.CreateConnectRequest) (*admin.CreateConnectResponse, error)
}
type connectDomainImpl struct {
}

func NewConnectDomain() ConnectDomain {
	return &connectDomainImpl{}
}
func (c *connectDomainImpl) Create(req *admin.CreateConnectRequest) (*admin.CreateConnectResponse, error) {
	tab := ls_connect_repo.LaneSiteConnectConfigurationTab{
		ResourceId:     req.ResourceId,
		NextResourceId: req.NextResourceId,
		ResourceType:   req.ResourceType,
		NextType:       req.NextType,
	}
	createId, err := repo.Create(&tab)
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
