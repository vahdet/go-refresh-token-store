package interfaces

import "github.com/vahdet/go-refresh-token-store-redis/models"
type (
	UserService interface {
		Get(id int64) (*models.UserRefreshToken, error)
		Create(user *models.UserRefreshToken) (*models.UserRefreshToken, error)
		Update(id int64, user *models.UserRefreshToken) (*models.UserRefreshToken, error)
		Delete(id int64) (*models.UserRefreshToken, error)
		Count(id int64) (int64, error)
	}
)
