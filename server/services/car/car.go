package car

import (
	"fmt"
	"main/server/db"
	"main/server/model"
	"main/server/request"
	"main/server/response"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func EquipCarService(ctx *gin.Context, equipRequest request.CarRequest, playerId string) {

	//check if the car is bought or not

	var exists bool
	query := "SELECT EXISTS(SELECT * FROM owned_cars WHERE player_id =? AND car_id=?)"
	err := db.QueryExecutor(query, &exists, playerId, equipRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	if !exists {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_NOT_FOUND, utils.FAILURE, nil, ctx)
		return
	}

	query = "Update owned_cars SET selected=false WHERE player_id=? AND selected=true"
	err = db.RawExecutor(query, playerId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	query = "UPDATE owned_cars SET selected=true WHERE player_id=? AND car_id=?"
	err = db.RawExecutor(query, playerId, equipRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse(utils.CAR_SELECETED_SUCCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)

}

func BuyCarService(ctx *gin.Context, carRequest request.CarRequest, playerId string) {

	var carDetails model.Car
	var playerDetails model.Player

	fmt.Println("car details is", carRequest.CarId)
	//if player already has that car
	var exists bool
	query := "SELECT EXISTS(SELECT * FROM owned_cars WHERE player_id =? AND car_id =?)"
	err := db.QueryExecutor(query, &exists, playerId, carRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	if exists {
		response.ShowResponse(utils.CAR_ALREADY_BOUGHT, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)

		return
	}

	//check if the car exists or not
	if !db.RecordExist("cars", carRequest.CarId, "car_id") {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_NOT_FOUND, utils.FAILURE, nil, ctx)
		return
	}
	err = db.FindById(&carDetails, carRequest.CarId, "car_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	// if !carDetails.Locked {
	// 	response.ShowResponse("Car already bought", utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
	// 	return
	// }
	// get the details of the player from the database
	err = db.FindById(&playerDetails, playerId, "player_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	var amount int64
	var currType string
	//check for the coins and level required to unlock the car
	if carDetails.CurrType == "coins" {
		currType = "coins"
		amount = playerDetails.Coins
	} else {
		currType = "cash"
		amount = playerDetails.Cash
	}

	if amount < int64(carDetails.CurrAmount) && currType == carDetails.CurrType {
		response.ShowResponse(utils.NOT_ENOUGH_COINS, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//deduct player money
	if carDetails.CurrType == "coins" {
		playerDetails.Coins -= int64(carDetails.CurrAmount)
	} else {
		playerDetails.Cash -= int64(carDetails.CurrAmount)
	}

	err = db.UpdateRecord(&playerDetails, playerId, "player_id").Error
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//Add the car to players account and make it as players current selected car
	playerCar := model.OwnedCars{
		PlayerId: playerId,
		CarId:    carRequest.CarId,
		Selected: true,
	}

	//set bought car defaults
	err = utils.SetPlayerCarDefaults(playerId, carRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	//creae player record
	err = db.CreateRecord(&playerCar)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse(utils.CAR_BOUGHT_SUCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)
}

func SellCarService(ctx *gin.Context, sellCarRequest request.CarRequest, playerId string) {

	if !db.RecordExist("owned_cars", sellCarRequest.CarId, "car_id") {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_NOT_FOUND, utils.FAILURE, nil, ctx)
		return
	}

	var playerDetails model.Player
	var carDetails model.Car
	err := db.FindById(&playerDetails, playerId, "player_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	err = db.FindById(&carDetails, sellCarRequest.CarId, "car_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//give some reward to user on selling his car

	//if buying cuurency of car is coins then give 30% of original buying amount else give 60% of original buying amount
	if carDetails.CurrType == "coins" {
		playerDetails.Coins += int64(0.3 * float64(carDetails.CurrAmount))
	} else {
		playerDetails.Cash += int64(0.6 * float64(carDetails.CurrAmount))
	}

	err = db.UpdateRecord(&playerDetails, playerId, "player_id").Error
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	//delete the car from players collection

	err = utils.DeleteCarDetails("owned_cars", playerId, sellCarRequest.CarId, ctx)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	err = utils.DeleteCarDetails("player_cars_stats", playerId, sellCarRequest.CarId, ctx)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	err = utils.DeleteCarDetails("player_car_upgrades", playerId, sellCarRequest.CarId, ctx)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	response.ShowResponse(utils.CAR_SOLD_SUCCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)

}

func RepairCarService(ctx *gin.Context, repairCarRequest request.CarRequest, playerId string) {

	var playerDetails model.Player
	err := db.FindById(&playerDetails, playerId, "player_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	var carDetails model.CarStats
	err = db.FindById(&carDetails, repairCarRequest.CarId, "car_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	var playerCarStats model.PlayerCarsStats
	query := "SELECT * FROM player_cars_stats WHERE car_id=? AND player_id=?"
	err = db.QueryExecutor(query, &playerCarStats, repairCarRequest.CarId, playerId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	durabilityDiff := carDetails.Durability - playerCarStats.Durability

	if durabilityDiff*5 > playerDetails.RepairParts {
		response.ShowResponse(utils.NOT_ENOUGH_REPAIR_PARTS, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	tx := db.BeginTransaction()
	if tx.Error != nil {
		response.ShowResponse(tx.Error.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	playerCarStats.Durability = carDetails.Durability
	query = "UPDATE player_cars_stats SET durability = ? WHERE player_id=? AND car_id=?"
	err = tx.Exec(query, carDetails.Durability, playerId, repairCarRequest.CarId).Error
	if err != nil {
		tx.Rollback() // Rollback the transaction if there's an error
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	if err := tx.Commit().Error; err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	response.ShowResponse(utils.CAR_REPAIR_SUCCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)
}

func GetAllCarsService(ctx *gin.Context) {
	var carDetails []model.Car
	query := "SELECT * FROM cars"
	err := db.QueryExecutor(query, &carDetails)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse(utils.DATA_FETCH_SUCCESS, utils.HTTP_OK, utils.SUCCESS, carDetails, ctx)
}

func GetCarByIdService(ctx *gin.Context, getReq request.CarRequest) {
	var carResponse response.CarResponse
	var carDetails model.Car
	var carStats model.CarStats
	var customization []model.DefaultCustomization

	//check if the car is present with given car id or not

	if !db.RecordExist("cars", getReq.CarId, "car_id") {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_NOT_FOUND, utils.FAILURE, nil, ctx)
		return
	}

	query := "SELECT * FROM cars WHERE car_id=?"
	err := db.QueryExecutor(query, &carDetails, getReq.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	query = "SELECT * FROM default_customizations WHERE car_id=?"
	err = db.QueryExecutor(query, &customization, getReq.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	query = "SELECT * FROM car_stats WHERE car_id=?"
	err = db.QueryExecutor(query, &carStats, getReq.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	carResponse.CarName = carDetails.CarName
	carResponse.Car = carDetails
	carResponse.CarStats = carStats
	carResponse.Customization = customization

	response.ShowResponse(utils.DATA_FETCH_SUCCESS, utils.HTTP_OK, utils.SUCCESS, carResponse, ctx)
}
