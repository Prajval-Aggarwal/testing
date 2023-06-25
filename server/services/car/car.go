package car

import (
	"main/server/db"
	"main/server/model"
	"main/server/request"
	"main/server/response"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func EquipCarService(ctx *gin.Context, equipRequest request.CarRequest, playerId string) {

	query := "Update owned_cars SET selected=false WHERE player_id=? AND selected=true"
	err := db.RawExecutor(query, playerId)
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

	response.ShowResponse("Current car selected successfully", utils.HTTP_OK, utils.SUCCESS, nil, ctx)

}
func BuyCarService(ctx *gin.Context, carRequest request.CarRequest, playerId string) {

	var carDetails model.Car
	var playerDetails model.Player

	//check if the car exists or not
	if !db.RecordExist("cars", carRequest.CarId, "car_id") {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_NOT_FOUND, utils.FAILURE, nil, ctx)
		return
	}
	err := db.FindById(&carDetails, carRequest.CarId, "car_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	if !carDetails.Locked {
		response.ShowResponse("Car already bought", utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
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
		if playerDetails.Level <= int(carDetails.Level) {
			response.ShowResponse(utils.UPGRADE_LEVEL, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		}
		response.ShowResponse(utils.NOT_ENOUGH_COINS, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)

		return
	}

	//Add the car to players account and make it as players current selected car
	playerCar := model.OwnedCars{
		PlayerId: playerId,
		CarId:    carRequest.CarId,
		Selected: true,
	}

	query := "UPDATE cars SET locked=false WHERE car_id=?"
	err = db.RawExecutor(query, carRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	err = db.CreateRecord(&playerCar)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	response.ShowResponse(utils.CAR_BOUGHT_SUCESS, utils.HTTP_OK, utils.SUCCESS, playerCar, ctx)
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
		playerDetails.Coins += int64(0.6 * float64(carDetails.CurrAmount))
	}

	err = db.UpdateRecord(&playerDetails, playerId, "player_id").Error
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	//delete the car from players collection
	query := "DELETE FROM owned_cars WHERE car_id =? AND player_id =?"
	err = db.RawExecutor(query, sellCarRequest.CarId, playerId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse(utils.CAR_SOLD_SUCCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)

}
