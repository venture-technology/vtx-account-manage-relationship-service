package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/vtx-account-manager/internal/middleware"
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
	api := router.Group("vtx-account-manager/api/v1")

	api.GET("/school/:cnpj/driver", middleware.ResponsibleMiddleware(), ct.SearchDriversInSchool) // para encontrar motoristas da escola
	api.POST("/sponsor", middleware.ResponsibleMiddleware(), ct.CreateSponsor)                    // para fechar um contrato com o motorista e escola
	api.DELETE("/sponsor", middleware.ResponsibleMiddleware(), ct.BreachSponsor)                  // para quebrar um contrato com o motorista e escola
	api.GET("/sponsor", middleware.ResponsibleMiddleware(), ct.GetSponsors)                       // para visualizar todos os motoristas
}

func (ct *ResponsibleController) CreateSponsor(c *gin.Context) {

	var input models.Sponsor

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, server.InternalServerErrorResponse(err))
		return
	}

	err := ct.responsibleservice.CreateSponsor(c, &input)

	if err != nil {
		log.Printf("error to create sponsor: %s", err.Error())
		c.JSON(http.StatusBadRequest, server.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, input)

}

func (ct *ResponsibleController) GetSponsors(c *gin.Context) {

}

func (ct *ResponsibleController) BreachSponsor(c *gin.Context) {

}

func (ct *ResponsibleController) SearchDriversInSchool(c *gin.Context) {

}
