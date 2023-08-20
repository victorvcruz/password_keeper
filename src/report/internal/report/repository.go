package report

import (
	"report.com/internal/platform/database"
)

type RepositoryClient interface {
	Create(report *Report) error
	Update(report *Report) error
	Delete(report *Report) error
	FindByUserId(userId string) ([]Report, error)
	FindByVaultId(vaultId string) ([]Report, error)
}

type repository struct {
	RepositoryClient
	db database.Client
}

func NewReportRepository(_db database.Client) RepositoryClient {
	return &repository{
		db: _db,
	}
}

func (r *repository) Create(report *Report) error {
	db, err := r.db.Begin()
	defer func() { r.db.CommitOrRollback(db, err) }()
	return db.Create(report).Error
}

func (r *repository) Update(report *Report) error {
	db, err := r.db.Begin()
	defer func() { r.db.CommitOrRollback(db, err) }()
	return db.Save(report).Error
}

func (r *repository) Delete(report *Report) error {
	db, err := r.db.Begin()
	defer func() { r.db.CommitOrRollback(db, err) }()
	return db.Delete(report).Error
}

func (r *repository) FindByUserId(userId string) ([]Report, error) {
	var reports []Report
	err := r.db.DB().Where("user_id = ? AND deleted_at is null", userId).Find(&reports).Error
	if err != nil {
		return nil, err
	}
	return reports, nil
}

func (r *repository) FindByVaultId(vaultId string) ([]Report, error) {
	var reports []Report
	err := r.db.DB().Where("vault_id = ? AND deleted_at is null", vaultId).Find(&reports).Error
	if err != nil {
		return nil, err
	}
	return reports, nil
}
