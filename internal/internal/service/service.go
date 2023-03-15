package service

import "os"

type InternalServiceClient interface {
	TokenRequest(login TokenRequest) (string, error)
	InsertDefaultService() error
}

type internalService struct {
	repository ServiceRepositoryClient
}

func NewInternalService(
	_repository ServiceRepositoryClient,
) InternalServiceClient {
	return &internalService{
		repository: _repository,
	}
}

func (i *internalService) InsertDefaultService() error {
	service := os.Getenv("DEFAULT_SERVICE")

	internal, err := i.repository.FindByService(service)
	if err != nil {
		return err
	}
	if internal != nil {
		return nil
	}

	var internalModel *Internal
	internalModel.FillFields(service, os.Getenv("DEFAULT_TOKEN"), os.Getenv("DEFAULT_PASSWORD"))
	err = i.repository.Create(internalModel)
	if err != nil {
		return err
	}
	return nil
}

func (i *internalService) TokenRequest(request TokenRequest) (string, error) {
	service, err := i.repository.FindByService(request.Service)
	if err != nil {
		return "", err
	}

	return service.Token, nil
}
