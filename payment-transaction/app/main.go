package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	// "github.com/renatospaka/payment-transaction/adapter/grpc/client"
	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
	// "github.com/renatospaka/payment-transaction/adapter/grpc/service"
	httpServer "github.com/renatospaka/payment-transaction/adapter/httpServer"
	repository "github.com/renatospaka/payment-transaction/adapter/postgres"
	"github.com/renatospaka/payment-transaction/adapter/web/controller"
	"github.com/renatospaka/payment-transaction/core/usecase"
	"github.com/renatospaka/payment-transaction/utils/configs"
)

func main() {
	log.Println("iniciando a aplicação")
	configs, err := configs.LoadConfig(".")
	if err != nil {
		log.Panic(err)
	}

	ctx := context.Background()

	//open connection to the database
	log.Println("iniciando conexão com o banco de dados")
	conn := "postgresql://" + configs.DBUser + ":" + configs.DBPassword + "@" + configs.DBHost + ":" + configs.DBPort + "/" + configs.DBName + "?sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	// client := grpcClient.NewGrpcClient(ctx)
	log.Printf("iniciando conexão com o servidor gRPC na porta :%s\n", configs.GRPCServerPort)
	options := make([]grpc.DialOption, 0)
	options = append(options, grpc.WithTransportCredentials(insecure.NewCredentials()))
	connGrpc, err := grpc.Dial("app1:" + configs.GRPCServerPort, options...)
	if err != nil {
		log.Panic(err)
	}
	defer connGrpc.Close()
	
	log.Println("iniciando gerador de transações")

	//web server
	repo := repository.NewPostgresDatabase(db)
	usecases := usecase.NewTransactionUsecase(repo)
	controllers := controller.NewTransactionController(usecases)
	webServer := httpServer.NewHttpServer(ctx, controllers)
	
	//grpc client
	// services := service.NewAuthorizationService(usecases)
	// cli := client.NewGrpcClient(ctx, connGrpc, services)

	// create a client to connect to the gRPC server an run ProcessTransaction
	log.Println("iniciando conexão com o servidor gRPC")
	client := pb.NewAuthorizationServiceClient(connGrpc)

	
	request := &pb.AuthorizationRequest{
		ClientId: " ce84577c-2011-4750-9c73-838669b4e02c",
		TransactionId: "9bb8cea2-fe27-4eed-95fd-1b8c5033d26c",
		Value: 100,
	}

	log.Println("enviando requisição para o servidor gRPC")
	response, err := client.Process(ctx, request)
	if err != nil {
		log.Println("erro ao enviar requisição para o servidor gRPC")
		log.Panic(err)
	}

	log.Println(response)
	
	
	//start web server
	log.Printf("gerador de transações escutando porta: %s\n", configs.WEBServerPort)
	http.ListenAndServe(":"+configs.WEBServerPort, webServer.Server)
}
