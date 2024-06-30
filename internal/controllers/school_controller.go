package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/vtx-account-manager/internal/middleware"
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

	api := router.Group("api/v1/vtx-account-manager")

	api.GET("/driver", middleware.SchoolMiddleware(), ct.GetDriver)   // para visualizar todos seus motoristas
	api.GET("/sponsor", middleware.SchoolMiddleware(), ct.GetSponsor) // para visualizar todos os sponsors
	api.GET(":cnpj/school", func(c *gin.Context) {
		driver := c.Query("driver")
		if driver == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "The param 'driver' not found"})
		}
	}, ct.IsDriver)

}

func (ct *SchoolController) GetDriver(c *gin.Context) {}

func (ct *SchoolController) GetSponsor(c *gin.Context) {}

func (ct *SchoolController) IsDriver(c *gin.Context) {}
