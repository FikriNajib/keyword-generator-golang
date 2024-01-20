package request

type CalculateRequest struct {
	UserID int `json:"user_id" validate:"required"`
}
