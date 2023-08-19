package folder

import "vault.com/internal/platform/database"

type RepositoryClient interface {
	Create(folder *Folder) error
	Update(folder *Folder) error
	Delete(folder *Folder) error
	FindByID(id uint) (*Folder, error)
}

type repository struct {
	db database.DatabaseClient
}

func NewFolderRepository(_db database.DatabaseClient) RepositoryClient {
	return &repository{
		db: _db,
	}
}

func (r *repository) Create(folder *Folder) error {
	return r.db.DB().Create(folder).Error
}

func (r *repository) Update(folder *Folder) error {
	return r.db.DB().Save(folder).Error
}

func (r *repository) Delete(folder *Folder) error {
	return r.db.DB().Delete(folder).Error
}

func (r *repository) FindByID(id uint) (*Folder, error) {
	var folder Folder
	if err := r.db.DB().First(&folder, id).Error; err != nil {
		return nil, err
	}
	return &folder, nil
}
