package domain

import "time"

type AccessTokenResult struct {
	AccessToken string `json:"access_token"`
	UserID int64 `json:"user_id"`
	ClientID int64 `json:"client_id"`
	ExpiresString string `json:"expires_string"`
	ExpiresDateTime time.Time `json:"expires_time"`
	ExpiresInt int64 `json:"expires_int"`
	IsAlreadyExpired bool `json:"is_already_expired"",omitempty`
}
