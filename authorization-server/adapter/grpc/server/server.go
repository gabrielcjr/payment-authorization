package server

import (
	"context"
	"log"
	"net"

	"github.com/renatospaka/authorization-server/adapter/grpc/pb"
	"github.com/renatospaka/authorization-server/adapter/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcServer struct {
	ctx      context.Context
	services *service.AuthorizationService
	Server   *grpc.Server
}

func NewGrpcServer(ctx context.Context, services *service.AuthorizationService) *GrpcServer {
	log.Println("iniciando conexão com o servidor gRPC")
	srv := &GrpcServer{
		ctx: ctx,
		services: services,
	}
	return srv
}

func (g *GrpcServer) Connect(port string) {
	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Panic(err)
	}
	defer lis.Close()

	g.Server = grpc.NewServer()
	pb.RegisterAuthorizationServiceServer(g.Server, g.services)
	reflection.Register(g.Server)		// to allow Evans to test the server
	if err := g.Server.Serve(lis); err != nil {
		log.Panic(err)
	}
}
