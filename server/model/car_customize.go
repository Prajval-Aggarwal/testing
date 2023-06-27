package model

type CarCustomization struct {
	Part          string  `json:"part"`
	ColorCategory string  `json:"colorCategory"`
	ColorType     string  `json:"colorType"`
	ColorCode     string  `json:"colorCode"`
	ColorName     string  `json:"colorName"`
	CurrType      string  `json:"currType"`
	CurrAmount    float64 `json:"currAmount"`
	Value         string  `json:"value"`
}

type DefaultCustomization struct {
	CarId         string  `json:"carId"`
	Part          string  `json:"part"`
	ColorCategory string  `json:"colorCategory"`
	ColorType     string  `json:"colorType"`
	ColorCode     string  `json:"colorCode"`
	ColorName     string  `json:"colorName"`
	CurrType      string  `json:"currType"`
	CurrAmount    float64 `json:"currAmount"`
	Value         string  `json:"value"`
}
