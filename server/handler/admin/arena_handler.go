package handler

import (
	"main/server/request"
	"main/server/response"
	"main/server/services/arena"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Add a new arena
// @Description Add a new arena to the system
// @Tags Arena
// @Accept json
// @Produce json
// @Param Authorization header string true "Admin Access token"
// @Param garageReq body request.AddArenaRequest true "Arena request payload"
// @Success 200 {object} response.Success "Arena added successful"
// @Failure 400 {object} response.Success "Bad request"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /admin/arena/add [post]
func AddArenaHandler(ctx *gin.Context) {

	var addArenaReq request.AddArenaRequest
	err := utils.RequestDecoding(ctx, &addArenaReq)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	err = addArenaReq.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	arena.AddArenaService(ctx, addArenaReq)
}

// DeleteArenaHandler deletes a arena with the given ID.
// @Summary Delete a Arena
// @Description Delete a Arena by its ID
// @Tags Arena
// @Accept json
// @Produce json
// @Param Authorization header string true "Admin Access token"
// @Param ArenaReq body request.DeletArenaReq true "Arena request payload"
// @Success 200 {object} response.Success "Arena deleted successful"
// @Failure 400 {object} response.Success "Bad request"
// @Failure 404 {string} string "Arena not found"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /admin/arena/delete [delete]
func DeleteArenaHandler(ctx *gin.Context) {
	var deleteReq request.DeletArenaReq
	err := utils.RequestDecoding(ctx, &deleteReq)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	err = deleteReq.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	arena.DeleteArenaService(ctx, deleteReq)
}

// UpdateArenaHandler updates a Arena with the given ID.
// @Summary Update a Arena
// @Description Update a Arena by its ID
// @Tags Arena
// @Accept json
// @Produce json
// @Param Authorization header string true "Admin Access token"
// @Param updateReq body request.UpdateArenaReq true "Update request payload"
// @Success 200 {object} response.Success "Arena updated successful"
// @Failure 400 {object} response.Success "Bad request"
// @Failure 404 {string} string "Arena not found"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /admin/arena/update [put]
func UpdateArenaHandler(ctx *gin.Context) {
	var updateReq request.UpdateArenaReq
	err := utils.RequestDecoding(ctx, &updateReq)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	err = updateReq.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	arena.UpdateArenaService(ctx, updateReq)
}

// GetArenaListHandler retrieves the list of all Arenas.
//
// @Summary Get All Arena List
// @Description Retrieve the list of all Arenas
// @Tags Arena
// @Accept json
// @Produce json
// @Param skip query integer false "Number of records to skip (default is 0)"
// @Param limit query integer false "Maximum number of records to fetch (default is 10)"
// @Success 200 {object} response.Success "Garage list fetched successfully"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /arena/get-all [get]
func GetArenaListHandler(ctx *gin.Context) {
	arena.GetAllArenaService(ctx)
}
