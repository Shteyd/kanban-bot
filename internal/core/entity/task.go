package entity

type Task struct {
	ID          string `json:"id"          example:"d50d5ef5-c9c0-4ae4-a6b4-e1289f3bb0e0"`
	BoardID     string `json:"board_id"    example:"d50d5ef5-c9c0-4ae4-a6b4-e1289f3bb0e0"`
	Position    int    `json:"position"    example:"1"`
	Name        string `json:"name"        example:"Create new idea"`
	Description string `json:"description" example:"Some text"`
}
