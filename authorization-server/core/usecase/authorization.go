package usecase

import (
	"log"

	"github.com/renatospaka/authorization-server/core/dto"
	"github.com/renatospaka/authorization-server/core/entity"
	"github.com/renatospaka/authorization-server/core/repository"
	"github.com/renatospaka/authorization-server/utils/dateTime"
)

type AuthorizationUsecase struct {
	repo repository.AuthorizationInterface
}

func NewAuthorizationUsecase(repo repository.AuthorizationInterface) *AuthorizationUsecase {
	return &AuthorizationUsecase{
		repo: repo,
	}
}

func (a *AuthorizationUsecase) ProcessAuthorization(auth *dto.AuthorizationProcessDto) (*dto.AuthorizationProcessResultDto, error) {
	log.Println("usecase.authorizations.process")

	response := &dto.AuthorizationProcessResultDto{
		ClientID: auth.ClientID,
		Value: auth.Value,
	}

	authorization, err := entity.NewAuthorization(auth.ClientID, auth.Value)
	if err != nil {
		response.ErrorMessage = err.Error()
		return response, err
	}
	
	authorization.Process()
	response.ID = authorization.GetID()
	response.Value = authorization.GetValue()
	response.Status = authorization.GetStatus()
	response.DeniedAt = dateTime.FormatDateToNull(authorization.DeniedAt())
	response.ApprovedAt = dateTime.FormatDateToNull(authorization.ApprovedAt())
	if !authorization.IsValid() {
		response.ErrorMessage = err.Error()
		return response, err
	}

	err = a.repo.Process(authorization)
	if err != nil {
		response.ErrorMessage = err.Error()
		return response, err
	}
	return response, nil
}
