package route

import (
	"demo-gin-api/controller"
	"demo-gin-api/docs"
	"demo-gin-api/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	apiGroup := r.Group(docs.SwaggerInfo.BasePath)
	setInformationRouter(apiGroup)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}

func setInformationRouter(apiGroup *gin.RouterGroup) {

	authRouter := apiGroup.Group("/user")
	{
		authRouter.POST("/login", controller.UserLoginPost)
	}

	newsRouter := apiGroup.Group("/news").Use(middleware.JWTAuth())
	{
		newsRouter.GET("", controller.NewsGetList)
		newsRouter.GET("/:id", controller.NewsGet)
		newsRouter.POST("", controller.NewsPost)
		newsRouter.PUT("/:id", controller.NewsPut)
		newsRouter.DELETE("/:id", controller.NewsDelete)
	}
}
