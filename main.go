package main

import (
	db "demo-gin-api/database"
	"demo-gin-api/route"
)

func main() {
	db.InitDB()
	route.InitRouter()
}
