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

	//Player garage routes
	server.engine.GET("/get-all-garages", gateway.UserAuthorization, handler.GetAllGarageListHandler)
	server.engine.GET("/buy-garage", gateway.UserAuthorization, handler.BuyCarHandler)
	server.engine.GET("/add-car-garage", gateway.UserAuthorization, handler.AddCarToGarageHandler)
	server.engine.GET("/upgrage-garage", gateway.UserAuthorization, handler.UpgradeGarageHandler)
	server.engine.GET("/get-garage", gateway.UserAuthorization, handler.GetPlayerGarageListHandler)

	server.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
