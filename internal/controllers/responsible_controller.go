package controllers

import (
	"log"
	"net/http"

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

	api.GET("/find/:school/driver", ct.SearchDriversInSchool)       // para encontrar motoristas da escola
	api.POST("/contract", ct.CreateContract)                        // para criar um contrato
	api.GET("/contract/:cpf", ct.GetContractByCpf)                  // para verificar todos os contratos
	api.GET("/contract/:record", ct.GetContract)                    // para verificar um contrato em especifico
	api.GET("/contract/:record/invoice", ct.GetInvoiceFromContract) // para verificar todas as faturas daquele contrato
	api.GET("/contract/:record/invoice/:id", ct.GetInvoice)         // para verificar uma fatura de um contrato especifico
	api.PATCH("/contract/:record", ct.UpdateContract)               // para atualizar o metodo de pagamento de um contrato
	api.PATCH("/contract/webhook/expired", ct.UpdateStatusContract) // para atualizar e setar contrato como cancelado (vistado apenas por webhooks da stripe)
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

func (ct *ResponsibleController) GetContractByCpf(c *gin.Context) {

}

func (ct *ResponsibleController) GetContract(c *gin.Context) {

}

func (ct *ResponsibleController) GetInvoiceFromContract(c *gin.Context) {

}

func (ct *ResponsibleController) GetInvoice(c *gin.Context) {

}

func (ct *ResponsibleController) UpdateContract(c *gin.Context) {

}

func (ct *ResponsibleController) UpdateStatusContract(c *gin.Context) {

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
