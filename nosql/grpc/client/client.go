package client

import (
	"nosql/config"
	"google.golang.org/grpc"
)

type ServiceManagerI interface {}

type grpcClients struct {
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	return &grpcClients{}, nil
}
