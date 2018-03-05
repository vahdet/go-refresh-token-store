package dal

import (
	"github.com/vahdet/go-refresh-token-store-redis/models"
	log "github.com/sirupsen/logrus"
	"fmt"
	"github.com/go-redis/redis"
)

const (
	tokenKeyPrefix = "refreshtoken"
	keySeparator = ":"
)

type TokenDal struct{}

func NewTokenDal() *TokenDal {
	return &TokenDal{}
}

func (dal *TokenDal) Get(id int64) (*models.UserRefreshToken, error) {

	var userRefreshToken models.UserRefreshToken

	userRefreshToken.UserId = id

	tokenFound, err := client.Get(getPrefixedDataStoreId(id)).Result()

	if err == redis.Nil {
		userRefreshToken.RefreshToken = ""
	} else if err != nil {
		log.WithFields(log.Fields{
			"id": id,
		}).Error(fmt.Sprintf("getting failed: '%#v'", err))
		return nil, err
	} else {
		userRefreshToken.RefreshToken = tokenFound
	}
	return &userRefreshToken, nil
}

func (dal *TokenDal) Create(token *models.UserRefreshToken) error {

	err := client.SetNX(getPrefixedDataStoreId(token.UserId), token.RefreshToken, 0).Err()
	if err != nil {
		log.WithFields(log.Fields{
			"userid": token.UserId,
		}).Error(fmt.Sprintf("creating token failed: '%#v'", err))
		return err
	}

	return nil
}

func (dal *TokenDal) Update(id int64, token *models.UserRefreshToken) error {

	err := client.SetXX(getPrefixedDataStoreId(id), token.RefreshToken, 0).Err()
	if err != nil {
		log.WithFields(log.Fields{
			"userid": id,
		}).Error(fmt.Sprintf("updating token failed: '%#v'", err))
		return err
	}
	return nil
}

func (dal *TokenDal) Delete(id int64) error {
	err := client.Del(getPrefixedDataStoreId(id)).Err()

	if err != nil {
		log.WithFields(log.Fields{
			"userid": id,
		}).Error(fmt.Sprintf("deleting token failed: '%#v'", err))
		return err
	}
	return nil
}

func (dal *TokenDal) Count(id int64) (int64, error) {
	return client.Exists(getPrefixedDataStoreId(id)).Result()
}

func getPrefixedDataStoreId(id int64) string {
	//return userKeyPrefix + keySeparator + id
	return fmt.Sprintf("%s%s%d", tokenKeyPrefix, keySeparator, id)
}
