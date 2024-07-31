package service

import (
	"context"
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"
	"github.com/venture-technology/vtx-account-manager/config"
	"github.com/venture-technology/vtx-account-manager/internal/repository"
	"github.com/venture-technology/vtx-account-manager/models"
)

func newPostgres(dbConfig config.Database) string {
	return "user=" + dbConfig.User +
		" password=" + dbConfig.Password +
		" dbname=" + dbConfig.Name +
		" host=" + dbConfig.Host +
		" port=" + dbConfig.Port +
		" sslmode=disable"
}

func mockDriver() *models.Driver {
	return &models.Driver{
		Name:  "João Silva",
		Email: "gustavorodrigueslima2004@gmail.com",
		CNH:   "26779665567",
	}
}

func setupDriverTestDb(t *testing.T) (*sql.DB, *DriverService) {

	t.Helper()

	config, err := config.Load("../../config/config.yaml")
	if err != nil {
		t.Fatalf("falha ao carregar a configuração: %v", err)
	}

	db, err := sql.Open("postgres", newPostgres(config.Database))
	if err != nil {
		t.Fatalf("falha ao conectar ao banco de dados: %v", err)
	}

	driverRepository := repository.NewDriverRepository(db)
	driverService := NewDriverService(driverRepository)

	return db, driverService

}

func TestCreatePartner(t *testing.T) {

	db, driverService := setupDriverTestDb(t)
	defer db.Close()

	driverMock := mockDriver()
	schoolMock := mockSchool()

	partner := models.Partner{
		Driver: *driverMock,
		School: *schoolMock,
	}

	err := driverService.CreatePartner(context.Background(), &partner)

	if err != nil {
		t.Errorf("Erro ao criar parceria: %v", err)
	}

}

func TestDriverGetPartners(t *testing.T) {

	db, driverService := setupDriverTestDb(t)
	defer db.Close()

	driverMock := mockDriver()

	_, err := driverService.GetPartners(context.Background(), &driverMock.CNH)

	if err != nil {
		t.Errorf("Erro ao encontrar lista de parceiros: %v", err.Error())
	}
}

func TestDriverGetSchool(t *testing.T) {

	db, driverService := setupDriverTestDb(t)
	defer db.Close()

	driverMock := mockDriver()
	schoolMock := mockSchool()

	_, err := driverService.GetSchool(context.Background(), &driverMock.CNH, &schoolMock.CNPJ)
	if err != nil {
		t.Errorf("Erro ao encontrar parceria: %v", err.Error())
	}

}

func TestDriverGetContract(t *testing.T) {
	log.Print("yet not implemented")
}

func TestDriverGetContractByShift(t *testing.T) {
	log.Print("yet not implemented")
}
