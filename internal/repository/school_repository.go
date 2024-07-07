package repository

import (
	"context"
	"database/sql"

	"github.com/venture-technology/vtx-account-manager/models"
)

type ISchoolRepository interface {
	GetDriver(ctx context.Context, cnpj *string) ([]models.Handshake, error)
	GetSponsors(ctx context.Context, cnh *string) ([]models.Sponsor, error)
	DeletePartner(ctx context.Context, cnh *string) error
}

type SchoolRepository struct {
	db *sql.DB
}

func NewSchoolRepository(db *sql.DB) *SchoolRepository {
	return &SchoolRepository{
		db: db,
	}
}

func (sr *SchoolRepository) GetDriver(ctx context.Context, cnpj *string) ([]models.Handshake, error) {

	return nil, nil

}

func (sr *SchoolRepository) GetSponsors(ctx context.Context, cnh *string) ([]models.Sponsor, error) {

	return nil, nil

}

func (sr *SchoolRepository) DeletePartner(ctx context.Context, cnh *string) error {

	return nil

}
