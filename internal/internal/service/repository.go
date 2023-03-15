package service

import "gorm.io/gorm"

type ServiceRepositoryClient interface {
	Create(internal *Internal) error
	Update(internal *Internal) error
	FindByService(service string) (*Internal, error)
}

type serviceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(_db *gorm.DB) ServiceRepositoryClient {
	return &serviceRepository{
		db: _db,
	}
}

func (u *serviceRepository) Create(internal *Internal) error {
	err := u.db.Create(internal).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *serviceRepository) Update(internal *Internal) error {
	err := u.db.Save(internal).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *serviceRepository) FindByService(service string) (*Internal, error) {
	var internalModel Internal
	err := u.db.Where("service = ? AND deleted_at is null", service).Find(&internalModel).Error
	if err != nil {
		return nil, err
	}

	if internalModel.Service == "" {
		return nil, nil
	}

	return &internalModel, nil
}
