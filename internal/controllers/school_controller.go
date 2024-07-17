package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/vtx-account-manager/internal/server"
	"github.com/venture-technology/vtx-account-manager/internal/service"
)

type SchoolController struct {
	schoolservice *service.SchoolService
}

func NewSchoolController(schoolservice *service.SchoolService) *SchoolController {
	return &SchoolController{
		schoolservice: schoolservice,
	}
}

func (ct *SchoolController) RegisterRoutes(router *gin.Engine) {

	api := router.Group("vtx-account-manager/api/v1")

	api.GET("/:cnpj/driver", ct.GetDriver)   // para visualizar todos seus motoristas
	api.GET("/:cnpj/sponsor", ct.GetSponsor) // para visualizar todos os sponsors

}

func (ct *SchoolController) GetDriver(c *gin.Context) {

	cnpj := c.Param("cnpj")

	drivers, err := ct.schoolservice.GetDriver(c, &cnpj)

	if err != nil {
		log.Printf("error to find drivers: %s", err.Error())
		c.JSON(http.StatusBadRequest, server.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, drivers)

}

func (ct *SchoolController) GetSponsor(c *gin.Context) {

	cnpj := c.Param("cnpj")

	sponsors, err := ct.schoolservice.GetSponsors(c, &cnpj)

	if err != nil {
		log.Printf("error to find sponsors: %s", err.Error())
		c.JSON(http.StatusBadRequest, server.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, sponsors)

}
