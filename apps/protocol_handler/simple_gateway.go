package protocol_handler

import (
	"LFS/protocol/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SimpleHandlerFunc func(ctx *gin.Context, req interface{}) (interface{}, error)

func SimpleGateway(handler SimpleHandlerFunc, target interface{}) gin.HandlerFunc {
	defer func() {
		if err := recover(); err != nil {
			_ = fmt.Sprintf("捕获到了panic:%s", err)
		}
	}()
	return func(context *gin.Context) {
		if err := context.BindJSON(target); err != nil {
			context.JSON(http.StatusBadRequest, common.HttpCommonResponse{
				ReturnCode: -1,
				Message:    fmt.Sprintf("get param fail: %s", err.Error()),
			})
			return
		}

		resp, err := handler(context, target)
		if err != nil {
			context.JSON(http.StatusBadRequest, common.HttpCommonResponse{
				ReturnCode: -1,
				Message:    fmt.Sprintf("get param fail: %s", err.Error()),
			})
			return
		}
		context.JSON(http.StatusOK, common.HttpCommonResponse{
			ReturnCode: 0,
			Message:    "OK",
			Data:       resp,
		})
	}
}
