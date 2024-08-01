package service

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/venture-technology/vtx-account-manager/config"
	"github.com/venture-technology/vtx-account-manager/internal/repository"
	"github.com/venture-technology/vtx-account-manager/models"
)

func mockSchool() *models.School {
	return &models.School{
		Name:  "E.M.E.F Professor Carlos Pasquale",
		Email: "gustavorodrigueslima2004@gmail.com",
		CNPJ:  "64025893000102",
	}
}

func setupSchoolTestDb(t *testing.T) (*sql.DB, *SchoolService) {

	t.Helper()

	config, err := config.Load("../../config/config.yaml")
	if err != nil {
		t.Fatalf("falha ao carregar a configuração: %v", err)
	}

	db, err := sql.Open("postgres", newPostgres(config.Database))
	if err != nil {
		t.Fatalf("falha ao conectar ao banco de dados: %v", err)
	}

	schoolRepository := repository.NewSchoolRepository(db)
	schoolService := NewSchoolService(schoolRepository)

	return db, schoolService

}

func TestSchoolGetAllDriversToSchool(t *testing.T) {

	db, schoolService := setupSchoolTestDb(t)
	defer db.Close()

	schoolMock := mockSchool()
	drivers, err := schoolService.GetAllDriversToSchool(context.Background(), &schoolMock.CNPJ)

	if err != nil {
		t.Errorf("Erro ao encontrar parceria: %v", err.Error())
	}

	log.Print(drivers)

}

func TestSchoolGetContract(t *testing.T) {
	log.Print("yet not implemented")
}

func TestSchoolDeletePartner(t *testing.T) {

	db, schoolService := setupSchoolTestDb(t)
	defer db.Close()

	schoolMock := mockSchool()
	driverMock := mockDriver()

	err := schoolService.DeletePartner(context.Background(), &schoolMock.CNPJ, &driverMock.CNH)

	if err != nil {
		t.Errorf("Erro ao deletar parceria: %v", err.Error())
	}

}
