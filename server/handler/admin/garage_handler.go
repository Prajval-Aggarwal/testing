package handler

import (
	"main/server/request"
	"main/server/response"
	"main/server/services/garage"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Add a new garage
// @Description Add a new garage to the system
// @Tags Garage
// @Accept json
// @Produce json
// @Param Authorization header string true "Admin Access token"
// @Param garageReq body request.AddGarageRequest true "Garage request payload"
// @Success 200 {object} response.Success "Garage added successful"
// @Failure 400 {object} response.Success "Bad request"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /admin/garage/add [post]
func AddGarageHandler(ctx *gin.Context) {
	var addGarageReq request.AddGarageRequest
	err := utils.RequestDecoding(ctx, &addGarageReq)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	err = addGarageReq.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	garage.AddGarageService(ctx, addGarageReq)
}

// DeleteGarageHandler deletes a garage with the given ID.
// @Summary Delete a garage
// @Description Delete a garage by its ID
// @Tags Garage
// @Accept json
// @Produce json
// @Param Authorization header string true "Admin Access token"
// @Param garageReq body request.DeletGarageReq true "Garage request payload"
// @Success 200 {object} response.Success "Garage deleted successful"
// @Failure 400 {object} response.Success "Bad request"
// @Failure 404 {string} response.Success "Garage not found"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /admin/garage/delete [delete]
func DeleteGarageHandler(ctx *gin.Context) {
	var deleteReq request.DeletGarageReq
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

	garage.DeleteGarageService(ctx, deleteReq)
}

// UpdateGarageHandler updates a garage with the given ID.
// @Summary Update a garage
// @Description Update a garage by its ID
// @Tags Garage
// @Accept json
// @Produce json
// @Param Authorization header string true "Admin Access token"
// @Param updateReq body request.UpdateGarageReq true "Update request payload"
// @Success 200 {object} response.Success "Garage updated successful"
// @Failure 400 {object} response.Success "Bad request"
// @Failure 404 {string} response.Success "Garage not found"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /admin/garage/update [put]
func UpdateGarageHandler(ctx *gin.Context) {
	var updateReq request.UpdateGarageReq
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

	garage.UpdateGarageService(ctx, updateReq)
}

// GetAllGarageListService retrieves the list of all garages.
//
// @Summary Get All Garage List
// @Description Retrieve the list of all garages
// @Tags Garage
// @Accept json
// @Produce json
// @Param skip query integer false "Number of records to skip (default is 0)"
// @Param limit query integer false "Maximum number of records to fetch (default is 10)"
// @Success 200 {object} response.Success "Garage list fetched successfully"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /garages/get-all [get]
func GetAllGarageListHandler(ctx *gin.Context) {
	garage.GetAllGarageListService(ctx)
}

// just for backend not for front end so no swagger
func AddGarageTypesHandler(ctx *gin.Context) {
	garage.AddGargeTypes(ctx)
}

// GetGarageTypesHandler retrieves the list of all garages.
//
// @Summary Get All Garage type List
// @Description Retrieve the list of all garages
// @Tags Garage
// @Accept json
// @Produce json
// @Success 200 {object} response.Success "Garage type list fetched successfully"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /garage/types [get]
func GetGarageTypesHandler(ctx *gin.Context) {
	garage.GetGarageTypes(ctx)
}
