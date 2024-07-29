package repository

import (
	"context"
	"database/sql"

	"github.com/venture-technology/vtx-account-manager/models"
)

type IResponsibleRepository interface {
	FindAllDriverAtSchool(ctx context.Context, cnpj *string) ([]models.Driver, error)
	CreateContract(ctx context.Context, contract *models.Contract) error
	GetPartners(ctx context.Context, cpf *string) ([]models.Contract, error)
	BreachContract(ctx context.Context, record *int) error
}

type ResponsibleRepository struct {
	db *sql.DB
}

func NewResponsibleRepository(db *sql.DB) *ResponsibleRepository {
	return &ResponsibleRepository{
		db: db,
	}
}

func (rr *ResponsibleRepository) CreateContract(ctx context.Context, contract *models.Contract) error {

	sqlQuery := `INSERT INTO contracts (
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
		contract.Driver.Name,
		contract.Driver.CNH,
		contract.Driver.Email,
		contract.School.Name,
		contract.School.CNPJ,
		contract.School.Email,
		contract.Child.Responsible.Name,
		contract.Child.Responsible.CPF,
		contract.Child.Responsible.Email,
		contract.Child.Responsible.Street,
		contract.Child.Responsible.Complement,
		contract.Child.Responsible.ZIP,
		contract.Child.Name,
		contract.Child.RG,
		contract.Child.Shift,
		contract.CreatedAt,
	)

	return err

}

func (rr *ResponsibleRepository) GetPartners(ctx context.Context, cpf *string) ([]models.Contract, error) {

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
		created_at FROM contracts WHERE cpf_responsible = $1`

	rows, err := rr.db.Query(sqlQuery, cpf)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contracts []models.Contract

	for rows.Next() {
		var contract models.Contract

		err := rows.Scan(
			&contract.Driver.Name,
			&contract.Driver.CNH,
			&contract.Driver.Email,
			&contract.School.Name,
			&contract.School.CNPJ,
			&contract.School.Email,
			&contract.Child.Responsible.Name,
			&contract.Child.Responsible.CPF,
			&contract.Child.Responsible.Email,
			&contract.Child.Responsible.Street,
			&contract.Child.Responsible.Complement,
			&contract.Child.Responsible.ZIP,
			&contract.Child.Name,
			&contract.Child.RG,
			&contract.Child.Shift,
			&contract.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		contracts = append(contracts, contract)
	}

	if err := rows.Scan(); err != nil {
		return nil, err
	}

	return contracts, nil

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

func (rr *ResponsibleRepository) BreachContract(ctx context.Context, record *int) error {

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
	_, err = tx.Exec("DELETE FROM contracts WHERE record = $1", record)
	return err

}
