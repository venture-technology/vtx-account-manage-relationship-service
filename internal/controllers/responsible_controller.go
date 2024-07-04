package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/venture-technology/vtx-account-manager/internal/middleware"
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

	api.GET("/school/:cnpj/driver", middleware.ResponsibleMiddleware(), ct.SearchDriversInSchool) // para encontrar motoristas da escola
	api.POST("/school/:cnpj/driver", middleware.ResponsibleMiddleware(), ct.CreateSponsor)        // para fechar um contrato com o motorista e escola
	api.DELETE("/school/:cnpj/driver", middleware.ResponsibleMiddleware(), ct.BreachSponsor)      // para quebrar um contrato com o motorista e escola
	api.GET("/sponsors", middleware.ResponsibleMiddleware(), ct.GetSponsors)                      // para visualizar todos os motoristas
}

func (ct *ResponsibleController) CreateSponsor(c *gin.Context) {
}

func (ct *ResponsibleController) BreachSponsor(c *gin.Context) {
}

func (ct *ResponsibleController) SearchDriversInSchool(c *gin.Context) {
}

func (ct *ResponsibleController) GetSponsors(c *gin.Context) {
}
