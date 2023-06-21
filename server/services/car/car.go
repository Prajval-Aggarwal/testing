package car

import (
	"main/server/db"
	"main/server/model"
	"main/server/request"
	"main/server/response"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func BuyCarService(ctx *gin.Context, carRequest request.BuyCarRequest, playerId string) {

	var carDetails model.Car
	var playerDetails model.Player

	//check if the car exists or not
	if !db.RecordExist("car", carRequest.CarId, "car_id") {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_NOT_FOUND, utils.FAILURE, nil, ctx)
		return
	}
	err := db.FindById(&carDetails, carRequest.CarId, "car_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// get the details of the player from the database
	err = db.FindById(&playerDetails, playerId, "player_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//check for the coins and level required to unlock the car
	if playerDetails.Coins < int64(carDetails.CurrAmount) {
		if playerDetails.Level != int(carDetails.Level) {
			response.ShowResponse("Upgrade your level to unlock the car", utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		}
		response.ShowResponse("Not enough coins to unlock it", utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)

		return
	}

	//Add the car to players account and make it as players current selected car
	playerCar := model.PlayerCars{
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
	response.ShowResponse("Car added to player successfully", utils.HTTP_OK, utils.SUCCESS, playerCar, ctx)
}
