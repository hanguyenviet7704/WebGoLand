package po

import "time"

type Tokens struct {
	ID                       int       `json:"id"`
	UserID                   int       `json:"user_id"`
	AccessToken              string    `json:"token"`
	RefreshToken             string    `json:"refresh_token"`
	Access_token_expires_at  time.Time `json:"access_token_expires_at"`
	Refresh_token_expires_at time.Time `json:"refresh_token_expires_at"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}
