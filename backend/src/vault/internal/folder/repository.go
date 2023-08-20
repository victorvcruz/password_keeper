package folder

import (
	"vault.com/internal/platform/database"
)

type RepositoryClient interface {
	Create(folder *Folder) (*Folder, error)
	Update(folder *Folder) (*Folder, error)
	Delete(folderId uint64) error
	FindByID(id uint64) (*Folder, error)
	FindAllByUserId(userID uint64) ([]Folder, error)
}

type repository struct {
	db database.Client
}

func NewFolderRepository(_db database.Client) RepositoryClient {
	return &repository{
		db: _db,
	}
}

func (r *repository) Create(folder *Folder) (*Folder, error) {
	db, err := r.db.Begin()
	defer func() { r.db.CommitOrRollback(db, err) }()

	if err = db.Create(folder).Error; err != nil {
		return nil, err
	}
	return folder, nil
}

func (r *repository) Update(folder *Folder) (*Folder, error) {
	db, err := r.db.Begin()
	defer func() { r.db.CommitOrRollback(db, err) }()

	if err = db.Save(folder).Error; err != nil {
		return nil, err
	}
	return folder, nil
}

func (r *repository) Delete(folderId uint64) error {
	db, err := r.db.Begin()
	defer func() { r.db.CommitOrRollback(db, err) }()

	return db.Where("id = ?", folderId).Delete(&Folder{}).Error
}

func (r *repository) FindByID(id uint64) (*Folder, error) {
	var folder Folder
	if err := r.db.DB().First(&folder, id).Error; err != nil {
		return nil, err
	}
	return &folder, nil
}

func (r *repository) FindAllByUserId(userID uint64) ([]Folder, error) {
	var folders []Folder
	if err := r.db.DB().Where("user_id = ?", userID).Find(&folders).Error; err != nil {
		return nil, err
	}
	return folders, nil
}
