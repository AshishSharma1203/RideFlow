package auth

type Claims struct {
	UserID string `json:"user_id"`
	Expiry int64 `json:"expiry"`
}