package service

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/venture-technology/vtx-account-manager/config"
	"github.com/venture-technology/vtx-account-manager/internal/repository"
	"github.com/venture-technology/vtx-account-manager/models"
)

func setupResponsibleTestDb(t *testing.T) (*sql.DB, *ResponsibleService) {
	t.Helper()

	config, err := config.Load("../../config/config.yaml")
	if err != nil {
		t.Fatalf("falha ao carregar a configuração: %v", err)
	}

	db, err := sql.Open("postgres", newPostgres(config.Database))
	if err != nil {
		t.Fatalf("falha ao conectar ao banco de dados: %v", err)
	}

	responsibleRepository := repository.NewResponsibleRepository(db)
	responsibleService := NewResponsibleService(responsibleRepository)

	return db, responsibleService
}

func mockChild() *models.Child {
	return &models.Child{
		Name: "Kauã Barbosa do Nascimento",
		RG:   "552381147",
		Responsible: models.Responsible{
			Name:            "Gesse Souza Lima",
			Email:           "gustavorodrigueslima2004@gmail.com",
			CPF:             "22321279826",
			Street:          "Rua Tiburcio de Souza",
			Number:          "2782",
			Complement:      "",
			ZIP:             "08140000",
			CustomerId:      "cus_QXeuluwEfuvSnt",
			PaymentMethodId: "pm_1PgZbILfFDLpePGLIZ5AOoIr",
		},
	}
}
