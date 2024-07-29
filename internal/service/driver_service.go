package service

import (
	"context"

	"github.com/venture-technology/vtx-account-manager/internal/repository"
	"github.com/venture-technology/vtx-account-manager/models"
)

type DriverService struct {
	driverrepository repository.IDriverRepository
}

func NewDriverService(repo repository.IDriverRepository) *DriverService {
	return &DriverService{
		driverrepository: repo,
	}
}

func (ds *DriverService) GetSchool(ctx context.Context, cnh, cnpj *string) (*models.Partner, error) {
	return ds.driverrepository.GetSchool(ctx, cnh, cnpj)
}

func (ds *DriverService) GetContracts(ctx context.Context, cnh *string) ([]models.Contract, error) {
	return ds.driverrepository.GetContracts(ctx, cnh)
}

func (ds *DriverService) CreatePartner(ctx context.Context, partner *models.Partner) error {
	return ds.driverrepository.CreatePartner(ctx, partner)
}

func (ds *DriverService) GetPartners(ctx context.Context, cnh *string) ([]models.Partner, error) {
	return ds.driverrepository.GetPartners(ctx, cnh)
}

func (ds *DriverService) GetContractsByShift(ctx context.Context, cnh, shift *string) ([]models.Contract, error) {
	return ds.driverrepository.GetContractsByShift(ctx, cnh, shift)
}
