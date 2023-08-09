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

	//extra routes
	server.engine.POST("/admin/signup", admin.AdminSignUpHandler)
	server.engine.POST("/add-garage-type", admin.AddGarageTypesHandler)
	server.engine.POST("/add-arena-type", admin.AddArenaTypesHandler)

	//Admin Login Route

	server.engine.POST("/forgot-password", admin.ForgotPasswordHandler)
	server.engine.PATCH("/reset-password", admin.ResetPasswordHandler)

	//Auth routes
	server.engine.POST("/guest-login", admin.GuestLoginHandler)
	server.engine.POST("/login", admin.LoginHandler)
	server.engine.PUT("/update-email", gateway.UserAuthorization, admin.UpdateEmailHandler)
	server.engine.PATCH("/update-pass", gateway.AdminAuthorization, admin.UpdatePasswordHandler)
	server.engine.GET("/admin", admin.GetAdminHandler)

	//player details
	server.engine.GET("/player-details", gateway.UserAuthorization, handler.GetPlayerDetailsHandler)

	//Player routes
	server.engine.POST("/car/buy", gateway.UserAuthorization, player.BuyCarHandler)
	server.engine.PUT("/car/equip", gateway.UserAuthorization, player.EquipCarHandler)
	server.engine.DELETE("/car/sell", gateway.UserAuthorization, player.SellCarHandler)
	server.engine.POST("/car/repair", gateway.UserAuthorization, player.RepairCarHandler)
	server.engine.GET("/car/get-all", player.GetAllCarsHandler)
	server.engine.POST("/car/get-by-id", player.GetCarByIdHandler)

	//Player garage routes
	server.engine.POST("/garage/buy", gateway.UserAuthorization, player.BuyGarageHandler)
	server.engine.GET("/garages/get-all", admin.GetAllGarageListHandler)
	server.engine.POST("/garage/add-car", gateway.UserAuthorization, player.AddCarToGarageHandler)
	server.engine.PUT("/garage/upgrade", gateway.UserAuthorization, player.UpgradeGarageHandler)
	server.engine.GET("/garage/get", gateway.UserAuthorization, player.GetPlayerGarageListHandler)

	//Admin garage routes
	server.engine.POST("/admin/garage/add", gateway.AdminAuthorization, admin.AddGarageHandler)
	server.engine.DELETE("/admin/garage/delete", gateway.AdminAuthorization, admin.DeleteGarageHandler)
	server.engine.PUT("/admin/garage/update", gateway.AdminAuthorization, admin.UpdateGarageHandler)
	server.engine.GET("/garage/types", admin.GetGarageTypesHandler)

	//Admin Battle Arena Routes
	server.engine.POST("/admin/arena", gateway.AdminAuthorization, admin.AddArenaHandler)
	server.engine.DELETE("/admin/arena", gateway.AdminAuthorization, admin.DeleteArenaHandler)
	server.engine.PUT("/admin/arena", gateway.AdminAuthorization, admin.UpdateArenaHandler)
	server.engine.GET("/arena/get", admin.GetArenaListHandler)
	server.engine.GET("/arena/types", admin.GetArenaTypeHandler)

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

	server.engine.GET("/get-customization", player.GetCarCustomiseHandler)
	server.engine.GET("/get-color-category", player.GetCarColorCategoriesHandler)
	server.engine.GET("/get-color-type", player.GetCarColorTypesHandler)
	server.engine.GET("/get-colors", player.GetCarColorsHandler)

	//add dummy data in db
	server.engine.GET("/add-dummy-data", handler.AddDummyDataHandler)

	//swagger routes
	server.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
