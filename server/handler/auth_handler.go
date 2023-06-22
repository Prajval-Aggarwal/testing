package handler

import (
	"main/server/request"
	"main/server/response"
	"main/server/services/auth"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

// @Description	Guest Log in a player
// @Accept			json
// @Produce		json
// @Success		200				{object}	response.Success
// @Failure		400				{object}	response.Success
// @Param			playerDetails	body		request.GuestLoginRequest	true	"Device id of the player"
// @Tags			Authentication
// @Router			/guest-login [post]
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

// @Description	Log in a player
// @Accept			json
// @Produce		json
// @Success		200				{object}	response.Success
// @Failure		400				{object}	response.Success
// @Param			playerDetails	body		request.LoginRequest	true	"Details of the player"
// @Tags			Authentication
// @Router			/login [post]
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

// @Description	Update Email of the player
// @Accept			json
// @Produce		json
// @Success		200				{object}	response.Success
// @Failure		400				{object}	response.Success
// @Failure		401				{object}	response.Success
// @Param			playerDetails	body	request.UpdateEmailRequest	true	"Email of the player"

// @Tags			Authentication
// @Router			/update-email [put]
func UpdateEmailHandler(ctx *gin.Context) {
	var email request.UpdateEmailRequest

	token := ctx.Request.Header.Get("Authorization")
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
	auth.UpdateEmailService(ctx, email, token)
}
