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

func (ds *DriverService) GetSchool(ctx context.Context, cnh *string) ([]models.School, error) {
	return ds.driverrepository.GetSchool(ctx, cnh)
}

func (ds *DriverService) GetSponsors(ctx context.Context, cnpj *string) ([]models.Sponsor, error) {
	return ds.driverrepository.GetSponsors(ctx, cnpj)
}
