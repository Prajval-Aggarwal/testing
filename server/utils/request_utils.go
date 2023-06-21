package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func RequestDecoding(ctx *gin.Context, data interface{}) error {

	reqBody, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(reqBody, &data)
	if err != nil {
		return err
	}
	return nil
}
