package grpc_service_clients

import (
	"google.golang.org/grpc"

	"github.com/perfectogo/template-service-with-kafka/internal/pkg/config"
)

type ServiceClients interface {
	Close()
}

type serviceClients struct {
	services []*grpc.ClientConn
}

func New(config *config.Config) (ServiceClients, error) {
	return &serviceClients{
		services: []*grpc.ClientConn{},
	}, nil
}

func (s *serviceClients) Close() {
	// closing investment service
	for _, conn := range s.services {
		conn.Close()
	}
}
