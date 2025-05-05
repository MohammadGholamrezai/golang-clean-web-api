package dto

type CreateUpdateCountryRequest struct {
	Name string `json:"name" binding:"required,min=3"`
}

type CreateUpdateCountryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type GetByIdResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
