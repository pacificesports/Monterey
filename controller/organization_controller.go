package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	if result.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "Failed to find organization with id: " + c.Param("organizationID")})
		return
	}
	c.JSON(http.StatusOK, result)
}

func CreateOrganization(c *gin.Context) {
	var input model.Organization
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if input.ID == "" {
		input.ID = uuid.New().String()
	}
	if result := service.CreateOrganization(input); result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": result.Error()})
		return
	}
	c.JSON(http.StatusOK, service.GetOrganizationByID(input.ID))
}

func GetTeamsForOrganization(c *gin.Context) {
	result := service.GetTeamsForOrganization(c.Param("organizationID"))
	c.JSON(http.StatusOK, result)
}

func AddTeamToOrganization(c *gin.Context) {
	if service.CheckTeamInOrganization(c.Param("organizationID"), c.Param("teamID")) {
		c.JSON(http.StatusConflict, gin.H{"message": "Team already in organization"})
		return
	}
	if err := service.AddTeamToOrganization(c.Param("organizationID"), c.Param("teamID")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Team added to organization"})
}

func RemoveTeamFromOrganization(c *gin.Context) {
	if !service.CheckTeamInOrganization(c.Param("organizationID"), c.Param("teamID")) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Team not in organization"})
		return
	}
	if err := service.RemoveTeamFromOrganization(c.Param("organizationID"), c.Param("teamID")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Team removed from organization"})
}
