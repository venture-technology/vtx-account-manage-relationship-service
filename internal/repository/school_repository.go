package repository

import (
	"context"
	"database/sql"

	"github.com/venture-technology/vtx-account-manager/models"
)

type ISchoolRepository interface {
	GetDriver(ctx context.Context, cnpj *string) ([]models.Handshake, error)
	GetSponsors(ctx context.Context, cnpj *string) ([]models.Sponsor, error)
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

	sqlQuery := `SELECT record, name_driver, cnh_driver, email_driver, created_at FROM partners WHERE cnpj_driver = $1`

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

func (sr *SchoolRepository) GetSponsors(ctx context.Context, cnpj *string) ([]models.Sponsor, error) {

	sqlQuery := `SELECT record, name_driver, cnh_driver, email_driver, name_school, cnpj_school, email_school, name_responsible, cpf_responsible, email_responsible, street_responsible, number_responsible, zip_responsible, name_child, rg_child, shift, created_at FROM sponsors WHERE cnpj_school = $1`

	rows, err := sr.db.Query(sqlQuery, cnpj)
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

func (sr *SchoolRepository) DeletePartner(ctx context.Context, cnh *string) error {

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
	_, err = tx.Exec("DELETE FROM partners WHERE cnh_driver = $1", cnh)
	return err

}
