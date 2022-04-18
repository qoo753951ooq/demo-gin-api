package controller

import (
	"demo-gin-api/model/vo"
	"demo-gin-api/service"
	"demo-gin-api/util"
	"net/http"
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

func NewsPost(ctx *gin.Context) {

	var nBody vo.NewsPostVO

	if err := ctx.ShouldBind(&nBody); err != nil {
		util.Failure(ctx, "string", http.StatusBadRequest, err.Error())
		return
	}

	news, err := service.PostNews(nBody)

	if err != nil {
		util.Failure(ctx, "string", http.StatusBadRequest, err.Error())
		return
	}

	util.Success(ctx, "json", news)
}
