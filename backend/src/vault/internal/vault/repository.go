package vault

import "vault.com/internal/platform/database"

type RepositoryClient interface {
	Create(vault *Vault) (*Vault, error)
	Update(vault *Vault) (*Vault, error)
	Delete(id uint64) error
	FindByID(id uint64) (*Vault, error)
	FindAllByUserID(userID uint64) ([]Vault, error)
	FindAllByFolderID(folderID uint64) ([]Vault, error)
}

type repository struct {
	db database.Client
}

func NewVaultRepository(_db database.Client) RepositoryClient {
	return &repository{
		db: _db,
	}
}

func (r *repository) Create(vault *Vault) (*Vault, error) {
	db, err := r.db.Begin()
	defer func() { r.db.CommitOrRollback(db, err) }()

	if err = db.Create(vault).Error; err != nil {
		return nil, err
	}
	return vault, nil
}

func (r *repository) Update(vault *Vault) (*Vault, error) {
	db, err := r.db.Begin()
	defer func() { r.db.CommitOrRollback(db, err) }()

	if err = db.Save(vault).Error; err != nil {
		return nil, err
	}
	return vault, nil
}

func (r *repository) Delete(id uint64) error {
	db, err := r.db.Begin()
	defer func() { r.db.CommitOrRollback(db, err) }()

	return db.Where("id = ?", id).Delete(&Vault{}).Error
}

func (r *repository) FindByID(id uint64) (*Vault, error) {
	var vault Vault
	if err := r.db.DB().First(&vault, id).Error; err != nil {
		return nil, err
	}
	return &vault, nil
}

func (r *repository) FindAllByUserID(userID uint64) ([]Vault, error) {
	var vaults []Vault
	if err := r.db.DB().Where("user_id = ?", userID).Find(&vaults).Error; err != nil {
		return nil, err
	}
	return vaults, nil
}

func (r *repository) FindAllByFolderID(folderID uint64) ([]Vault, error) {
	var vaults []Vault
	if err := r.db.DB().Where("folder_id = ?", folderID).Find(&vaults).Error; err != nil {
		return nil, err
	}
	return vaults, nil
}
