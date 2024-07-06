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

	api.GET("/:cnh", ct.GetSchool) // para verificar se uma escola é parceira de um motorista
	api.POST("/partner", ct.CreatePartner)
	api.GET("/:cnh/school", middleware.DriverMiddleware(), ct.GetPartners)       // para visualizar todas as suas escolas
	api.GET("/:cnh/sponsor", middleware.DriverMiddleware(), ct.GetSponsor)       // para visualizar todos os sponsors
	api.GET("/:cnh/shift", middleware.DriverMiddleware(), ct.GetSponsorsByShift) // para buscar todos os sponsors de acordo com o horário da escola
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

func (ct *DriverController) GetPartners(c *gin.Context) {

	cnh := c.Param("cnh")

	partners, err := ct.driverservice.GetPartners(c, &cnh)

	if err != nil {
		log.Printf("error while found partners: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "partners don't found"})
		return
	}

	c.JSON(http.StatusOK, partners)

}

func (ct *DriverController) GetSchool(c *gin.Context) {

	cnpj := c.Query("school")

	if cnpj == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The param 'school' not found"})
	}

	cnh := c.Param("cnh")

	partner, err := ct.driverservice.GetSchool(c, &cnh, &cnpj)
	if err != nil {
		log.Printf("error while found school: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "school don't found"})
		return
	}

	c.JSON(http.StatusOK, partner)

}

func (ct *DriverController) GetSponsor(c *gin.Context) {

	cnh := c.Param("cnh")

	sponsors, err := ct.driverservice.GetSponsors(c, &cnh)
	if err != nil {
		log.Printf("error while found sponsors: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "sponsors don't found"})
		return
	}

	c.JSON(http.StatusOK, sponsors)

}

func (ct *DriverController) GetSponsorsByShift(c *gin.Context) {

	shift := c.Query("shift")
	if shift == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The param 'shift' not found"})
	}

	cnh := c.Param("cnh")

	sponsors, err := ct.driverservice.GetSponsorsByShift(c, &cnh, &shift)
	if err != nil {
		log.Printf("error while found sponsors: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "sponsors don't found"})
		return
	}

	c.JSON(http.StatusOK, sponsors)

}
