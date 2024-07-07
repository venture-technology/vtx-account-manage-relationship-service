package repository

import (
	"context"
	"database/sql"

	"github.com/venture-technology/vtx-account-manager/models"
)

type IDriverRepository interface {
	GetSchool(ctx context.Context, cnh, cnpj *string) (*models.Handshake, error)
	GetSponsors(ctx context.Context, cnh *string) ([]models.Sponsor, error)
	GetSponsorsByShift(ctx context.Context, cnh, shift *string) ([]models.Sponsor, error)
	CreatePartner(ctx context.Context, handshake *models.Handshake) error
	GetPartners(ctx context.Context, cnh *string) ([]models.Handshake, error)
}

type DriverRepository struct {
	db *sql.DB
}

func NewDriverRepository(db *sql.DB) *DriverRepository {
	return &DriverRepository{
		db: db,
	}
}

func (dr *DriverRepository) GetSchool(ctx context.Context, cnh, cnpj *string) (*models.Handshake, error) {

	sqlQuery := `SELECT record, name_driver, cnh_driver, email_driver, name_school, cnpj_school, email_school, created_at FROM partners WHERE cnh_driver = $1 AND cnpj_school = $2 LIMIT 1`

	var partner models.Handshake

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

func (dr *DriverRepository) GetSponsors(ctx context.Context, cnh *string) ([]models.Sponsor, error) {

	sqlQuery := `SELECT record, name_driver, cnh_driver, email_driver, name_school, cnpj_school, email_school, name_responsible, cpf_responsible, email_responsible, street_responsible, number_responsible, zip_responsible, name_child, rg_child, shift, created_at FROM sponsors WHERE cnh_driver = $1`

	rows, err := dr.db.Query(sqlQuery, cnh)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sponsors []models.Sponsor

	for rows.Next() {

		var sponsor models.Sponsor

		err := rows.Scan(
			&sponsor.Record,
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
			&sponsor.Child.Responsible.Number,
			&sponsor.Child.Responsible.Complement,
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

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sponsors, nil

}

func (dr *DriverRepository) CreatePartner(ctx context.Context, handshake *models.Handshake) error {

	sqlQuery := `INSERT INTO partners (name_driver, cnh_driver, email_driver, name_school, cnpj_school, email_school, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := dr.db.Exec(sqlQuery, handshake.Driver.Name, handshake.Driver.CNH, handshake.Driver.Email, handshake.School.Name, handshake.School.CNPJ, handshake.School.Email)

	return err
}

func (dr *DriverRepository) GetPartners(ctx context.Context, cnh *string) ([]models.Handshake, error) {

	sqlQuery := `SELECT record, name_school, cnpj_school, email_school, created_at FROM partners WHERE cnh_driver = $1`

	rows, err := dr.db.Query(sqlQuery, cnh)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var partners []models.Handshake

	for rows.Next() {
		var partner models.Handshake

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

func (dr *DriverRepository) GetSponsorsByShift(ctx context.Context, cnh, shift *string) ([]models.Sponsor, error) {

	sqlQuery := `SELECT record, name_driver, cnh_driver, email_driver, name_school, cnpj_school, email_school, name_responsible, cpf_responsible, email_responsible, street_responsible, number_responsible, zip_responsible, name_child, rg_child, shift, created_at FROM sponsors WHERE cnh_driver = $1 AND shift = $2`

	rows, err := dr.db.Query(sqlQuery, cnh, shift)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sponsors []models.Sponsor

	for rows.Next() {

		var sponsor models.Sponsor

		err := rows.Scan(
			&sponsor.Record,
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
			&sponsor.Child.Responsible.Number,
			&sponsor.Child.Responsible.Complement,
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

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sponsors, nil

}
