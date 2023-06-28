package request

type LicenseRequest struct {
	CarId string `json:"carId"`
	Value string `json:"value"`
}
type InteriorCustomizeRequest struct {
	CarId     string `json:"carId"`
	ColorName string `json:"colorName"`
	ColorCode string `json:"colorCode"`
}

type WheelCustomizeRequest struct {
	CarId         string `json:"carId"`
	ColorCategory string `json:"colorCategory"`
	ColorName     string `json:"colorName"`
	ColorCode     string `json:"colorCode"`
}

type ColorCustomizationRequest struct {
	CarId         string `json:"carId"`
	ColorCategory string `json:"colorCategory"`
	ColorType     string `json:"colorType"`
	ColorName     string `json:"colorName"`
	ColorCode     string `json:"colorCode"`
}
