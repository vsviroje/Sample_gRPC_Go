package models

type User struct {
	ID      int     `json:"user_id"`
	FName   string  `json:"first_name"`
	City    string  `json:"city"`
	Height  float32 `json:"height"`
	Phone   int32   `json:"phone"`
	Married bool    `json:"married"`
}
