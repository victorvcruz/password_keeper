package folder

import (
	"vault.com/internal/platform/database"
)

type ServiceClient interface {
	Create(req Request) (Response, error)
	FindAllByUserId(userId uint64) ([]Response, error)
	Update(req Request) (Response, error)
	Delete(folderId uint64) error
	Find(folderId uint64) (Response, error)
}

type service struct {
	db         database.Client
	repository RepositoryClient
}

func NewFolderService(
	_db database.Client,
	_repository RepositoryClient,
) ServiceClient {
	return &service{
		db:         _db,
		repository: _repository,
	}
}

func (s service) Create(req Request) (Response, error) {
	folder, err := s.repository.Create(req.ToModel())
	if err != nil {
		return Response{}, err
	}
	return folder.ToResponse(), nil
}

func (s service) Update(req Request) (Response, error) {
	folder, err := s.repository.Update(req.ToModel())
	if err != nil {
		return Response{}, err
	}
	return folder.ToResponse(), nil
}

func (s service) Delete(folderId uint64) error {
	err := s.repository.Delete(folderId)
	if err != nil {
		return err
	}
	return nil
}

func (s service) Find(folderId uint64) (Response, error) {
	folders, err := s.repository.FindByID(folderId)
	if err != nil {
		return Response{}, err
	}
	return folders.ToResponse(), nil
}

func (s service) FindAllByUserId(userId uint64) ([]Response, error) {
	folders, err := s.repository.FindAllByUserId(userId)
	if err != nil {
		return nil, err
	}

	response := make([]Response, len(folders))
	for i := range folders {
		response[i] = folders[i].ToResponse()
	}

	return response, nil
}
