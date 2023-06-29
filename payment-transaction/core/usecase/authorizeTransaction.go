package usecase

import (
	"context"
	"log"
	"math/rand"

	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
	pkgEntity "github.com/renatospaka/payment-transaction/utils/entity"
)

func (t *TransactionUsecase) authorizeTransaction(ctx context.Context, in *pb.AuthorizationRequest) (*pb.AuthorizationResponse, error) {
	log.Println("usecase.transactions.authorize")

	// Issue #21 open to solve this lack of communication to the server.
	// Meanwhile, this is an workaround to respond to the caller
	min, max, status := 0, 100, "pending"

	random := rand.Intn(max-min) + min
	if random <= 30 {
		status = "denied"
	} else {
		status = "approved"
	}

	uuid := pkgEntity.NewID()
	response := &pb.AuthorizationResponse{
		AuthorizationId: uuid.String(),
		ClientId:        in.ClientId,
		TransactionId:   in.TransactionId,
		Status:          status,
		Value:           in.Value,
	}
	return response, nil
}
