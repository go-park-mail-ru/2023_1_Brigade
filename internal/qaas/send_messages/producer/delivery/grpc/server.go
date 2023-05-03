package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"net"
	"project/internal/generated"
	"project/internal/qaas/send_messages/producer"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (p *producerServiceGRPCServer) ProduceMessage(ctx context.Context, bytes *generated.Bytes) (*emptypb.Empty, error) {
	err := p.producerUsecase.ProduceMessage(ctx, model_conversion.FromProtoBytesToBytes(bytes))
	return new(empty.Empty), err
}
