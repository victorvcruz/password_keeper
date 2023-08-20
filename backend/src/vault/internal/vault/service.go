package vault

import (
	"vault.com/internal/platform/database"
	"vault.com/internal/utils/errors"
)

type ServiceClient interface {
	Find(id uint64) (Response, error)
	FindAll(userId uint64) ([]Response, error)
	Create(req Request) (Response, error)
	Update(req Request) (Response, error)
	Delete(id uint64) error
	FindAllByFolder(folderId uint64) ([]Response, error)
}

type service struct {
	db         database.Client
	repository RepositoryClient
}

func (s service) Find(id uint64) (Response, error) {
	vault, err := s.repository.FindByID(id)
	if err != nil {
		return Response{}, err
	}

	if vault == nil {
		return Response{}, &errors.NotFound{Msg: "not found vault"}
	}
	return vault.ToResponse(), nil
}

func (s service) FindAll(userId uint64) ([]Response, error) {
	vaults, err := s.repository.FindAllByUserID(userId)
	if err != nil {
		return nil, err
	}

	if vaults == nil {
		return nil, &errors.NotFound{Msg: "not found vaults"}
	}

	response := make([]Response, len(vaults))
	for i := range vaults {
		response[i] = vaults[i].ToResponse()
	}
	return response, nil
}

func (s service) Create(req Request) (Response, error) {
	vault, err := s.repository.Create(req.ToModel())
	if err != nil {
		return Response{}, err
	}
	return vault.ToResponse(), nil
}

func (s service) Update(req Request) (Response, error) {
	vault, err := s.repository.Update(req.ToModel())
	if err != nil {
		return Response{}, err
	}
	return vault.ToResponse(), nil
}

func (s service) Delete(id uint64) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s service) FindAllByFolder(folderId uint64) ([]Response, error) {
	vaults, err := s.repository.FindAllByFolderID(folderId)
	if err != nil {
		return nil, err
	}

	if vaults == nil {
		return nil, &errors.NotFound{Msg: "not found vaults"}
	}

	response := make([]Response, len(vaults))
	for i := range vaults {
		response[i] = vaults[i].ToResponse()
	}
	return response, nil
}

func NewVaultService(
	_db database.Client,
	_repository RepositoryClient,
) ServiceClient {
	return &service{
		db:         _db,
		repository: _repository,
	}
}
