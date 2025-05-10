package dto

type CreateUpdateCountryRequest struct {
	Name string `json:"name" binding:"required,min=3"`
}

type CountryResponse struct {
	Id     int            `json:"id"`
	Name   string         `json:"name"`
	Cities []CityResponse `json:"cities"`
}

type GetByIdResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CreateUpdateCityRequest struct {
	Name string `json:"name" binding:"required,min=3"`
}

type CityResponse struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Country CountryResponse `json:"country"`
}
