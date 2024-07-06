package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/vtx-account-manager/internal/middleware"
	"github.com/venture-technology/vtx-account-manager/internal/service"
	"github.com/venture-technology/vtx-account-manager/models"
)

type DriverController struct {
	driverservice *service.DriverService
}

func NewDriverController(driverservice *service.DriverService) *DriverController {
	return &DriverController{
		driverservice: driverservice,
	}
}

func (ct *DriverController) RegisterRoutes(router *gin.Engine) {

	api := router.Group("api/v1/vtx-account-manager")

	api.POST("/partner", ct.CreatePartner)
	api.GET("/school", middleware.DriverMiddleware(), ct.GetSchool)   // para visualizar todas as suas escolas
	api.GET("/sponsor", middleware.DriverMiddleware(), ct.GetSponsor) // para visualizar todos os sponsors
	api.GET(":cnh/school", func(c *gin.Context) {
		school := c.Query("school")
		if school == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "The param 'school' not found"})
		}
	}, ct.IsSchoolArePartner)
}

func (ct *DriverController) CreatePartner(c *gin.Context) {

	var input models.Handshake

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	err := ct.driverservice.CreatePartner(c, &input)

	if err != nil {
		log.Printf("error to create partners: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error at create partner"})
		return
	}

	c.JSON(http.StatusCreated, input)

}

func (ct *DriverController) GetSchool(c *gin.Context) {
}

func (ct *DriverController) GetSponsor(c *gin.Context) {

}

func (ct *DriverController) IsSchoolArePartner(c *gin.Context) {

}
