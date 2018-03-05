package utils

import (
	"github.com/vahdet/go-refresh-token-store-redis/models"
	pb "github.com/vahdet/go-refresh-token-store-redis/proto"
)

func ConvertModelToProto(token *models.UserRefreshToken) (*pb.UserToken, error) {

	return &pb.UserToken{
		UserId: 	token.UserId,
		Token:		token.RefreshToken,
	}, nil
}

func ConvertProtoToModel(proto *pb.UserToken) (*models.UserRefreshToken, error) {

	return &models.UserRefreshToken{
		UserId:			proto.UserId,
		RefreshToken:  	proto.Token,
	}, nil
}

