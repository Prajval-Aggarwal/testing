package handler

import (
	"main/server/request"
	"main/server/response"
	"main/server/services/arena"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

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
