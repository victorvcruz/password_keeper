package report

import "gorm.io/gorm"

type ReportRepositoryClient interface {
	Create(report *Report) error
	Update(report *Report) error
	Delete(report *Report) error
	FindByUserId(userId string) ([]Report, error)
	FindByVaultId(vaultId string) ([]Report, error)
}

type reportRepository struct {
	ReportRepositoryClient
	db *gorm.DB
}

func NewReportRepository(_db *gorm.DB) ReportRepositoryClient {
	return &reportRepository{
		db: _db,
	}
}

func (r *reportRepository) Create(report *Report) error {
	err := r.db.Create(report).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *reportRepository) Update(report *Report) error {
	err := r.db.Save(report).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *reportRepository) Delete(report *Report) error {
	err := r.db.Delete(report).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *reportRepository) FindByUserId(userId string) ([]Report, error) {
	var reports []Report

	err := r.db.Where("user_id = ? AND deleted_at is null", userId).Find(&reports).Error
	if err != nil {
		return nil, err
	}

	return reports, nil
}

func (r *reportRepository) FindByVaultId(vaultId string) ([]Report, error) {
	var reports []Report

	err := r.db.Where("vault_id = ? AND deleted_at is null", vaultId).Find(&reports).Error
	if err != nil {
		return nil, err
	}

	return reports, nil
}
