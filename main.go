package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/vtx-account-manager/config"
	"github.com/venture-technology/vtx-account-manager/internal/controllers"
	"github.com/venture-technology/vtx-account-manager/internal/repository"
	"github.com/venture-technology/vtx-account-manager/internal/service"

	_ "github.com/lib/pq"
)

func main() {

	router := gin.Default()

	config, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := sql.Open("postgres", newPostgres(config.Database))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = migrate(db, config.Database.Schema)
	if err != nil {
		log.Fatalf("failed to execute migrations: %v", err)
	}

	DriverRepository := repository.NewDriverRepository(db)
	responsibleRepository := repository.NewResponsibleRepository(db)
	schoolRepository := repository.NewSchoolRepository(db)

	DriverService := service.NewDriverService(DriverRepository)
	responsibleService := service.NewResponsibleService(responsibleRepository)
	schoolService := service.NewSchoolService(schoolRepository)

	DriverController := controllers.NewDriverController(DriverService)
	responsibleController := controllers.NewResponsibleController(responsibleService)
	schoolController := controllers.NewSchoolController(schoolService)

	log.Print(responsibleController, schoolController)

	DriverController.RegisterRoutes(router)
	// responsibleController.RegisterRoutes(router)
	schoolController.RegisterRoutes(router)

	log.Printf("initing service: %s", config.Name)
	router.Run(fmt.Sprintf(":%d", config.Server.Port))

}

func newPostgres(dbConfig config.Database) string {
	return "user=" + dbConfig.User +
		" password=" + dbConfig.Password +
		" dbname=" + dbConfig.Name +
		" host=" + dbConfig.Host +
		" port=" + dbConfig.Port +
		" sslmode=disable"
}

func migrate(db *sql.DB, filepath string) error {
	schema, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return err
	}

	return nil
}
