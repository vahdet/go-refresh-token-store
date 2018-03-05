package services


import (
	log "github.com/sirupsen/logrus"
	"github.com/vahdet/go-refresh-token-store-redis/models"
	"gopkg.in/go-playground/validator.v9"
	"fmt"
	"github.com/vahdet/go-refresh-token-store-redis/dal/interfaces"
)

var validate *validator.Validate

type TokenService struct {
	dal interfaces.TokenDal
}

func NewTokenService(dal interfaces.TokenDal) *TokenService {
	return &TokenService{dal}
}

func (s *TokenService) Get(id int64) (*models.UserRefreshToken, error) {
	return s.dal.Get(id)
}

func (s *TokenService) Create(token *models.UserRefreshToken) (*models.UserRefreshToken, error) {
	// validation of the input
	if err := validate.Struct(token); err != nil {
		valErr := err.(validator.ValidationErrors)
		log.WithFields(log.Fields{
			"userid": token.UserId,
		}).Error(fmt.Sprintf("validation failed: '%#v'", valErr))
		return nil, err
	}
	// Data Access Layer call
	if err := s.dal.Create(token); err != nil {
		log.WithFields(log.Fields{
			"userid": token.UserId,
		}).Error(fmt.Sprintf("creation failed: '%#v'", err))
		return nil, err
	}
	return s.dal.Get(token.UserId)
}

func (s *TokenService) Update(id int64, token *models.UserRefreshToken) (*models.UserRefreshToken, error) {
	// validation of the input
	if err := validate.Struct(token); err != nil {
		valErr := err.(validator.ValidationErrors)
		log.WithFields(log.Fields{
			"userid": token.UserId,
		}).Error(fmt.Sprintf("validation failed: '%#v'", valErr))
		return nil, err
	}
	// Data Access Layer call
	if err := s.dal.Update(id, token); err != nil {
		log.WithFields(log.Fields{
			"id": id,
		}).Error(fmt.Sprintf("Update failed: '%#v'", err))
		return nil, err
	}
	return s.dal.Get(token.UserId)
}

func (s *TokenService) Delete(id int64) (*models.UserRefreshToken, error) {
	// Check if exists
	user, err := s.dal.Get(id)
	if err != nil {
		log.WithFields(log.Fields{
			"id": id,
		}).Error(fmt.Sprintf("getting failed: '%#v'", err))
		return nil, err
	}
	// Data Access Layer call
	err = s.dal.Delete(id)
	return user, err
}

func (s *TokenService) Count(id int64) (int64, error) {
	return s.dal.Count(id)
}
