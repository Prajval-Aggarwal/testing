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

<<<<<<< HEAD
	server.engine.POST("/admin/login", admin.AdminLoginHandler)
=======
>>>>>>> 43ea51cb16ea49a0369579929718436edab5da04
	server.engine.POST("/admin/signup", admin.AdminSignUpHandler)
	server.engine.POST("/forgot-password", admin.ForgotPasswordHandler)
	server.engine.PATCH("/reset-password", admin.ResetPasswordHandler)

	//Auth routes
	server.engine.POST("/guest-login", admin.GuestLoginHandler)
	server.engine.POST("/login", admin.LoginHandler)
<<<<<<< HEAD
	server.engine.PUT("/update-email", gateway.UserAuthorization, admin.UpdateEmailHandler)
=======
	server.engine.PUT("/update-email", admin.UpdateEmailHandler)
>>>>>>> 43ea51cb16ea49a0369579929718436edab5da04

	//player details
	server.engine.GET("/player-details", handler.GetPlayerDetailsHandler)

	//Player routes
<<<<<<< HEAD
	server.engine.POST("/car/buy", gateway.UserAuthorization, player.BuyCarHandler)
	server.engine.PUT("/car/equip", gateway.UserAuthorization, player.EquipCarHandler)
	server.engine.DELETE("/car/sell", gateway.UserAuthorization, player.SellCarHandler)
	server.engine.POST("/car/repair", gateway.UserAuthorization, player.RepairCarHandler)
=======
	server.engine.POST("/car/buy", gateway.AdminAuthorization, player.BuyCarHandler)
	server.engine.PUT("/car/equip", gateway.AdminAuthorization, player.EquipCarHandler)
	server.engine.DELETE("/car/sell", gateway.AdminAuthorization, player.SellCarHandler)
	server.engine.POST("/car/repair", gateway.AdminAuthorization, player.RepairCarHandler)
>>>>>>> 43ea51cb16ea49a0369579929718436edab5da04
	server.engine.GET("/car/get-all", player.GetAllCarsHandler)
	server.engine.POST("/car/get-by-id", player.GetCarByIdHandler)

	//Player garage routes
<<<<<<< HEAD
	server.engine.POST("/garage/buy", gateway.UserAuthorization, player.BuyGarageHandler)
	server.engine.GET("/garages/get-all", admin.GetAllGarageListHandler)
	server.engine.POST("/garage/add-car", gateway.UserAuthorization, player.AddCarToGarageHandler)
	server.engine.PUT("/garage/upgrade", gateway.UserAuthorization, player.UpgradeGarageHandler)
	server.engine.GET("/garage/get", gateway.UserAuthorization, player.GetPlayerGarageListHandler)

	//Admin garage routes
	server.engine.POST("/admin/garage/add", gateway.AdminAuthorization, admin.AddGarageHandler)
	server.engine.DELETE("/admin/garage/delete", gateway.AdminAuthorization, admin.DeleteGarageHandler)
	server.engine.PUT("/admin/garage/update", gateway.AdminAuthorization, admin.UpdateGarageHandler)

	//Car upgrade routes
	server.engine.PUT("/car/upgrade/engine", gateway.UserAuthorization, player.UpgradeEngineHandler)
	server.engine.PUT("/car/upgrade/turbo", gateway.UserAuthorization, player.UpgradeTurboHandler)
	server.engine.PUT("/car/upgrade/intake", gateway.UserAuthorization, player.UpgradeIntakeHandler)
	server.engine.PUT("/car/upgrade/nitrous", gateway.UserAuthorization, player.UpgradeNitrousHandler)
	server.engine.PUT("/car/upgrade/body", gateway.UserAuthorization, player.UpgradeBodyHandler)
	server.engine.PUT("/car/upgrade/tires", gateway.UserAuthorization, player.UpgradeTiresHandler)
	server.engine.PUT("/car/upgrade/transmission", gateway.UserAuthorization, player.UpgradeTransmissionHandler)

	//car Customiztion routes
	server.engine.PUT("/car/customise/color", gateway.UserAuthorization, player.ColorCustomizeHandler)
	server.engine.PUT("/car/customise/wheels", gateway.UserAuthorization, player.WheelsCustomizeHandler)
	server.engine.PUT("/car/customise/interior", gateway.UserAuthorization, player.InteriorCustomizeHandler)
	server.engine.PUT("/car/customise/license", gateway.UserAuthorization, player.LicenseCustomizeHandler)
=======
	server.engine.POST("/garage/buy", gateway.AdminAuthorization, player.BuyGarageHandler)
	server.engine.GET("/garages/get-all", admin.GetAllGarageListHandler)
	server.engine.POST("/garage/add-car", gateway.AdminAuthorization, player.AddCarToGarageHandler)
	server.engine.PUT("/garage/upgrade", gateway.AdminAuthorization, player.UpgradeGarageHandler)
	server.engine.GET("/garage/get", gateway.AdminAuthorization, player.GetPlayerGarageListHandler)

	//Admin garage routes
	server.engine.POST("/garage/add", gateway.AdminAuthorization, admin.AddGarageHandler)
	server.engine.DELETE("/garage/delete", gateway.AdminAuthorization, admin.DeleteGarageHandler)
	server.engine.PUT("/garage/update", gateway.AdminAuthorization, admin.UpdateGarageHandler)

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
>>>>>>> 43ea51cb16ea49a0369579929718436edab5da04

	server.engine.GET("/get-customization", player.GetCarCustomiseHandler)
	server.engine.GET("/get-color-category", player.GetCarColorCategoriesHandler)
	server.engine.GET("/get-color-type", player.GetCarColorTypesHandler)
	server.engine.GET("/get-colors", player.GetCarColorsHandler)

	//add dummy data in db
	server.engine.GET("/add-dummy-data", handler.AddDummyDataHandler)

	//swagger routes
	server.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
