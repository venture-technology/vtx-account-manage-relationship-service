package service

import (
	"context"

	"github.com/venture-technology/vtx-account-manager/internal/repository"
	"github.com/venture-technology/vtx-account-manager/models"
)

type SchoolService struct {
	schoolrepository repository.ISchoolRepository
}

func NewSchoolService(repo repository.ISchoolRepository) *SchoolService {
	return &SchoolService{
		schoolrepository: repo,
	}
}

func (ss *SchoolService) GetDriver(ctx context.Context, cnpj *string) ([]models.Handshake, error) {
	return ss.schoolrepository.GetDriver(ctx, cnpj)
}

func (ss *SchoolService) GetSponsors(ctx context.Context, cnh *string) ([]models.Sponsor, error) {
	return ss.schoolrepository.GetSponsors(ctx, cnh)
}
