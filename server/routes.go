package server

import (
	_ "main/docs"
	"main/server/gateway"
	"main/server/handler"

	admin "main/server/handler/admin"
	player "main/server/handler/player"

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

	//Admin Login Route

	server.engine.POST("/admin/signup", admin.AdminSignUpHandler)
	server.engine.POST("/forgot-password", admin.ForgotPasswordHandler)
	server.engine.PATCH("/reset-password", admin.ResetPasswordHandler)

	//Auth routes
	server.engine.POST("/guest-login", admin.GuestLoginHandler)
	server.engine.POST("/login", admin.LoginHandler)
	server.engine.PUT("/update-email", admin.UpdateEmailHandler)

	//player details
	server.engine.GET("/player-details", handler.GetPlayerDetailsHandler)

	//Player routes
	server.engine.POST("/car/buy", player.BuyCarHandler)
	server.engine.PUT("/car/equip", player.EquipCarHandler)
	server.engine.DELETE("/car/sell", player.SellCarHandler)
	server.engine.POST("/car/repair", player.RepairCarHandler)
	server.engine.GET("/car/get-all", player.GetAllCarsHandler)
	server.engine.POST("/car/get-by-id", player.GetCarByIdHandler)

	//Player garage routes
	server.engine.POST("/garage/buy", player.BuyGarageHandler)
	server.engine.GET("/garages/get-all", admin.GetAllGarageListHandler)
	server.engine.POST("/garage/add-car", player.AddCarToGarageHandler)
	server.engine.PUT("/garage/upgrade", player.UpgradeGarageHandler)
	server.engine.GET("/garage/get", player.GetPlayerGarageListHandler)

	//Admin garage routes
	server.engine.POST("/admin/garage/add", gateway.AdminAuthorization, admin.AddGarageHandler)
	server.engine.DELETE("/admin/garage/delete", gateway.AdminAuthorization, admin.DeleteGarageHandler)
	server.engine.PUT("/admin/garage/update", gateway.AdminAuthorization, admin.UpdateGarageHandler)

	//Car upgrade routes
	server.engine.PUT("/car/upgrade/engine", player.UpgradeEngineHandler)
	server.engine.PUT("/car/upgrade/turbo", player.UpgradeTurboHandler)
	server.engine.PUT("/car/upgrade/intake", player.UpgradeIntakeHandler)
	server.engine.PUT("/car/upgrade/nitrous", player.UpgradeNitrousHandler)
	server.engine.PUT("/car/upgrade/body", player.UpgradeBodyHandler)
	server.engine.PUT("/car/upgrade/tires", player.UpgradeTiresHandler)
	server.engine.PUT("/car/upgrade/transmission", player.UpgradeTransmissionHandler)

	//car Customiztion routes
	server.engine.PUT("/car/customise/color", player.ColorCustomizeHandler)
	server.engine.PUT("/car/customise/wheels", player.WheelsCustomizeHandler)
	server.engine.PUT("/car/customise/interior", player.InteriorCustomizeHandler)
	server.engine.PUT("/car/customise/license", player.LicenseCustomizeHandler)

	server.engine.GET("/get-customization", player.GetCarCustomiseHandler)
	server.engine.GET("/get-color-category", player.GetCarColorCategoriesHandler)
	server.engine.GET("/get-color-type", player.GetCarColorTypesHandler)
	server.engine.GET("/get-colors", player.GetCarColorsHandler)

	//add dummy data in db
	server.engine.GET("/add-dummy-data", handler.AddDummyDataHandler)

	//swagger routes
	server.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
