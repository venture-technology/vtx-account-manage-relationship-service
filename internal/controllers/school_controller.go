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

	api.GET("/:cnpj/driver", ct.GetAllDriversToSchool)       // para visualizar todos seus motoristas
	api.GET("/:cnpj/contract", ct.GetContract)               // para visualizar todos os contracts
	api.DELETE("/:cnpj/contract/:record", ct.DeleteContract) // para deletar uma parceria

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

func (ct *SchoolController) GetContract(c *gin.Context) {

	cnpj := c.Param("cnpj")

	contracts, err := ct.schoolservice.GetContracts(c, &cnpj)

	if err != nil {
		log.Printf("error to find contracts: %s", err.Error())
		c.JSON(http.StatusBadRequest, server.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, contracts)

}

func (ct *SchoolController) DeleteContract(c *gin.Context) {

	cnpj := c.Param("cnpj")
	cnh := c.Param("cnh")

	// necessário verificar se eles não tem nenhum contract antes, portanto uma escola não pode encerrar contrato com o motorista se tiver alunos registrados.

	err := ct.schoolservice.DeletePartner(c, &cnpj, &cnh)

	if err != nil {
		log.Printf("delete partner error: %s", err.Error())
		c.JSON(http.StatusBadRequest, server.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted w successfully"})

}
