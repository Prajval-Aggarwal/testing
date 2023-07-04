package model

type CarCustomization struct {
	Part          string  `json:"part,omitempty"`
	ColorCategory string  `json:"colorCategory,omitempty"`
	ColorType     string  `json:"colorType,omitempty"`
	ColorCode     string  `json:"colorCode,omitempty"`
	ColorName     string  `json:"colorName,omitempty"`
	CurrType      string  `json:"currType,omitempty"`
	CurrAmount    float64 `json:"currAmount,omitempty"`
	Value         string  `json:"value,omitempty"`
}

type DefaultCustomization struct {
	CarId         string  `json:"carId,omitempty"`
	Part          string  `json:"part,omitempty"`
	ColorCategory string  `json:"colorCategory,omitempty"`
	ColorType     string  `json:"colorType,omitempty"`
	ColorCode     string  `json:"colorCode,omitempty"`
	ColorName     string  `json:"colorName,omitempty"`
	CurrType      string  `json:"currType,omitempty"`
	CurrAmount    float64 `json:"currAmount,omitempty"`
	Value         string  `json:"value,omitempty"`
}
