package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewsGetList(ctx *gin.Context) {
	ctx.String(http.StatusOK, "%s", "Test")
}
