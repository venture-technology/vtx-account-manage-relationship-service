package service

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/venture-technology/vtx-account-manager/config"
	"github.com/venture-technology/vtx-account-manager/internal/repository"
)

func setupTestDb(t *testing.T) (*sql.DB, *ResponsibleService) {
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

// func TestCalculateValueSubscription(t *testing.T) {
// 	db, responsibleService := setupTestDb(t)
// 	defer db.Close()

// 	dist, err := responsibleService.getDistance(context.Background(), "Avenida Itamerendiba, 30, 08120520", "Avenida Barão de Alagoas, 223, 08120000")

// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}

// 	log.Print("dist", dist)
// }
