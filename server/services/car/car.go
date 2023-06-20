package car

import (
	"main/server/db"
	"main/server/model"
	"main/server/request"
	"main/server/response"
	"main/server/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AddCarService(ctx *gin.Context, carRequest request.AddCarRequest) {
	carUuid := uuid.New()
	carModel := model.Car{
		CarId:      carUuid,
		CarName:    carRequest.CarName,
		Level:      carRequest.Level,
		CurrType:   carRequest.CurrType,
		CurrAmount: carRequest.CurrAmount,
		MaxLevel:   carRequest.MaxLevel,
		Class:      carRequest.Class,
		Status:     "locked",
	}

	carStats := model.CarStats{
		CarId:      carUuid,
		Power:      carRequest.Power,
		Grip:       carRequest.Grip,
		ShiftTime:  carRequest.ShiftTime,
		Weight:     carRequest.Weight,
		OR:         carRequest.OR,
		Durability: carRequest.Durability,
	} 

	err := db.CreateRecord(&carModel)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, "failure", nil, ctx)
		return
	}

	err = db.CreateRecord(&carStats)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, "failure", nil, ctx)
		return
	}

	response.ShowResponse(
		"Cars added successfully",
		utils.HTTP_OK,
		"success",
		carStats,
		ctx,
	)

}
