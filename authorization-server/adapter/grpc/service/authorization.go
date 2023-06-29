package service

import (
	"context"

	"github.com/renatospaka/authorization-server/adapter/grpc/pb"
	"github.com/renatospaka/authorization-server/core/dto"
	"github.com/renatospaka/authorization-server/core/usecase"
)

// handle it as it were a controller
type AuthorizationService struct {
	usecases *usecase.AuthorizationUsecase
	pb.UnimplementedAuthorizationServiceServer
}

func NewAuthorizationService(usecases *usecase.AuthorizationUsecase) *AuthorizationService {
	return &AuthorizationService{
		usecases: usecases,
	}
}

// Process the authorization request and return to the gRPC caller
func (a *AuthorizationService) Process(ctx context.Context, in *pb.AuthorizationRequest) (*pb.AuthorizationResponse, error) {
	auth := &dto.AuthorizationProcessDto{
		Value:         in.Value,
		ClientID:      in.ClientId,
		TransactionID: in.TransactionId,
	}
	authResponse := &pb.AuthorizationResponse{}

	response, err := a.usecases.ProcessAuthorization(auth)
	authResponse = &pb.AuthorizationResponse{
		AuthorizationId: response.ID,
		ClientId:        response.ClientID,
		TransactionId:   response.TransactionID,
		Status:          response.Status,
		Value:           response.Value,
	}

	if err != nil {
		authResponse.ErrorMessage = response.ErrorMessage
		return authResponse, nil
	}
	return authResponse, nil
}
