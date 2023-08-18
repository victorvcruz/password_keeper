package vault

import "vault.com/internal/platform/database"

type RepositoryClient interface{}

type repository struct {
	db database.DatabaseClient
}

func NewVaultRepository(_db database.DatabaseClient) RepositoryClient {
	return &repository{
		db: _db,
	}
}
