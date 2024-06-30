package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/vtx-account-manager/internal/middleware"
	"github.com/venture-technology/vtx-account-manager/internal/service"
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

	api.GET("/school", middleware.DriverMiddleware(), ct.GetSchool)   // para visualizar todas as suas escolas
	api.GET("/sponsor", middleware.DriverMiddleware(), ct.GetSponsor) // para visualizar todos os sponsors
	api.GET(":cnh/school", func(c *gin.Context) {
		school := c.Query("school")
		if school == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "The param 'school' not found"})
		}
	}, ct.IsSchool)
}

func (ct *DriverController) GetSchool(c *gin.Context) {
}

func (ct *DriverController) GetSponsor(c *gin.Context) {

}

func (ct *DriverController) IsSchool(c *gin.Context) {

}
