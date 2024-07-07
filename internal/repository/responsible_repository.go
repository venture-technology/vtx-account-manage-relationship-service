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

	sqlQuery := `INSERT INTO sponsors (
		name_driver, 
		cnh_driver, 
		email_driver, 
		name_school, 
		cnpj_school, 
		email_school, 
		name_responsible, 
		cpf_responsible, 
		email_responsible, 
		street_responsible, 
		number_responsible, 
		complement_responsible, 
		zip_responsible, 
		name_child, 
		rg_child, 
		shift, 
		created_at
	)
	VALUES ($1, $2. $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)`

	_, err := rr.db.Exec(sqlQuery,
		sponsor.Driver.Name,
		sponsor.Driver.CNH,
		sponsor.Driver.Email,
		sponsor.School.Name,
		sponsor.School.CNPJ,
		sponsor.School.Email,
		sponsor.Child.Responsible.Name,
		sponsor.Child.Responsible.CPF,
		sponsor.Child.Responsible.Email,
		sponsor.Child.Responsible.Street,
		sponsor.Child.Responsible.Complement,
		sponsor.Child.Responsible.ZIP,
		sponsor.Child.Name,
		sponsor.Child.RG,
		sponsor.Child.Shift,
		sponsor.CreatedAt,
	)

	return err

}

func (rr *ResponsibleRepository) GetPartners(ctx context.Context, cpf *string) ([]models.Sponsor, error) {

	sqlQuery := `SELECT 
		name_driver, 
		cnh_driver, 
		email_driver, 
		name_school, 
		cnpj_school, 
		email_school, 
		name_responsible, 
		cpf_responsible, 
		email_responsible, 
		street_responsible, 
		number_responsible, 
		complement_responsible, 
		zip_responsible, 
		name_child, 
		rg_child, 
		shift, 
		created_at FROM sponsors WHERE cpf_responsible = $1`

	rows, err := rr.db.Query(sqlQuery, cpf)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sponsors []models.Sponsor

	for rows.Next() {
		var sponsor models.Sponsor

		err := rows.Scan(
			&sponsor.Driver.Name,
			&sponsor.Driver.CNH,
			&sponsor.Driver.Email,
			&sponsor.School.Name,
			&sponsor.School.CNPJ,
			&sponsor.School.Email,
			&sponsor.Child.Responsible.Name,
			&sponsor.Child.Responsible.CPF,
			&sponsor.Child.Responsible.Email,
			&sponsor.Child.Responsible.Street,
			&sponsor.Child.Responsible.Complement,
			&sponsor.Child.Responsible.ZIP,
			&sponsor.Child.Name,
			&sponsor.Child.RG,
			&sponsor.Child.Shift,
			&sponsor.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		sponsors = append(sponsors, sponsor)
	}

	if err := rows.Scan(); err != nil {
		return nil, err
	}

	return sponsors, nil

}

func (rr *ResponsibleRepository) FindAllDriverAtSchool(ctx context.Context, cnpj *string) ([]models.Driver, error) {

	sqlQuery := `SELECT record, name_driver, cnh_driver, email_driver FROM partners WHERE cnpj_school = $1`

	rows, err := rr.db.Query(sqlQuery, cnpj)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var drivers []models.Driver

	for rows.Next() {
		var driver models.Driver

		err := rows.Scan(
			&driver.Name,
			&driver.CNH,
			&driver.Email,
		)

		if err != nil {
			return nil, err
		}

		drivers = append(drivers, driver)
	}

	if err := rows.Scan(); err != nil {
		return nil, err
	}

	return drivers, nil

}

func (rr *ResponsibleRepository) BreachSponsor(ctx context.Context, record *int) error {

	tx, err := rr.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	_, err = tx.Exec("DELETE FROM sponsors WHERE record = $1", record)
	return err

}
