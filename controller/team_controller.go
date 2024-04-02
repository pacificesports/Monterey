package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	if result.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "Failed to find team with id: " + c.Param("teamID")})
		return
	}
	c.JSON(http.StatusOK, result)
}

func CreateTeam(c *gin.Context) {
	var input model.Team
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if input.ID == "" {
		input.ID = uuid.New().String()
	}
	if result := service.CreateTeam(input); result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": result.Error()})
		return
	}
	c.JSON(http.StatusOK, service.GetTeamByID(input.ID))
}

func DeleteTeam(c *gin.Context) {
	result := service.DeleteTeam(c.Param("teamID"))
	if result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": result.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Team deleted successfully"})
}
