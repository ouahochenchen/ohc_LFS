package connect

import (
	"LFS/internal/domain/connect"
	"LFS/protocol/admin"
	"github.com/gin-gonic/gin"
)

type ConnectUseCase interface {
	CreateConnect(ctx *gin.Context, info interface{}) (interface{}, error)
}
type connectUseCaseImpl struct {
	connectService connect.ConnectDomain
}

func NewConnectUseCase(connectService connect.ConnectDomain) ConnectUseCase {
	return &connectUseCaseImpl{
		connectService: connectService,
	}
}
func (c *connectUseCaseImpl) CreateConnect(ctx *gin.Context, info interface{}) (interface{}, error) {
	req := info.(*admin.CreateConnectRequest)
	return c.connectService.Create(req)
	//var req admin.CreateConnectRequest
	//err := ctx.BindJSON(&req)
	//create, err1 := c.connectService.Create(&req)
	//if err1 != nil || err != nil {
	//	ctx.JSON(http.StatusBadRequest, common.HttpCommonResponse{
	//		ReturnCode: -1,
	//		Message:    fmt.Sprintf("get param fail is %s", err.Error()),
	//	})
	//	return
	//}
	//ctx.JSON(http.StatusOK, common.HttpCommonResponse{
	//	ReturnCode: 0,
	//	Message:    "OK",
	//	Data: admin.CreateConnectResponse{
	//		Id:             &create,
	//		ResourceId:     &req.ResourceId,
	//		NextResourceId: &req.NextResourceId,
	//	},
	//})
	//return
}
