package folder

import "vault.com/internal/platform/database"

type RepositoryClient interface {
	Create(folder *Folder) (*Folder, error)
	Update(folder *Folder) (*Folder, error)
	Delete(folderId uint) error
	FindByID(id uint) (*Folder, error)
	FindAllByUserId(userID uint) ([]Folder, error)
}

type repository struct {
	db database.DatabaseClient
}

func NewFolderRepository(_db database.DatabaseClient) RepositoryClient {
	return &repository{
		db: _db,
	}
}

func (r *repository) Create(folder *Folder) (*Folder, error) {
	if err := r.db.DB().Create(folder).Error; err != nil {
		return nil, err
	}
	return folder, nil
}

func (r *repository) Update(folder *Folder) (*Folder, error) {
	if err := r.db.DB().Save(folder).Error; err != nil {
		return nil, err
	}
	return folder, nil
}

func (r *repository) Delete(folderId uint) error {
	return r.db.DB().Where("id = ?", folderId).Delete(&Folder{}).Error
}

func (r *repository) FindByID(id uint) (*Folder, error) {
	var folder Folder
	if err := r.db.DB().First(&folder, id).Error; err != nil {
		return nil, err
	}
	return &folder, nil
}

func (r *repository) FindAllByUserId(userID uint) ([]Folder, error) {
	var folders []Folder
	if err := r.db.DB().Where("user_id = ?", userID).Find(&folders).Error; err != nil {
		return nil, err
	}
	return folders, nil
}
