package dto

type NewProduct struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type UpdateProduct struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
