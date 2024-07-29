package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/vtx-account-manager/internal/server"
	"github.com/venture-technology/vtx-account-manager/internal/service"
	"github.com/venture-technology/vtx-account-manager/models"
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

	api := router.Group("vtx-account-manager/api/v1/driver")

	api.GET("/ping", ct.Ping)
	api.GET("/:cnh", ct.GetSchool)                 // para verificar se uma escola é parceira de um motorista
	api.POST("/partner", ct.CreatePartner)         // para criar uma parceria entre escola e motorista
	api.GET("/:cnh/school", ct.GetPartners)        // para visualizar todas as suas escolas
	api.GET("/:cnh/contract", ct.GetContract)      // para visualizar todos os contracts
	api.GET("/:cnh/shift", ct.GetContractsByShift) // para buscar todos os contracts de acordo com o horário da escola
}

func (ct *DriverController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}

func (ct *DriverController) CreatePartner(c *gin.Context) {

	var input models.Partner

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, server.InternalServerErrorResponse(err))
		return
	}

	err := ct.driverservice.CreatePartner(c, &input)

	if err != nil {
		log.Printf("error to create partners: %s", err.Error())
		c.JSON(http.StatusBadRequest, server.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, input)

}

func (ct *DriverController) GetPartners(c *gin.Context) {

	cnh := c.Param("cnh")

	partners, err := ct.driverservice.GetPartners(c, &cnh)

	if err != nil {
		log.Printf("error while found partners: %s", err.Error())
		c.JSON(http.StatusBadRequest, server.InternalServerErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, partners)

}

func (ct *DriverController) GetSchool(c *gin.Context) {

	cnpj := c.Query("school")

	if cnpj == "" {
		c.JSON(http.StatusBadRequest, server.NotParamErrorResponse("school"))
	}

	cnh := c.Param("cnh")

	partner, err := ct.driverservice.GetSchool(c, &cnh, &cnpj)
	if err != nil {
		log.Printf("error while found school: %s", err.Error())
		c.JSON(http.StatusBadRequest, server.NotFoundObjectErrorResponse("school"))
		return
	}

	c.JSON(http.StatusOK, partner)

}

func (ct *DriverController) GetContract(c *gin.Context) {

	cnh := c.Param("cnh")

	contracts, err := ct.driverservice.GetContracts(c, &cnh)
	if err != nil {
		log.Printf("error while found contracts: %s", err.Error())
		c.JSON(http.StatusBadRequest, server.NotFoundObjectErrorResponse("contract"))
		return
	}

	c.JSON(http.StatusOK, contracts)

}

func (ct *DriverController) GetContractsByShift(c *gin.Context) {

	shift := c.Query("shift")
	if shift == "" {
		c.JSON(http.StatusBadRequest, server.NotParamErrorResponse("shift"))
	}

	cnh := c.Param("cnh")

	contracts, err := ct.driverservice.GetContractsByShift(c, &cnh, &shift)
	if err != nil {
		log.Printf("error while found contracts: %s", err.Error())
		c.JSON(http.StatusBadRequest, server.NotFoundObjectErrorResponse("contract"))
		return
	}

	c.JSON(http.StatusOK, contracts)

}
