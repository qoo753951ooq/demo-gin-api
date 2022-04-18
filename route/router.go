package route

import (
	"demo-gin-api/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() {

	gin.SetMode(gin.TestMode)
	r := gin.Default()

	apiGroup := r.Group("/demo-gin-api")
	setInformationRouter(apiGroup)

	r.Run(":8080")
}

func setInformationRouter(apiGroup *gin.RouterGroup) {

	newsRouter := apiGroup.Group("/news")
	{
		newsRouter.GET("", controller.NewsGetList)
		newsRouter.GET("/:id", controller.NewsGet)
		newsRouter.POST("", controller.NewsPost)
		newsRouter.PUT("/:id", controller.NewsPut)
	}
}
