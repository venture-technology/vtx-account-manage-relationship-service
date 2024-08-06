package service

import (
	"fmt"
	"log"
	"testing"

	"github.com/venture-technology/vtx-account-manager/models"
)

func mockContract() *models.Contract {
	return &models.Contract{
		Status:      "Active",
		Description: "Test Contract",
		Driver:      *mockDriver(),
		School:      *mockSchool(),
		Child:       *mockChild(),
	}

}

func TestGetDistance(t *testing.T) {

	db, responsibleService := setupResponsibleTestDb(t)
	defer db.Close()

	contract := mockContract()

	origin := fmt.Sprintf("%s, %s, %s", contract.Child.Responsible.Street, contract.Child.Responsible.Number, contract.Child.Responsible.ZIP)
	destination := fmt.Sprintf("%s, %s, %s", contract.School.Street, contract.School.Number, contract.School.ZIP)

	distance, err := responsibleService.GetDistance(origin, destination)

	if err != nil {
		t.Errorf(err.Error())
	}

	log.Print("distance em km: ", *distance)

}

func TestCalculateValueSubscription(t *testing.T) {

	db, responsibleService := setupResponsibleTestDb(t)
	defer db.Close()

	value := responsibleService.CalculateContractValue(3.50, 25)

	log.Print(value)

}

func TestCreateProduct(t *testing.T) {

	db, responsibleService := setupResponsibleTestDb(t)
	defer db.Close()

	contract := mockContract()

	prod, err := responsibleService.CreateProduct(contract)

	if err != nil {
		t.Errorf(err.Error())
	}

	log.Print(prod.ID)

}

func TestCreatePrice(t *testing.T) {

	db, responsibleService := setupResponsibleTestDb(t)
	defer db.Close()

	contract := mockContract()
	contract.StripeSubscription.ProductSubscriptionId = "prod_QblxXRrYBIIc9y"

	amount := responsibleService.CalculateContractValue(3.50, 25)

	contract.Amount = int64(amount)

	pr, err := responsibleService.CreatePrice(contract)

	if err != nil {
		t.Errorf(err.Error())
	}

	log.Print(pr.ID)

}

func TestCreateSubscription(t *testing.T) {

	db, responsibleService := setupResponsibleTestDb(t)
	defer db.Close()

	contract := mockContract()
	contract.Child.Responsible.CustomerId = "cus_QXeuluwEfuvSnt"
	contract.StripeSubscription.PriceSubscriptionId = "price_1PkYeqLfFDLpePGLnJyo5YjD"

	subs, err := responsibleService.CreateSubscription(contract)

	if err != nil {
		t.Errorf(err.Error())
	}

	log.Print(subs.ID)

}

func TestListSubscriptions(t *testing.T) {

	db, responsibleService := setupResponsibleTestDb(t)
	defer db.Close()

	contract := mockContract()
	contract.Child.Responsible.CustomerId = "cus_QXeuluwEfuvSnt"

	subs, err := responsibleService.ListSubscriptions(contract)

	if err != nil {
		t.Errorf(err.Error())
	}

	log.Print(subs)

}

func TestGetSubscription(t *testing.T) {

	db, responsibleService := setupResponsibleTestDb(t)
	defer db.Close()

	contract := mockContract()
	contract.StripeSubscription.ProductSubscriptionId = "sub_1Pkb11LfFDLpePGLhwaB2vL8"

	sub, err := responsibleService.GetSubscription("sub_1Pkb11LfFDLpePGLhwaB2vL8")

	if err != nil {
		t.Errorf(err.Error())
	}

	log.Print(sub)

}

func TestDeleteSubscription(t *testing.T) {

	db, responsibleService := setupResponsibleTestDb(t)
	defer db.Close()

	contract := mockContract()
	contract.StripeSubscription.SubscriptionId = "sub_1Pkb11LfFDLpePGLhwaB2vL8"

	delSub, err := responsibleService.DeleteSubscription(contract)

	if err != nil {
		t.Errorf(err.Error())
	}

	log.Print(delSub)

}

func TestGetInvoice(t *testing.T) {

	db, responsibleService := setupResponsibleTestDb(t)
	defer db.Close()

	inv, err := responsibleService.GetInvoice("in_1PkanDLfFDLpePGL56BzursV")
	if err != nil {
		t.Errorf(err.Error())
	}

	log.Print(inv)

}

func TestListInvoices(t *testing.T) {

	db, responsibleService := setupResponsibleTestDb(t)
	defer db.Close()

	contract := mockContract()
	contract.StripeSubscription.SubscriptionId = "sub_1Pkb11LfFDLpePGLhwaB2vL8"

	invoices, err := responsibleService.ListInvoices(contract)

	if err != nil {
		t.Errorf(err.Error())
	}

	log.Print(invoices)

}

func TestCalculateRemainingValueSubscription(t *testing.T) {

	db, responsibleService := setupResponsibleTestDb(t)
	defer db.Close()

	contract := mockContract()
	contract.StripeSubscription.SubscriptionId = "sub_1Pkb11LfFDLpePGLhwaB2vL8"

	invoices, err := responsibleService.ListInvoices(contract)

	if err != nil {
		t.Errorf(err.Error())
	}

	val := responsibleService.CalculateRemainingValueSubscription(invoices)

	log.Print(val)
}

func TestFineResponsible(t *testing.T) {

	db, responsibleService := setupResponsibleTestDb(t)
	defer db.Close()

	contract := mockContract()
	contract.StripeSubscription.SubscriptionId = "sub_1Pkb11LfFDLpePGLhwaB2vL8"

	invoices, err := responsibleService.ListInvoices(contract)

	if err != nil {
		t.Errorf(err.Error())
	}

	amount := responsibleService.CalculateRemainingValueSubscription(invoices)

	log.Print(amount)

	pi, err := responsibleService.FineResponsible(contract, int64(amount.Fines))

	if err != nil {
		t.Errorf(err.Error())
	}

	log.Print(pi.Amount, pi.ID)

}
