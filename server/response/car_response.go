package response

import "main/server/model"

type CarResponse struct {
	CarName       string                       `json:"carName"`
	Car           model.Car                    `json:"info"`
	CarStats      model.CarStats               `json:"stats"`
	Customization []model.DefaultCustomization `json:"customization"`
}
