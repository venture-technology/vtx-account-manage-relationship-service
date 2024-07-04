package repository

import "database/sql"

type IResponsibleRepository interface {
	FindAllDriverAtSchool()
	CreateSponsor()
	GetPartners()
	BreachSponsor()
}

type ResponsibleRepository struct {
	db *sql.DB
}

func NewResponsibleRepository(db *sql.DB) *ResponsibleRepository {
	return &ResponsibleRepository{
		db: db,
	}
}

func (rr *ResponsibleRepository) CreateSponsor() {

}

func (rr *ResponsibleRepository) GetPartners() {

}

func (rr *ResponsibleRepository) FindAllDriverAtSchool() {

}

func (rr *ResponsibleRepository) BreachSponsor() {

}
