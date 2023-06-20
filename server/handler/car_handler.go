package handler

import (
	"main/server/request"
	"main/server/response"
	"main/server/services/car"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func AddCarHandler(ctx *gin.Context) {
	var carRequest request.AddCarRequest
	err := utils.RequestDecoding(ctx, &carRequest)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, "failure", nil, ctx)
		return
	}
	err = carRequest.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, "Failure", nil, ctx)
		return
	}
	car.AddCarService(ctx, carRequest)
}
