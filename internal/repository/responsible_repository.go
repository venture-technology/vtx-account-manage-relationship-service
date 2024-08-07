package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/venture-technology/vtx-account-manager/models"
)

type IResponsibleRepository interface {
	FindAllDriverAtSchool(ctx context.Context, cnpj *string) ([]models.Driver, error)
	CreateContract(ctx context.Context, contract *models.Contract) error
	FindContractsByCpf(ctx context.Context, cpf, status *string) ([]models.Contract, error)
	DeleteContract(ctx context.Context, record uuid.UUID) error
	ExpireContract(ctx context.Context, record uuid.UUID) error
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
		record,
		title_stripe_subscription,
		description_stripe_subscription,
		id_stripe_subsciption,
		id_price_subscription,
		id_product_subscription,
		name_school,
		cnpj_school,
		email_school,
		street_school,
		complement_school,
		zip_school,
		name_driver,
		cnh_driver,
		email_driver,
		street_driver,
		complement_driver,
		zip_driver,
		name_responsible,
		cpf_responsible,
		customer_id_responsible,
		payment_method_id_responsible,
		email_responsible,
		street_responsible,
		complement_responsible,
		zip_responsible,
		name_child,
		rg_child,
		shift,
		expire_at,
		status
	)
	VALUES ($1, $2. $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, "currently")`

	_, err := rr.db.Exec(sqlQuery,
		contract.Record,
		contract.StripeSubscription.Title,
		contract.Description,
		contract.StripeSubscription.SubscriptionId,
		contract.StripeSubscription.PriceSubscriptionId,
		contract.StripeSubscription.ProductSubscriptionId,
		contract.School.Name,
		contract.School.CNPJ,
		contract.School.Email,
		contract.School.Street,
		contract.School.Complement,
		contract.School.ZIP,
		contract.Driver.Name,
		contract.Driver.CNH,
		contract.Driver.Email,
		contract.Driver.Street,
		contract.Driver.Complement,
		contract.Driver.ZIP,
		contract.Child.Responsible.Name,
		contract.Child.Responsible.CPF,
		contract.Child.Responsible.CustomerId,
		contract.Child.Responsible.PaymentMethodId,
		contract.Child.Responsible.Email,
		contract.Child.Responsible.Street,
		contract.Child.Responsible.Complement,
		contract.Child.Responsible.ZIP,
		contract.Child.Name,
		contract.Child.RG,
		contract.Child.Shift,
		contract.ExpireAt,
	)

	return err

}

func (rr *ResponsibleRepository) FindContractsByCpf(ctx context.Context, cpf, status *string) ([]models.Contract, error) {

	sqlQuery := `SELECT 
		record,
		title_stripe_subscription,
		description_stripe_subscription,
		id_stripe_subsciption,
		id_price_subscription,
		id_product_subscription,
		name_school,
		cnpj_school,
		email_school,
		street_school,
		complement_school,
		zip_school,
		name_driver,
		cnh_driver,
		email_driver,
		street_driver,
		complement_driver,
		zip_driver,
		name_responsible,
		cpf_responsible,
		customer_id_responsible,
		payment_method_id_responsible,
		email_responsible,
		street_responsible,
		complement_responsible,
		zip_responsible,
		name_child,
		rg_child,
		shift,
		created_at,
		expire_at, 
		status FROM contracts WHERE cpf_responsible = $1 AND status = $2`

	rows, err := rr.db.Query(sqlQuery, *cpf, *status)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contracts []models.Contract

	for rows.Next() {
		var contract models.Contract

		err := rows.Scan(
			&contract.Record,
			&contract.StripeSubscription.Title,
			&contract.Description,
			&contract.StripeSubscription.SubscriptionId,
			&contract.StripeSubscription.PriceSubscriptionId,
			&contract.StripeSubscription.ProductSubscriptionId,
			&contract.School.Name,
			&contract.School.CNPJ,
			&contract.School.Email,
			&contract.School.Street,
			&contract.School.Complement,
			&contract.School.ZIP,
			&contract.Driver.Name,
			&contract.Driver.CNH,
			&contract.Driver.Email,
			&contract.Driver.Street,
			&contract.Driver.Complement,
			&contract.Driver.ZIP,
			&contract.Child.Responsible.Name,
			&contract.Child.Responsible.CPF,
			&contract.Child.Responsible.CustomerId,
			&contract.Child.Responsible.PaymentMethodId,
			&contract.Child.Responsible.Email,
			&contract.Child.Responsible.Street,
			&contract.Child.Responsible.Complement,
			&contract.Child.Responsible.ZIP,
			&contract.Child.Name,
			&contract.Child.RG,
			&contract.Child.Shift,
			&contract.CreatedAt,
			&contract.ExpireAt,
			&contract.Status,
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

	sqlQuery := `SELECT name_driver, cnh_driver, email_driver FROM partners WHERE cnpj_school = $1`

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

func (rr *ResponsibleRepository) DeleteContract(ctx context.Context, record uuid.UUID) error {

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

func (rr *ResponsibleRepository) ExpireContract(ctx context.Context, record uuid.UUID) error {

	sqlQuery := `UPDATE contracts SET status = $1 WHERE record = $2`
	_, err := rr.db.ExecContext(ctx, sqlQuery, "expired")
	return err

}
