package service

import (
	"context"

	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
	"github.com/renatospaka/payment-transaction/core/usecase"
)

type AuthorizationService struct {
	usecases *usecase.TransactionUsecase
}

func NewAuthorizationService(usecases *usecase.TransactionUsecase) *AuthorizationService {
	return &AuthorizationService{
		usecases: usecases,
	}
}

func (a *AuthorizationService) Process(ctx context.Context, in *pb.AuthorizationRequest) (*pb.AuthorizationResponse, error) {
	auth := &pb.AuthorizationRequest{
		ClientId:      in.ClientId,
		TransactionId: in.TransactionId,
		Value:         in.Value,
	}

	response, err := a.usecases.Authorize(auth)
	authResponse := &pb.AuthorizationResponse{
		AuthorizationId: response.TransactionId,
		ClientId:        response.ClientId,
		TransactionId:   response.TransactionId,
		Status:          response.Status,
		Value:           response.Value,
	}

	if err != nil {
		authResponse.ErrorMessage = response.ErrorMessage
		return authResponse, nil
	}
	return authResponse, nil
}
