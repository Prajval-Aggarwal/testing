package server

import (
	_ "main/docs"
	"main/server/gateway"
	"main/server/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes(server *Server) {

	//CORS allow
	server.engine.Use(gateway.CORSMiddleware())

	server.engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Sever listening",
		})
	})

	//Auth routes
	server.engine.POST("/guest-login", handler.GuestLoginHandler)
	server.engine.POST("/login", handler.LoginHandler)
	server.engine.PUT("/update-email", gateway.UserAuthorization, handler.UpdateEmailHandler)

	//Player routes
	server.engine.POST("/car/buy", gateway.UserAuthorization, handler.BuyCarHandler)
	server.engine.PUT("/car/equip", gateway.UserAuthorization, handler.EquipCarHandler)
	server.engine.DELETE("/car/sell", gateway.UserAuthorization, handler.SellCarHandler)
	server.engine.POST("/car/repair", gateway.UserAuthorization, handler.RepairCarHandler)
	server.engine.GET("/car/get-all", handler.GetAllCarsHandler)
	server.engine.POST("/car/get-by-id", handler.GetCarByIdHandler)
	//Player garage routes
	server.engine.POST("/garage/buy", gateway.UserAuthorization, handler.BuyGarageHandler)
	server.engine.GET("/garages/get-all", handler.GetAllGarageListHandler)
	server.engine.POST("/garage/add-car", gateway.UserAuthorization, handler.AddCarToGarageHandler)
	server.engine.PUT("/garage/upgrade", gateway.UserAuthorization, handler.UpgradeGarageHandler)
	server.engine.GET("/garage/get", gateway.UserAuthorization, handler.GetPlayerGarageListHandler)

	//Car upgrade routes
	server.engine.PUT("/car/upgrade/engine", gateway.UserAuthorization, handler.UpgradeEngineHandler)
	server.engine.PUT("/car/upgrade/turbo", gateway.UserAuthorization, handler.UpgradeTurboHandler)
	server.engine.PUT("/car/upgrade/intake", gateway.UserAuthorization, handler.UpgradeIntakeHandler)
	server.engine.PUT("/car/upgrade/nitrous", gateway.UserAuthorization, handler.UpgradeNitrousHandler)
	server.engine.PUT("/car/upgrade/body", gateway.UserAuthorization, handler.UpgradeBodyHandler)
	server.engine.PUT("/car/upgrade/tires", gateway.UserAuthorization, handler.UpgradeTiresHandler)
	server.engine.PUT("/car/upgrade/transmission", gateway.UserAuthorization, handler.UpgradeTransmissionHandler)

	//car Customiztion routes
	server.engine.PUT("/car/customise/color", gateway.UserAuthorization, handler.ColorCustomizeHandler)
	server.engine.PUT("/car/customise/wheels", gateway.UserAuthorization, handler.WheelsCustomizeHandler)
	server.engine.PUT("/car/customise/interior", gateway.UserAuthorization, handler.InteriorCustomizeHandler)
	server.engine.PUT("/car/customise/license", gateway.UserAuthorization, handler.LicenseCustomizeHandler)

	server.engine.GET("/get-customization", handler.GetCarCustomiseHandler)
	server.engine.GET("/get-color-category", handler.GetCarColorCategoriesHandler)
	server.engine.GET("/get-color-type", handler.GetCarColorTypesHandler)
	server.engine.GET("/get-colors", handler.GetCarColorsHandler)

	//add dummy data in db
	server.engine.GET("/add-dummy-data", handler.AddDummyDataHandler)

	//swagger routes
	server.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
