package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"net"
	"project/internal/generated"
	producer "project/internal/microservices/producer/usecase"
	"project/internal/pkg/model_conversion"
)

type producerServiceGRPCServer struct {
	grpcServer      *grpc.Server
	producerUsecase producer.Usecase
}

func NewProducerServiceGRPCServer(grpcServer *grpc.Server, producerUsecase producer.Usecase) *producerServiceGRPCServer {
	return &producerServiceGRPCServer{
		grpcServer:      grpcServer,
		producerUsecase: producerUsecase,
	}
}

func (p *producerServiceGRPCServer) StartGRPCServer(listenURL string) error {
	lis, err := net.Listen("tcp", listenURL)
	if err != nil {
		return err
	}

	generated.RegisterProducerServer(p.grpcServer, p)

	return p.grpcServer.Serve(lis)
}

func (p *producerServiceGRPCServer) ProduceMessage(ctx context.Context, message *generated.ProducerMessage) (*empty.Empty, error) {
	err := p.producerUsecase.ProduceMessage(ctx, model_conversion.FromProtoProducerMessageToProducerMessage(message))
	return new(empty.Empty), err
}
