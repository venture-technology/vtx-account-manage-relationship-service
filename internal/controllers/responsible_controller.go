package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/vtx-account-manager/internal/server"
	"github.com/venture-technology/vtx-account-manager/internal/service"
	"github.com/venture-technology/vtx-account-manager/models"
)

type ResponsibleController struct {
	responsibleservice *service.ResponsibleService
}

func NewResponsibleController(responsibleservice *service.ResponsibleService) *ResponsibleController {
	return &ResponsibleController{
		responsibleservice: responsibleservice,
	}
}

func (ct *ResponsibleController) RegisterRoutes(router *gin.Engine) {
	api := router.Group("vtx-account-manager/api/v1/responsible")

	api.GET("/school/:cnpj/driver", ct.SearchDriversInSchool) // para encontrar motoristas da escola
	api.POST("/contract", ct.CreateContract)                  // para fechar um contrato com o motorista e escola
	api.GET("/contract/:cpf", ct.GetPartners)                 // para visualizar todos os motoristas
	api.DELETE("/contract/:record", ct.BreachContract)        // para quebrar um contrato com o motorista e escola
}

func (ct *ResponsibleController) CreateContract(c *gin.Context) {

	var input models.Contract

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, server.InternalServerErrorResponse(err))
		return
	}

	err := ct.responsibleservice.CreateContract(c, &input)

	if err != nil {
		log.Printf("error to create contract: %s", err.Error())
		c.JSON(http.StatusBadRequest, server.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, input)

}

func (ct *ResponsibleController) GetPartners(c *gin.Context) {

	cpf := c.Param("cpf")

	partners, err := ct.responsibleservice.GetPartners(c, &cpf)

	if err != nil {
		log.Printf("error to find partners: %s", err.Error())
		c.JSON(http.StatusBadRequest, server.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, partners)

}

func (ct *ResponsibleController) BreachContract(c *gin.Context) {

	record, _ := strconv.Atoi(c.Param("record"))

	err := ct.responsibleservice.BreachContract(c, &record)

	if err != nil {
		log.Printf("error to breach contract: %s", err.Error())
		c.JSON(http.StatusBadRequest, server.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted with successfully"})

}

func (ct *ResponsibleController) SearchDriversInSchool(c *gin.Context) {

	cnpj := c.Param("cnpj")

	drivers, err := ct.responsibleservice.FindAllDriverAtSchool(c, &cnpj)

	if err != nil {
		log.Printf("error to find drivers at the school: %s", err.Error())
		c.JSON(http.StatusBadRequest, server.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, drivers)

}
