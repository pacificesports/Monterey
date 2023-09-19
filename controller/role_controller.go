package controller

import (
	"github.com/gin-gonic/gin"
	"monterey/service"
	"net/http"
)

func GetRolesForUserForOrganization(c *gin.Context) {
	result := service.GetRolesForUserForOrganization(c.Param("organizationID"), c.Param("userID"))
	c.JSON(http.StatusOK, result)
}

func SetRolesForUserForOrganization(c *gin.Context) {
	var input []string
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	service.DeleteRolesForUserForOrganization(c.Param("organizationID"), c.Param("userID"))
	if err := service.SetRolesForUserForOrganization(c.Param("organizationID"), c.Param("userID"), input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, service.GetRolesForUserForOrganization(c.Param("organizationID"), c.Param("userID")))
}

func GetRolesForUserForTeam(c *gin.Context) {
	result := service.GetRolesForUserForTeam(c.Param("teamID"), c.Param("userID"))
	c.JSON(http.StatusOK, result)
}

func SetRolesForUserForTeam(c *gin.Context) {
	var input []string
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	service.DeleteRolesForUserForTeam(c.Param("teamID"), c.Param("userID"))
	if err := service.SetRolesForUserForTeam(c.Param("teamID"), c.Param("userID"), input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, service.GetRolesForUserForTeam(c.Param("teamID"), c.Param("userID")))
}
