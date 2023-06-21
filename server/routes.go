package server

import (
	_ "main/docs"
	"main/server/gateway"
	"main/server/handler"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes(server *Server) {

	//Player routes
	server.engine.POST("/buy-car", gateway.UserAuthorization, handler.BuyCarHandler)
	server.engine.PUT("/equip-car", gateway.UserAuthorization, handler.EquipCarHandler)
	server.engine.DELETE("/sell-car", gateway.UserAuthorization, handler.SellCarHandler)

	server.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
