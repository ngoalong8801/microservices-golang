package common

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

const (
	UnauthorizedForbidden = "You do not have permission to access this feature"
)

func GetMetadataFromContext(ctx *gin.Context) metadata.MD {
	md, exists := ctx.Get("md")
	var meta metadata.MD
	if exists {
		meta = md.(metadata.MD)
	}
	return meta
}
