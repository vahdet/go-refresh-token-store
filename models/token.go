package models

// structs
type (
	UserRefreshToken struct {
		UserId     		int64     	`json:"id"`
		RefreshToken 	string		`json:"refreshToken"`
	}
)