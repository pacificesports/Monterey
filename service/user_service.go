package service

import (
	"encoding/json"
	"io"
	"monterey/model"
	"monterey/utils"
	"net/http"
)

func SetUserForOrganization(user model.OrganizationUser) error {
	if DB.Where("organization_id = ? AND user_id = ?", user.OrganizationID, user.UserID).Updates(&user).RowsAffected == 0 {
		utils.SugarLogger.Infoln("New user with id: " + user.UserID + " added to organization with id: " + user.OrganizationID)
		if result := DB.Create(&user); result.Error != nil {
			return result.Error
		}
	} else {
		utils.SugarLogger.Infoln("User with id: " + user.UserID + " in organization with id: " + user.OrganizationID + " has been updated!")
	}
	return nil
}

func GetAllUsersForOrganization(organizationID string) []model.OrganizationUser {
	var users []model.OrganizationUser
	result := DB.Where("organization_id = ?", organizationID).Find(&users)
	if result.Error != nil {
		utils.SugarLogger.Errorln(result.Error.Error())
	}
	for i := range users {
		users[i].User = FetchUserDetails(users[i].UserID)
		users[i].Roles = GetRolesForUserForOrganization(users[i].OrganizationID, users[i].UserID)
	}
	return users
}

func GetUserForOrganization(organizationID string, userID string) model.OrganizationUser {
	var user model.OrganizationUser
	result := DB.Where("organization_id = ? AND user_id = ?", organizationID, userID).First(&user)
	if result.Error != nil {
		utils.SugarLogger.Errorln(result.Error.Error())
	}
	user.User = FetchUserDetails(user.UserID)
	user.Roles = GetRolesForUserForOrganization(user.OrganizationID, user.UserID)
	return user
}

func RemoveUserFromOrganization(organizationID string, userID string) error {
	result := DB.Where("organization_id = ? AND user_id = ?", organizationID, userID).Delete(&model.OrganizationUser{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllOrganizationsForUser(userID string) []model.Organization {
	var userOrgs []model.OrganizationUser
	organizations := make([]model.Organization, 0)
	result := DB.Where("user_id = ?", userID).Find(&userOrgs)
	if result.Error != nil {
		utils.SugarLogger.Errorln(result.Error.Error())
	}
	for _, org := range userOrgs {
		organizations = append(organizations, GetOrganizationByID(org.OrganizationID))
	}
	return organizations
}

func SetUserForTeam(user model.TeamUser) error {
	if DB.Where("team_id = ? AND user_id = ?", user.TeamID, user.UserID).Updates(&user).RowsAffected == 0 {
		utils.SugarLogger.Infoln("New user with id: " + user.UserID + " added to team with id: " + user.TeamID)
		if result := DB.Create(&user); result.Error != nil {
			return result.Error
		}
	} else {
		utils.SugarLogger.Infoln("User with id: " + user.UserID + " in team with id: " + user.TeamID + " has been updated!")
	}
	return nil
}

func GetAllUsersForTeam(teamID string) []model.TeamUser {
	var users []model.TeamUser
	result := DB.Where("team_id = ?", teamID).Find(&users)
	if result.Error != nil {
		utils.SugarLogger.Errorln(result.Error.Error())
	}
	for i := range users {
		users[i].User = FetchUserDetails(users[i].UserID)
		users[i].Roles = GetRolesForUserForTeam(users[i].TeamID, users[i].UserID)
	}
	return users
}

func GetUserForTeam(teamID string, userID string) model.TeamUser {
	var user model.TeamUser
	result := DB.Where("team_id = ? AND user_id = ?", teamID, userID).First(&user)
	if result.Error != nil {
		utils.SugarLogger.Errorln(result.Error.Error())
	}
	user.User = FetchUserDetails(user.UserID)
	user.Roles = GetRolesForUserForTeam(user.TeamID, user.UserID)
	return user
}

func RemoveUserFromTeam(teamID string, userID string) error {
	result := DB.Where("team_id = ? AND user_id = ?", teamID, userID).Delete(&model.TeamUser{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllTeamsForUser(userID string) []model.Team {
	var userTeams []model.TeamUser
	teams := make([]model.Team, 0)
	result := DB.Where("user_id = ?", userID).Find(&userTeams)
	if result.Error != nil {
		utils.SugarLogger.Errorln(result.Error.Error())
	}
	for _, t := range userTeams {
		teams = append(teams, GetTeamByID(t.TeamID))
	}
	return teams
}

func FetchUserDetails(userID string) json.RawMessage {
	var responseJson json.RawMessage = []byte("{}")
	mappedService := MatchRoute("users", "-")
	if mappedService.ID != 0 {
		proxyClient := &http.Client{}
		//proxyRequest, _ := http.NewRequest("GET", "http://localhost"+":"+strconv.Itoa(mappedService.Port)+"/schools/"+schoolID, nil) // Use this when not running in Docker
		proxyRequest, _ := http.NewRequest("GET", mappedService.URL+"/users/"+userID, nil)
		proxyRequest.Header.Set("Request-ID", "-")
		proxyResponse, err := proxyClient.Do(proxyRequest)
		if err != nil {
			utils.SugarLogger.Errorln("Failed to get user information from " + mappedService.Name + ": " + err.Error())
			return responseJson
		}
		defer proxyResponse.Body.Close()
		proxyResponseBodyBytes, _ := io.ReadAll(proxyResponse.Body)
		json.Unmarshal(proxyResponseBodyBytes, &responseJson)
	}
	return responseJson
}
