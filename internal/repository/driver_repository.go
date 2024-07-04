package repository

import (
	"context"
	"database/sql"

	"github.com/venture-technology/vtx-account-manager/models"
)

type IDriverRepository interface {
	GetSchool(ctx context.Context, cnh *string) ([]models.School, error)
	GetSponsors(ctx context.Context, cnpj *string) ([]models.Sponsor, error)
}

type DriverRepository struct {
	db *sql.DB
}

func NewDriverRepository(db *sql.DB) *DriverRepository {
	return &DriverRepository{
		db: db,
	}
}

func (dr *DriverRepository) GetSchool(ctx context.Context, cnh *string) ([]models.School, error) {

	return nil, nil

}

func (dr *DriverRepository) GetSponsors(ctx context.Context, cnpj *string) ([]models.Sponsor, error) {

	return nil, nil

}
