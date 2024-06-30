package repository

import "database/sql"

type ISchoolRepository interface {
	GetDriver()
	GetSponsors()
}

type SchoolRepository struct {
	db *sql.DB
}

func NewSchoolRepository(db *sql.DB) *SchoolRepository {
	return &SchoolRepository{
		db: db,
	}
}
