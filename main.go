package main

import (
	db "demo-gin-api/database"
	"demo-gin-api/route"
)

// @title Swagger Demo API
// @version 1.0
// @description Demo Swagger 說明文件
// @termsOfService http://swagger.io/terms/
// @contact.name Alan Syu
// @contact.email qoo753951ooq@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:8080
// @BasePath /demo-gin-api
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	db.InitDB()
	route.InitRouter()
}
