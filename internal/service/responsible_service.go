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

	// validar distancia

	// calcular valor de contrato

	// criar produto

	// criar preço

	// crio inscrição

	return rs.responsiblerepository.CreateContract(ctx, contract)
}

func (rs *ResponsibleService) FindContractsByCpf(ctx context.Context, cpf, status *string) ([]models.Contract, error) {
	return rs.responsiblerepository.FindContractsByCpf(ctx, cpf, status)
}

func (rs *ResponsibleService) GetContract() {}

func (rs *ResponsibleService) UpdateContract() {}

func (rs *ResponsibleService) FindAllDriverAtSchool(ctx context.Context, cnpj *string) ([]models.Driver, error) {
	return rs.responsiblerepository.FindAllDriverAtSchool(ctx, cnpj)
}
