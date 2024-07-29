package repository

import (
	"context"
	"database/sql"

	"github.com/venture-technology/vtx-account-manager/models"
)

type ISchoolRepository interface {
	GetAllDriversToSchool(ctx context.Context, cnpj *string) ([]models.Handshake, error)
	GetContracts(ctx context.Context, cnpj *string) ([]models.Contract, error)
	DeletePartner(ctx context.Context, cnpj, cnh *string) error
}

type SchoolRepository struct {
	db *sql.DB
}

func NewSchoolRepository(db *sql.DB) *SchoolRepository {
	return &SchoolRepository{
		db: db,
	}
}

func (sr *SchoolRepository) GetAllDriversToSchool(ctx context.Context, cnpj *string) ([]models.Handshake, error) {

	sqlQuery := `SELECT record, name_driver, cnh_driver, email_driver, created_at FROM partners WHERE cnpj_school = $1`

	rows, err := sr.db.Query(sqlQuery, cnpj)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var partners []models.Handshake

	for rows.Next() {
		var partner models.Handshake

		err := rows.Scan(
			&partner.Record,
			&partner.Driver.Name,
			&partner.Driver.CNH,
			&partner.Driver.Email,
			&partner.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		partners = append(partners, partner)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return partners, nil

}

func (sr *SchoolRepository) GetContracts(ctx context.Context, cnpj *string) ([]models.Contract, error) {

	sqlQuery := `SELECT record, name_driver, cnh_driver, email_driver, name_school, cnpj_school, email_school, name_responsible, cpf_responsible, email_responsible, street_responsible, number_responsible, zip_responsible, name_child, rg_child, shift, created_at FROM contracts WHERE cnpj_school = $1`

	rows, err := sr.db.Query(sqlQuery, cnpj)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contracts []models.Contract

	for rows.Next() {

		var contract models.Contract

		err := rows.Scan(
			&contract.Record,
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
			&contract.Child.Responsible.Number,
			&contract.Child.Responsible.Complement,
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

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return contracts, nil

}

func (sr *SchoolRepository) DeletePartner(ctx context.Context, cnpj, cnh *string) error {

	tx, err := sr.db.Begin()
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
	_, err = tx.Exec("DELETE FROM partners WHERE cnh_driver = $1 AND cnpj_school = $2", cnh, cnpj)
	return err

}
