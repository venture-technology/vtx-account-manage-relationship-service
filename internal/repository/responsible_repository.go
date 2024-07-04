package repository

import (
	"context"
	"database/sql"

	"github.com/venture-technology/vtx-account-manager/models"
)

type IResponsibleRepository interface {
	FindAllDriverAtSchool(ctx context.Context, cnpj *string) ([]models.Driver, error)
	CreateSponsor(ctx context.Context, sponsor *models.Sponsor) error
	GetPartners(ctx context.Context, cpf *string) ([]models.Sponsor, error)
	BreachSponsor(ctx context.Context, record *int) error
}

type ResponsibleRepository struct {
	db *sql.DB
}

func NewResponsibleRepository(db *sql.DB) *ResponsibleRepository {
	return &ResponsibleRepository{
		db: db,
	}
}

func (rr *ResponsibleRepository) CreateSponsor(ctx context.Context, sponsor *models.Sponsor) error {

	return nil

}

func (rr *ResponsibleRepository) GetPartners(ctx context.Context, cpf *string) ([]models.Sponsor, error) {

	return nil, nil

}

func (rr *ResponsibleRepository) FindAllDriverAtSchool(ctx context.Context, cnpj *string) ([]models.Driver, error) {

	return nil, nil

}

func (rr *ResponsibleRepository) BreachSponsor(ctx context.Context, record *int) error {

	return nil

}
