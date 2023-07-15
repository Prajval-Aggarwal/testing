package handler

import (
	"fmt"
	"main/server/request"
	"main/server/response"
	"main/server/services/arena"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Challenge Arena
// @Description Challenge the arena with the given request
// @Tags Arena
// @Accept json
// @Produce json
// @Param playerId header string true "The ID of the player"
// @Param challengereq body request.ChallengeReq true "Challenge Request"
// @Success 200 {object} response.Success "Success"
// @Failure 400 {object} response.Success "Bad request"
// @Failure  401 {object} response.Success "Unauthorised"
// @Router /arena/challenge [post]
func ChallengeArenaHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get("playerId")
	fmt.Println("player id is", playerId)
	if !exists {
		response.ShowResponse(utils.UNAUTHORIZED, utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}

	var challengeReq request.ChallengeReq
	err := utils.RequestDecoding(ctx, challengeReq)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	err = challengeReq.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	arena.ChallengeArenaService(ctx, challengeReq, playerId.(string))

}

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
// @Router /arena/end [post]
func EndChallengeHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get("playerId")
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
	arena.EndChallengeService(ctx, endChallReq, playerId.(string))
}

// func AddCarArenaHandler(ctx *gin.Context) {
// 	playerId, exists := ctx.Get("playerId")
// 	if !exists {

// 		response.ShowResponse(utils.UNAUTHORIZED, utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
// 		return
// 	}
// 	var addcarReq request.AddCarArenaRequest
// 	err := utils.RequestDecoding(ctx, &addcarReq)
// 	if err != nil {
// 		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
// 		return
// 	}

// 	err = addcarReq.Validate()
// 	if err != nil {
// 		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
// 		return
// 	}

// 	arena.AddCarArenaService(ctx, addcarReq, playerId.(string))

// }

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
// @Router /arena/replace-car [put]
func ReplaceArenaCarHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get("playerId")
	if !exists {

		response.ShowResponse(utils.UNAUTHORIZED, utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}

	var replaceReq request.ReplaceReq
	err := utils.RequestDecoding(ctx, &replaceReq)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	arena.ReplaceArenaService(ctx, replaceReq, playerId.(string))

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

// @Summary Get Arenas By id
// @Description Gets a particular arena
// @Tags Arena
// @Accept json
// @Produce json
// @Success 200 {object} response.Success "Success"
// @Failure 400 {object} response.Success "Bad request"
// @Router /arena/get-id [get]
func GetArenaByIdHandler(ctx *gin.Context) {
	var getReq request.GetArenaReq
	err := utils.RequestDecoding(ctx, &getReq)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	err = getReq.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	arena.GetArenaByIdService(ctx, getReq)

}
