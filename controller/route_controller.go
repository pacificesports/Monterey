package controller

import (
	"context"
	"log"
	"monterey/config"
	"monterey/service"
	"monterey/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	router.GET("/"+strings.ToLower(config.Service.Name)+"/ping", Ping)
	router.GET("/users/organizations/:userID", GetAllOrganizationsForUser)
	router.GET("/users/teams/:userID", GetAllTeamsForUser)

	router.GET("/organizations", GetAllOrganizations)
	router.GET("/organizations/:organizationID", GetOrganizationByID)
	router.POST("/organizations", CreateOrganization)
	router.GET("/organizations/:organizationID/teams", GetTeamsForOrganization)
	router.POST("/organizations/:organizationID/teams/:teamID", AddTeamToOrganization)
	router.DELETE("/organizations/:organizationID/teams/:teamID", RemoveTeamFromOrganization)
	router.GET("/organizations/:organizationID/users", GetAllUsersForOrganization)
	router.GET("/organizations/:organizationID/users/:userID", GetUserForOrganization)
	router.POST("/organizations/:organizationID/users", SetUserForOrganization)
	router.DELETE("/organizations/:organizationID/users/:userID", RemoveUserFromOrganization)
	router.GET("/organizations/:organizationID/users/:userID/roles", GetRolesForUserForOrganization)
	router.POST("/organizations/:organizationID/users/:userID/roles", SetRolesForUserForOrganization)

	router.GET("/teams", GetAllTeams)
	router.GET("/teams/:teamID", GetTeamByID)
	router.POST("/teams", CreateTeam)
	router.DELETE("/teams/:teamID", DeleteTeam)
	router.GET("/teams/:teamID/users", GetAllUsersForTeam)
	router.GET("/teams/:teamID/users/:userID", GetUserForTeam)
	router.POST("/teams/:teamID/users", SetUserForTeam)
	router.DELETE("/teams/:teamID/users/:userID", RemoveUserFromTeam)
	router.GET("/teams/:teamID/users/:userID/roles", GetRolesForUserForTeam)
	router.POST("/teams/:teamID/users/:userID/roles", SetRolesForUserForTeam)
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.SugarLogger.Infoln("GATEWAY REQUEST ID: " + c.GetHeader("Request-ID"))
		c.Next()
	}
}

func AuthChecker() gin.HandlerFunc {
	return func(c *gin.Context) {

		var requestUserID string

		ctx := context.Background()
		client, err := service.FirebaseAdmin.Auth(ctx)
		if err != nil {
			log.Fatalf("error getting Auth client: %v\n", err)
		}
		if c.GetHeader("Authorization") != "" {
			token, err := client.VerifyIDToken(ctx, strings.Split(c.GetHeader("Authorization"), "Bearer ")[1])
			if err != nil {
				utils.SugarLogger.Errorln("error verifying ID token")
				requestUserID = "null"
			} else {
				utils.SugarLogger.Infoln("Decoded User ID: " + token.UID)
				requestUserID = token.UID
			}
		} else {
			utils.SugarLogger.Infoln("No user token provided")
			requestUserID = "null"
		}
		c.Set("userID", requestUserID)
		// The main authentication gateway per request path
		// The requesting user's ID and roles are pulled and used below
		// Any path can also be quickly halted if not ready for prod
		c.Next()
	}
}

func contains(s []string, element string) bool {
	for _, i := range s {
		if i == element {
			return true
		}
	}
	return false
}
