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
	server.engine.POST("/car/buy", gateway.AdminAuthorization, player.BuyCarHandler)
	server.engine.PUT("/car/equip", gateway.AdminAuthorization, player.EquipCarHandler)
	server.engine.DELETE("/car/sell", gateway.AdminAuthorization, player.SellCarHandler)
	server.engine.POST("/car/repair", gateway.AdminAuthorization, player.RepairCarHandler)
	server.engine.GET("/car/get-all", player.GetAllCarsHandler)
	server.engine.POST("/car/get-by-id", player.GetCarByIdHandler)

	//Player garage routes
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
	server.engine.PUT("/car/upgrade/engine", gateway.AdminAuthorization, player.UpgradeEngineHandler)
	server.engine.PUT("/car/upgrade/turbo", gateway.AdminAuthorization, player.UpgradeTurboHandler)
	server.engine.PUT("/car/upgrade/intake", gateway.AdminAuthorization, player.UpgradeIntakeHandler)
	server.engine.PUT("/car/upgrade/nitrous", gateway.AdminAuthorization, player.UpgradeNitrousHandler)
	server.engine.PUT("/car/upgrade/body", gateway.AdminAuthorization, player.UpgradeBodyHandler)
	server.engine.PUT("/car/upgrade/tires", gateway.AdminAuthorization, player.UpgradeTiresHandler)
	server.engine.PUT("/car/upgrade/transmission", gateway.AdminAuthorization, player.UpgradeTransmissionHandler)

	//car Customiztion routes
	server.engine.PUT("/car/customise/color", gateway.AdminAuthorization, player.ColorCustomizeHandler)
	server.engine.PUT("/car/customise/wheels", gateway.AdminAuthorization, player.WheelsCustomizeHandler)
	server.engine.PUT("/car/customise/interior", gateway.AdminAuthorization, player.InteriorCustomizeHandler)
	server.engine.PUT("/car/customise/license", gateway.AdminAuthorization, player.LicenseCustomizeHandler)

	server.engine.GET("/get-customization", player.GetCarCustomiseHandler)
	server.engine.GET("/get-color-category", player.GetCarColorCategoriesHandler)
	server.engine.GET("/get-color-type", player.GetCarColorTypesHandler)
	server.engine.GET("/get-colors", player.GetCarColorsHandler)

	//Arena Routes
	server.engine.POST("/arena/end", gateway.AdminAuthorization, handler.EndChallengeHandler)
	server.engine.POST("/arena/add-car", gateway.AdminAuthorization, handler.AddCarToSlotHandler)
	server.engine.POST("/arena/replace-car", gateway.AdminAuthorization, handler.ReplaceCarHandler)
	server.engine.GET("/arena/get", handler.GetArenaHandler)

	//add dummy data in db
	server.engine.GET("/add-dummy-data", handler.AddDummyDataHandler)

	//swagger routes
	server.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
