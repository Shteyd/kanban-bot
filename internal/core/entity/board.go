package entity

type Board struct {
	ID          string `json:"id"          example:"d50d5ef5-c9c0-4ae4-a6b4-e1289f3bb0e0"`
	CustomerID  string `json:"customer_id" example:"d50d5ef5-c9c0-4ae4-a6b4-e1289f3bb0e0"`
	Article     string `json:"article"     example:"WEB"`
	Description string `json:"description" example:"Some text"`
}
