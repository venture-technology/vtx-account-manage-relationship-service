package repository

import "database/sql"

type IDriverRepository interface {
	GetSchool()
	GetSponsors()
}

type DriverRepository struct {
	db *sql.DB
}

func NewDriverRepository(db *sql.DB) *DriverRepository {
	return &DriverRepository{
		db: db,
	}
}

func (dr *DriverRepository) GetSchool() {

}

func (dr *DriverRepository) GetSponsors() {

}
