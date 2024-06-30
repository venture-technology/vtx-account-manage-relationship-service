package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/venture-technology/vtx-account-manager/internal/service"
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

	api := router.Group("api/v1/vtx-account-manager")

	api.GET("/school/:cpf/driver")    // para encontrar motoristas da escola
	api.POST("/school/:cpf/driver")   // para fechar um contrato com o motorista e escola
	api.DELETE("/school/:cpf/driver") // para quebrar um contrato com o motorista e escola
	api.GET("/sponsors")              // para visualizar todos os motoristas

}
