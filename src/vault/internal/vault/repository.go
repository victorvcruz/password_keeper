package vault

import "vault.com/internal/platform/database"

type RepositoryClient interface {
	Create(vault *Vault) (*Vault, error)
	Update(vault *Vault) (*Vault, error)
	Delete(id uint) error
	FindByID(id uint) (*Vault, error)
	FindAllByUserID(userID uint) ([]Vault, error)
}

type repository struct {
	db database.DatabaseClient
}

func NewVaultRepository(_db database.DatabaseClient) RepositoryClient {
	return &repository{
		db: _db,
	}
}

func (r *repository) Create(vault *Vault) (*Vault, error) {
	if err := r.db.DB().Create(vault).Error; err != nil {
		return nil, err
	}
	return vault, nil
}

func (r *repository) Update(vault *Vault) (*Vault, error) {
	if err := r.db.DB().Save(vault).Error; err != nil {
		return nil, err
	}
	return vault, nil
}

func (r *repository) Delete(id uint) error {
	return r.db.DB().Where("id = ?", id).Delete(&Vault{}).Error
}

func (r *repository) FindByID(id uint) (*Vault, error) {
	var vault Vault
	if err := r.db.DB().First(&vault, id).Error; err != nil {
		return nil, err
	}
	return &vault, nil
}

func (r *repository) FindAllByUserID(userID uint) ([]Vault, error) {
	var vaults []Vault
	if err := r.db.DB().Where("user_id = ?", userID).Find(&vaults).Error; err != nil {
		return nil, err
	}
	return vaults, nil
}
