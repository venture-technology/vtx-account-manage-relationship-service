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

func (ds *DriverService) GetSchool(ctx context.Context, cnh, cnpj *string) (*models.Handshake, error) {
	return ds.driverrepository.GetSchool(ctx, cnh, cnpj)
}

func (ds *DriverService) GetSponsors(ctx context.Context, cnh *string) ([]models.Sponsor, error) {
	return ds.driverrepository.GetSponsors(ctx, cnh)
}

func (ds *DriverService) CreatePartner(ctx context.Context, handshake *models.Handshake) error {
	return ds.driverrepository.CreatePartner(ctx, handshake)
}

func (ds *DriverService) GetPartners(ctx context.Context, cnh *string) ([]models.Handshake, error) {
	return ds.driverrepository.GetPartners(ctx, cnh)
}

func (ds *DriverService) GetSponsorsByShift(ctx context.Context, cnh, shift *string) ([]models.Sponsor, error) {
	return ds.driverrepository.GetSponsorsByShift(ctx, cnh, shift)
}
