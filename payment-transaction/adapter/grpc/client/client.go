package client

import (
	"context"
	"log"

	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
	"github.com/renatospaka/payment-transaction/adapter/grpc/service"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	ctx      context.Context
	conn     *grpc.ClientConn
	services *service.AuthorizationService
	client   pb.AuthorizationServiceClient
}

func NewGrpcClient(ctx context.Context, conn *grpc.ClientConn, services *service.AuthorizationService) *GrpcClient {
	log.Println("estabelecendo conexão com o servidor gRPC")
	cli := &GrpcClient{
		ctx:      ctx,
		conn:     conn,
		services: services,
	}

	cli.client = pb.NewAuthorizationServiceClient(cli.conn)
	return cli
}
