package controller

import (
	"github.com/gin-gonic/gin"
	"monterey/model"
	"monterey/service"
	"net/http"
)

func GetAllOrganizations(c *gin.Context) {
	result := service.GetAllOrganizations()
	c.JSON(http.StatusOK, result)
}

func GetOrganizationByID(c *gin.Context) {
	result := service.GetOrganizationByID(c.Param("organizationID"))
	c.JSON(http.StatusOK, result)
}

func CreateOrganization(c *gin.Context) {
	var input model.Organization
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if result := service.CreateOrganization(input); result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": result.Error()})
		return
	}
	c.JSON(http.StatusOK, service.GetOrganizationByID(input.ID))
}
