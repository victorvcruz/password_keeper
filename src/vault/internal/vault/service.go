package vault

import "vault.com/internal/platform/database"

type ServiceClient interface {
	Find() (Response, error)
	FindAll() ([]Response, error)
	Create(auth *Request) (Response, error)
	Update(auth *Request) (Response, error)
	Delete() (Response, error)
}

type service struct {
	db         database.DatabaseClient
	repository RepositoryClient
}

func (s service) Find() (Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) FindAll() ([]Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Create(auth *Request) (Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Update(auth *Request) (Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Delete() (Response, error) {
	//TODO implement me
	panic("implement me")
}

func NewVaultService(
	_db database.DatabaseClient,
	_repository RepositoryClient,
) ServiceClient {
	return &service{
		db:         _db,
		repository: _repository,
	}
}
