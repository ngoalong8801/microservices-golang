package common

import (
	"gateway-service/model"
	"github.com/gin-gonic/gin"
	"net/http"
	grpc "proto/grpc/common"
)

func ReturnErrorResponse(ctx *gin.Context, code int, desc string) {
	ctx.JSON(code, model.ErrorResponse{
		ErrorCode:        http.StatusText(code),
		ErrorDescription: desc,
	})
}

func AsErrorResponse(error *grpc.Error, ctx *gin.Context) {
	if error == nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{
			ErrorCode:        grpc.ErrorCode_INTERNAL_SERVER_ERROR.String(),
			ErrorDescription: "Unknown error",
		})
		return
	}

	if len(error.Details) > 0 {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{
			ErrorCode:        error.Code.String(),
			ErrorDescription: error.Message,
			Errors:           error.Details,
			Exception:        error.Exception,
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{
			ErrorCode:        error.Code.String(),
			ErrorDescription: error.Message,
			Errors:           error.Details,
			Exception:        error.Exception,
		})
	}
}

func AsSuccessResponse(data interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, data)
}
