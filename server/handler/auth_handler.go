package handler

import (
	"fmt"
	"main/server/request"
	"main/server/response"
	"main/server/services/auth"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

// GuestLoginService handles guest login and token generation.
//
// @Summary Guest Login
// @Description Perform guest login and generate access token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param guestLoginRequest body request.GuestLoginRequest true "Guest Login Request"
// @Success 200 {object} response.Success "Login successful"
// @Failure 400 {object} response.Success "Bad request"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /guest-login [post]
func GuestLoginHandler(ctx *gin.Context) {
	var guestLoginReq request.GuestLoginRequest
	err := utils.RequestDecoding(ctx, &guestLoginReq)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, "Failure", nil, ctx)
		return
	}
	err = guestLoginReq.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, "Failure", nil, ctx)
		return
	}
	auth.GuestLoginService(ctx, guestLoginReq)
}

// LoginService handles user login and token generation.
//
// @Summary User Login
// @Description Perform user login and generate access token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param loginDetails body request.LoginRequest true "Login Details"
// @Success 200 {object} response.Success "Login successful"
// @Failure 400 {object} response.Success "Bad request"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /login [post]
func LoginHandler(ctx *gin.Context) {
	var loginDetails request.LoginRequest
	err := utils.RequestDecoding(ctx, &loginDetails)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, "Failure", nil, ctx)
		return
	}
	err = loginDetails.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, "Failure", nil, ctx)
		return
	}

	auth.LoginService(ctx, loginDetails)
}

// UpdateEmailService updates the email of a player.
//
// @Summary Update Email
// @Description Update the email address of a player
// @Tags Player
// @Accept json
// @Produce json
// @Param email body request.UpdateEmailRequest true "Update Email Request"
// @Param Authorization header string true "Access Token"
// @Success 200 {object} response.Success "Email updated successfully"
// @Failure 400 {object} response.Success "Bad request"
// @Failure 401 {object} response.Success "Unauthorized"
// @Failure 404 {object} response.Success "Player not found"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /update-email [put]
func UpdateEmailHandler(ctx *gin.Context) {
	var email request.UpdateEmailRequest
	playerId, exists := ctx.Get("playerId")
	fmt.Println("player id is:", playerId, exists)
	if !exists {
		response.ShowResponse(utils.UNAUTHORIZED, utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}
	err := utils.RequestDecoding(ctx, &email)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, "Failure", nil, ctx)
		return
	}
	err = email.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, "Failure", nil, ctx)
		return
	}
	auth.UpdateEmailService(ctx, email, playerId.(string))
}
