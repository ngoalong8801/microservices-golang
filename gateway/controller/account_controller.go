package controller

import (
	"gateway-service/common"
	"gateway-service/grpc/client"
	"gateway-service/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	gcommon "proto/grpc/common"
)

type AccountController struct {
	accountClient *client.AccountClient
}

func NewAccountController(accountClient *client.AccountClient) *AccountController {
	return &AccountController{
		accountClient: accountClient,
	}
}

func (controller *AccountController) GetUser(ctx *gin.Context) {

	req := &gcommon.FindByIdRequest{
		Id: ctx.Param("id"),
	}

	res, err := controller.accountClient.GetUser(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		log.Println("Failed when get user account", err.Error())
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.GetSuccess() {
		ctx.JSON(http.StatusOK, res.GetData())
	} else {
		ctx.JSON(http.StatusBadRequest, model.AsErrorResponse(res.GetError()))
	}
}
