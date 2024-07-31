package repository

import (
	"context"
	"database/sql"

	"github.com/venture-technology/vtx-account-manager/models"
)

type IDriverRepository interface {
	GetSchool(ctx context.Context, cnh, cnpj *string) (*models.Partner, error)
	GetContracts(ctx context.Context, cnh *string) ([]models.Contract, error)
	GetContractsByShift(ctx context.Context, cnh, shift *string) ([]models.Contract, error)
	CreatePartner(ctx context.Context, partner *models.Partner) error
	GetPartners(ctx context.Context, cnh *string) ([]models.Partner, error)
}

type DriverRepository struct {
	db *sql.DB
}

func NewDriverRepository(db *sql.DB) *DriverRepository {
	return &DriverRepository{
		db: db,
	}
}

func (dr *DriverRepository) GetSchool(ctx context.Context, cnh, cnpj *string) (*models.Partner, error) {

	sqlQuery := `SELECT record, name_driver, cnh_driver, email_driver, name_school, cnpj_school, email_school, created_at FROM partners WHERE cnh_driver = $1 AND cnpj_school = $2 LIMIT 1`

	var partner models.Partner

	err := dr.db.QueryRow(sqlQuery, cnh, cnpj).Scan(
		&partner.Record,
		&partner.Driver.Name,
		&partner.Driver.CNH,
		&partner.Driver.Email,
		&partner.School.Name,
		&partner.School.CNPJ,
		&partner.School.Email,
		&partner.CreatedAt,
	)

	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}

	return &partner, nil

}

func (dr *DriverRepository) GetContracts(ctx context.Context, cnh *string) ([]models.Contract, error) {

	sqlQuery := `SELECT record, name_driver, cnh_driver, email_driver, name_school, cnpj_school, email_school, name_responsible, cpf_responsible, email_responsible, street_responsible, number_responsible, zip_responsible, name_child, rg_child, shift, created_at FROM contracts WHERE cnh_driver = $1`

	rows, err := dr.db.Query(sqlQuery, cnh)
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

func (dr *DriverRepository) CreatePartner(ctx context.Context, partner *models.Partner) error {

	sqlQuery := `INSERT INTO partners (name_driver, cnh_driver, email_driver, name_school, cnpj_school, email_school) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := dr.db.Exec(sqlQuery, partner.Driver.Name, partner.Driver.CNH, partner.Driver.Email, partner.School.Name, partner.School.CNPJ, partner.School.Email)

	return err
}

func (dr *DriverRepository) GetPartners(ctx context.Context, cnh *string) ([]models.Partner, error) {

	sqlQuery := `SELECT record, name_school, cnpj_school, email_school, created_at FROM partners WHERE cnh_driver = $1`

	rows, err := dr.db.Query(sqlQuery, cnh)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var partners []models.Partner

	for rows.Next() {
		var partner models.Partner

		err := rows.Scan(
			&partner.Record,
			&partner.School.Name,
			&partner.School.CNPJ,
			&partner.School.Email,
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

func (dr *DriverRepository) GetContractsByShift(ctx context.Context, cnh, shift *string) ([]models.Contract, error) {

	sqlQuery := `SELECT record, name_driver, cnh_driver, email_driver, name_school, cnpj_school, email_school, name_responsible, cpf_responsible, email_responsible, street_responsible, number_responsible, zip_responsible, name_child, rg_child, shift, created_at FROM contracts WHERE cnh_driver = $1 AND shift = $2`

	rows, err := dr.db.Query(sqlQuery, cnh, shift)
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
