package service

import (
	"nosql/config"
	"nosql/grpc/client"
	"nosql/pkg/logger"
	"nosql/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.Logger, strg storage.StorageI, svcs client.ServiceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()
	reflection.Register(grpcServer)
	return
}
