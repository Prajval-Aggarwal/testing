package handler

import (
	"main/server/request"
	"main/server/response"
	"main/server/services/arena"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

// @Summary End Challenge
// @Description Ends the current challenge and saves the data
// @Tags Arena
// @Accept json
// @Produce json
// @Param playerId header string true "The ID of the player"
// @Param challengereq body request.EndChallengeReq true "End Challenge Request"
// @Success 200 {object} response.Success "Success"
// @Failure 400 {object} response.Success "Bad request"
// @Failure  401 {object} response.Success "Unauthorised"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /arena/end [post]
func EndChallengeHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get(utils.PLAYER_ID)
	if !exists {

		response.ShowResponse(utils.UNAUTHORIZED, utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}
	var endChallReq request.EndChallengeReq
	err := utils.RequestDecoding(ctx, &endChallReq)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	err = endChallReq.Validate()
	if err != nil {

		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	arena.EndChallengeService(ctx, endChallReq, playerId.(string))
}

// @Summary Get Arenas
// @Description Gets the list of all the arenas
// @Tags Arena
// @Accept json
// @Produce json
// @Success 200 {object} response.Success "Success"
// @Failure 400 {object} response.Success "Bad request"
// @Router /arena/get [get]
func GetArenaHandler(ctx *gin.Context) {
	arena.GetArenaService(ctx)
}

// AddCarToSlotHandler adds a car to the player's slot in a specific arena.
// @Summary Add a car to slot
// @Description Add a car to the player's slot in a specific arena
// @Tags Arena
// @Accept json
// @Produce json
// @Param Authorization header string true "Player Access token"
// @Param addCarReq body request.AddCarArenaRequest true "Add car to slot request payload"
// @Success 200 {object} response.Success "Car added to slot successfully"
// @Failure 400 {object} response.Success "Bad request. Invalid payload"
// @Failure 404 {object} response.Success "Car or player not found"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /arena/add-car [post]
func AddCarToSlotHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get(utils.PLAYER_ID)
	if !exists {
		response.ShowResponse(utils.UNAUTHORIZED, utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}
	var addCarReq request.AddCarArenaRequest
	err := utils.RequestDecoding(ctx, &addCarReq)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	arena.AddCarToSlotService(ctx, addCarReq, playerId.(string))

}

// @Summary Replace Car
// @Description Add or replaces the car in the arena car slot
// @Tags Arena
// @Accept json
// @Produce json
// @Param playerId header string true "The ID of the player"
// @Param challengereq body request.ReplaceReq true "Replace car Request"
// @Success 200 {object} response.Success "Success"
// @Failure 400 {object} response.Success "Bad request"
// @Failure  401 {object} response.Success "Unauthorised"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /arena/replace-car [put]
func ReplaceCarHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get(utils.PLAYER_ID)
	if !exists {
		response.ShowResponse(utils.UNAUTHORIZED, utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}
	var addCarReq request.AddCarArenaRequest
	err := utils.RequestDecoding(ctx, &addCarReq)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	arena.AddCarToSlotService(ctx, addCarReq, playerId.(string))
}
