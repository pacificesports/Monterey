package controller

import (
	"github.com/gin-gonic/gin"
	"monterey/model"
	"monterey/service"
	"net/http"
)

func GetAllTeams(c *gin.Context) {
	result := service.GetAllTeams()
	c.JSON(http.StatusOK, result)
}

func GetTeamByID(c *gin.Context) {
	result := service.GetTeamByID(c.Param("teamID"))
	c.JSON(http.StatusOK, result)
}

func CreateTeam(c *gin.Context) {
	var input model.Team
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if result := service.CreateTeam(input); result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": result.Error()})
		return
	}
	c.JSON(http.StatusOK, service.GetTeamByID(input.ID))
}
