package service

import (
	"context"

	"github.com/venture-technology/vtx-account-manager/internal/repository"
	"github.com/venture-technology/vtx-account-manager/models"
)

type ResponsibleService struct {
	responsiblerepository repository.IResponsibleRepository
}

func NewResponsibleService(repo repository.IResponsibleRepository) *ResponsibleService {
	return &ResponsibleService{
		responsiblerepository: repo,
	}
}

func (rs *ResponsibleService) CreateContract(ctx context.Context, contract *models.Contract) error {
	return rs.responsiblerepository.CreateContract(ctx, contract)
}

func (rs *ResponsibleService) GetPartners(ctx context.Context, cpf *string) ([]models.Contract, error) {
	return rs.responsiblerepository.GetPartners(ctx, cpf)
}

func (rs *ResponsibleService) FindAllDriverAtSchool(ctx context.Context, cnpj *string) ([]models.Driver, error) {
	return rs.responsiblerepository.FindAllDriverAtSchool(ctx, cnpj)
}

func (rs *ResponsibleService) DeleteContract(ctx context.Context, record *int) error {
	return rs.responsiblerepository.DeleteContract(ctx, record)
}
