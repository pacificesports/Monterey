package controller

import (
	"github.com/gin-gonic/gin"
	"monterey/model"
	"monterey/service"
	"net/http"
)

func GetAllUsersForOrganization(c *gin.Context) {
	result := service.GetAllUsersForOrganization(c.Param("organizationID"))
	c.JSON(http.StatusOK, result)
}

func GetUserForOrganization(c *gin.Context) {
	result := service.GetUserForOrganization(c.Param("organizationID"), c.Param("userID"))
	c.JSON(http.StatusOK, result)
}

func SetUserForOrganization(c *gin.Context) {
	var input model.OrganizationUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := service.SetUserForOrganization(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, service.GetAllUsersForOrganization(c.Param("organizationID")))
}

func RemoveUserFromOrganization(c *gin.Context) {
	if err := service.RemoveUserFromOrganization(c.Param("organizationID"), c.Param("userID")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User removed from organization"})
}

func GetAllOrganizationsForUser(c *gin.Context) {
	result := service.GetAllOrganizationsForUser(c.Param("userID"))
	c.JSON(http.StatusOK, result)
}

func GetAllUsersForTeam(c *gin.Context) {
	result := service.GetAllUsersForTeam(c.Param("teamID"))
	c.JSON(http.StatusOK, result)
}

func GetUserForTeam(c *gin.Context) {
	result := service.GetUserForTeam(c.Param("teamID"), c.Param("userID"))
	c.JSON(http.StatusOK, result)
}

func SetUserForTeam(c *gin.Context) {
	var input model.TeamUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := service.SetUserForTeam(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, service.GetAllUsersForTeam(c.Param("teamID")))
}

func RemoveUserFromTeam(c *gin.Context) {
	if err := service.RemoveUserFromTeam(c.Param("teamID"), c.Param("userID")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User removed from team"})
}

func GetAllTeamsForUser(c *gin.Context) {
	result := service.GetAllTeamsForUser(c.Param("userID"))
	c.JSON(http.StatusOK, result)
}
