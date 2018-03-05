package grpcserver


import (
	"fmt"

	log "github.com/sirupsen/logrus"
	pb "github.com/vahdet/go-refresh-token-store-redis/proto"
	"golang.org/x/net/context"

	"github.com/vahdet/go-refresh-token-store-redis/services/interfaces"
	"github.com/vahdet/go-refresh-token-store-redis/app/utils"
)

type UserServer struct {
	Service interfaces.UserService
}

func (s *UserServer) Get(ctx context.Context, in *pb.UserId) (*pb.UserToken, error) {
	res, err := s.Service.Get(in.Value)
	if err != nil {
		log.WithFields(log.Fields{
			"id": in.Value,
		}).Error(fmt.Sprintf("Getting failed: '%#v'", err))
		return nil, err
	}

	return utils.ConvertModelToProto(res)
}

func (s *UserServer) Create(ctx context.Context, in *pb.UserToken) (*pb.UserId, error) {
	token, err := utils.ConvertProtoToModel(in)
	if err != nil {
		log.WithFields(log.Fields{
			"userid": in.UserId,
			"token": in.Token,
		}).Error(fmt.Sprintf("Conversion failed: '%#v'", err))
		return nil, err
	}

	res, err := s.Service.Create(token)
	if err != nil {
		log.WithFields(log.Fields{
			"userid": in.UserId,
			"token": in.Token,
		}).Error(fmt.Sprintf("Creation failed: '%#v'", err))
		return nil, err
	}

	return &pb.UserId{
		Value: res.UserId,
	}, nil
}

func (s *UserServer) Update(ctx context.Context, in *pb.UserToken) (*pb.UserId, error) {
	token, err := utils.ConvertProtoToModel(in)
	if err != nil {
		log.WithFields(log.Fields{
			"userid": in.UserId,
			"token": in.Token,
		}).Error(fmt.Sprintf("Conversion failed: '%#v'", err))
		return nil, err
	}

	res, err := s.Service.Update(token.UserId, token)
	if err != nil {
		log.WithFields(log.Fields{
			"userid": in.UserId,
		}).Error(fmt.Sprintf("Update failed: '%#v'", err))
		return nil, err
	}

	return &pb.UserId{
		Value: res.UserId,
	}, nil
}

func (s *UserServer) Delete(ctx context.Context, in *pb.UserId) (*pb.UserId, error) {
	res, err := s.Service.Delete(in.Value)
	if err != nil {
		log.WithFields(log.Fields{
			"id": in.Value,
		}).Error(fmt.Sprintf("Deletion failed: '%#v'", err))
		return nil, err
	}

	return &pb.UserId{
		Value: res.UserId,
	}, nil
}
