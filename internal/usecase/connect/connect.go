package connect

import (
	"LFS/internal/domain/connect"
	"LFS/protocol/admin"
	"LFS/protocol/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ConnectUseCase interface {
	CreateConnect(ctx *gin.Context)
}
type connectUseCaseImpl struct {
}

func NewConnectUseCase() ConnectUseCase {
	return &connectUseCaseImpl{}
}
func (*connectUseCaseImpl) CreateConnect(ctx *gin.Context) {
	var req admin.CreateConnectRequest
	err := ctx.BindJSON(&req)
	create, err1 := connect.NewConnectDomain().Create(&req)
	if err1 != nil && err != nil {
		ctx.JSON(http.StatusBadRequest, common.HttpCommonResponse{
			ReturnCode: -1,
			Message:    fmt.Sprintf("get param fail is %s", err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusBadRequest, common.HttpCommonResponse{
		ReturnCode: 0,
		Message:    "OK",
		Data: admin.CreateConnectResponse{
			Id:             &create,
			ResourceId:     &req.ResourceId,
			NextResourceId: &req.NextResourceId,
		},
	})
	return
}
