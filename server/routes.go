package server

import (
	_ "main/docs"
	"main/server/gateway"
	"main/server/handler"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes(server *Server) {

	//CORS allow
	server.engine.Use(gateway.CORSMiddleware())

	//Player routes
	server.engine.POST("/buy-car", gateway.UserAuthorization, handler.BuyCarHandler)
	server.engine.PUT("/equip-car", gateway.UserAuthorization, handler.EquipCarHandler)
	server.engine.DELETE("/sell-car", gateway.UserAuthorization, handler.SellCarHandler)

	//Player garage routes
	server.engine.GET("/get-all-garages", gateway.UserAuthorization, handler.GetAllGarageListHandler)
	server.engine.POST("/buy-garage", gateway.UserAuthorization, handler.BuyCarHandler)
	server.engine.POST("/add-car-garage", gateway.UserAuthorization, handler.AddCarToGarageHandler)
	server.engine.PUT("/upgrage-garage", gateway.UserAuthorization, handler.UpgradeGarageHandler)
	server.engine.GET("/get-garage", gateway.UserAuthorization, handler.GetPlayerGarageListHandler)

	//Auth routes
	server.engine.POST("/guest-login", handler.GuestLoginHandler)
	server.engine.POST("/login", handler.LoginHandler)
	server.engine.PUT("/update-email", handler.UpdateEmailHandler)

	//swagger routes
	server.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
