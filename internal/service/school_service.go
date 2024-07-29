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

func (ss *SchoolService) GetAllDriversToSchool(ctx context.Context, cnpj *string) ([]models.Handshake, error) {
	return ss.schoolrepository.GetAllDriversToSchool(ctx, cnpj)
}

func (ss *SchoolService) GetContracts(ctx context.Context, cnpj *string) ([]models.Contract, error) {
	return ss.schoolrepository.GetContracts(ctx, cnpj)
}

func (ss *SchoolService) DeletePartner(ctx context.Context, cnpj, cnh *string) error {
	return ss.schoolrepository.DeletePartner(ctx, cnpj, cnh)
}
