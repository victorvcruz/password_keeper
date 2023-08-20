package report

import (
	"report.com/internal/platform/database"
)

type ReportRepositoryClient interface {
	Create(report *Report) error
	Update(report *Report) error
	Delete(report *Report) error
	FindByUserId(userId string) ([]Report, error)
	FindByVaultId(vaultId string) ([]Report, error)
}

type reportRepository struct {
	ReportRepositoryClient
	db database.Client
}

func NewReportRepository(_db database.Client) ReportRepositoryClient {
	return &reportRepository{
		db: _db,
	}
}

func (r *reportRepository) Create(report *Report) error {
	db, err := r.db.Begin()
	defer func() { r.db.CommitOrRollback(db, err) }()
	return db.Create(report).Error
}

func (r *reportRepository) Update(report *Report) error {
	db, err := r.db.Begin()
	defer func() { r.db.CommitOrRollback(db, err) }()
	return db.Save(report).Error
}

func (r *reportRepository) Delete(report *Report) error {
	db, err := r.db.Begin()
	defer func() { r.db.CommitOrRollback(db, err) }()
	return db.Delete(report).Error
}

func (r *reportRepository) FindByUserId(userId string) ([]Report, error) {
	var reports []Report
	err := r.db.DB().Where("user_id = ? AND deleted_at is null", userId).Find(&reports).Error
	if err != nil {
		return nil, err
	}
	return reports, nil
}

func (r *reportRepository) FindByVaultId(vaultId string) ([]Report, error) {
	var reports []Report
	err := r.db.DB().Where("vault_id = ? AND deleted_at is null", vaultId).Find(&reports).Error
	if err != nil {
		return nil, err
	}
	return reports, nil
}
