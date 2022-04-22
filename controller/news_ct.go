package controller

import (
	"demo-gin-api/model/vo"
	"demo-gin-api/service"
	"demo-gin-api/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary getList
// @Tags information
// @Produce  json
// @param starttime query string false "開始時間 (yyyy-MM-dd)" default(2001-01-01)
// @param endtime query string false "結束時間 (yyyy-MM-dd)" default(2001-01-02)
// @param top query int false "取前幾筆" default(10)
// @Security BearerAuth
// @Success 200 {object} []vo.NewsVO
// @Router /news [get]
func NewsGetList(ctx *gin.Context) {

	starttime := ctx.Query("starttime")
	endtime := ctx.Query("endtime")
	top := ctx.Query("top")

	util.Success(ctx, "json", service.GetNewsList(starttime, endtime, top))
}

// @Summary getOne
// @Tags information
// @Produce  json
// @param id path int true "編號"
// @Security BearerAuth
// @Success 200 {object} vo.NewsVO
// @Router /news/{id} [get]
func NewsGet(ctx *gin.Context) {

	idStr := ctx.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	util.Success(ctx, "json", service.GetNews(id))
}

// @Summary addOne
// @Description 新增新聞
// @Tags information
// @Accept mpfd
// @Produce json
// @param postVO formData vo.NewsPostVO true "formData for NewsPostVO content"
// @Security BearerAuth
// @Success 200 {object} vo.NewsVO
// @Failure 400 {string} string
// @router /news [post]
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

// @Summary editOne
// @Description 編輯新聞
// @Tags information
// @Accept json
// @Produce plain
// @param id path int true "編號"
// @param putVO body vo.NewsPutVO true "body for NewsPutVO content"
// @Security BearerAuth
// @Success 200 {string} string
// @Failure 400 {string} string
// @router /news/{id} [put]
func NewsPut(ctx *gin.Context) {

	var nBody vo.NewsPutVO

	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err = ctx.ShouldBindJSON(&nBody); err != nil {
		util.Failure(ctx, "string", http.StatusBadRequest, err.Error())
		return
	}

	result, err := service.PutNews(id, nBody)

	if err != nil {
		util.Failure(ctx, "string", http.StatusBadRequest, err.Error())
		return
	}

	util.Success(ctx, "string", result)
}

// @Summary deleteOne
// @Description 刪除新聞
// @Tags information
// @Produce plain
// @param id path int true "編號"
// @Security BearerAuth
// @Success 200 {string} string
// @Failure 400 {string} string
// @router /news/{id} [delete]
func NewsDelete(ctx *gin.Context) {

	idStr := ctx.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	result, err := service.DeleteNews(id)

	if err != nil {
		util.Failure(ctx, "string", http.StatusBadRequest, err.Error())
		return
	}

	util.Success(ctx, "string", result)
}
