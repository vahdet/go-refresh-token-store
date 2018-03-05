package main

import (
	"fmt"
	"github.com/vahdet/go-refresh-token-store-redis/dal"
	"github.com/vahdet/go-refresh-token-store-redis/services"
	"net"
	pb "github.com/vahdet/go-refresh-token-store-redis/proto"
	"google.golang.org/grpc"
	"github.com/vahdet/go-refresh-token-store-redis/grpcserver"
	"google.golang.org/grpc/reflection"
	log "github.com/sirupsen/logrus"
)

const (
	defaultGrpcPort = "5302"
)

func main() {

	// Initialize DataStore
	dal.InitDataStoreClient()
	defer dal.CloseDataStoreClient()

	// Serve GRPC
	lis, err := net.Listen("tcp", defaultGrpcPort)
	if err != nil {
		panic(fmt.Errorf("failed to listen: %v", err))
	}

	s := grpc.NewServer()
	pb.RegisterTokenServiceServer(s, &grpcserver.UserServer{Service: services.NewTokenService(dal.NewTokenDal())})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
