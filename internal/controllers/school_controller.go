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

	api := router.Group("vtx-account-manager/api/v1/school")

	api.GET("/:cnpj/driver", ct.GetAllDriversToSchool)     // para visualizar todos seus motoristas
	api.GET("/:cnpj/sponsor", ct.GetSponsor)               // para visualizar todos os sponsors
	api.DELETE("/:cnpj/sponsor/:record", ct.DeleteSponsor) // para deletar uma parceria

}

func (ct *SchoolController) GetAllDriversToSchool(c *gin.Context) {

	cnpj := c.Param("cnpj")

	drivers, err := ct.schoolservice.GetAllDriversToSchool(c, &cnpj)

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

func (ct *SchoolController) DeleteSponsor(c *gin.Context) {

	cnpj := c.Param("cnpj")
	cnh := c.Param("cnh")

	// necessário verificar se eles não tem nenhum sponsor antes, portanto uma escola não pode encerrar contrato com o motorista se tiver alunos registrados.

	err := ct.schoolservice.DeletePartner(c, &cnpj, &cnh)

	if err != nil {
		log.Printf("delete partner error: %s", err.Error())
		c.JSON(http.StatusBadRequest, server.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted w successfully"})

}
