package controller

import (
	"demo-gin-api/service"
	"demo-gin-api/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewsGetList(ctx *gin.Context) {

	starttime := ctx.Query("starttime")
	endtime := ctx.Query("endtime")
	top := ctx.Query("top")

	util.Success(ctx, "json", service.GetNewsList(starttime, endtime, top))
}

func NewsGet(ctx *gin.Context) {

	idStr := ctx.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	util.Success(ctx, "json", service.GetNews(id))
}
