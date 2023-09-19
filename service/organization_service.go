package service

import (
	"monterey/model"
	"monterey/utils"
)

func GetAllOrganizations() []model.Organization {
	var organizations []model.Organization
	result := DB.Find(&organizations)
	if result.Error != nil {
		utils.SugarLogger.Errorln(result.Error.Error())
	}
	return organizations
}

func GetOrganizationByID(organizationID string) model.Organization {
	var organization model.Organization
	result := DB.Where("id = ?", organizationID).First(&organization)
	if result.Error != nil {
		utils.SugarLogger.Errorln(result.Error.Error())
	}
	return organization
}

func CreateOrganization(organization model.Organization) error {
	if DB.Where("id = ?", organization.ID).Select("*").Updates(&organization).RowsAffected == 0 {
		utils.SugarLogger.Infoln("New organization created with id: " + organization.ID)
		if result := DB.Create(&organization); result.Error != nil {
			return result.Error
		}
	} else {
		utils.SugarLogger.Infoln("Organization with id: " + organization.ID + " has been updated!")
	}
	return nil
}

func GetTeamsForOrganization(organizationID string) []model.Team {
	var orgTeams []model.TeamOrganization
	teams := make([]model.Team, 0)
	result := DB.Where("organization_id = ?", organizationID).Find(&orgTeams)
	if result.Error != nil {
		utils.SugarLogger.Errorln(result.Error.Error())
	}
	for _, orgTeam := range orgTeams {
		teams = append(teams, GetTeamByID(orgTeam.TeamID))
	}
	return teams
}

func CheckTeamInOrganization(organizationID string, teamID string) bool {
	var orgTeam model.TeamOrganization
	result := DB.Where("organization_id = ? AND team_id = ?", organizationID, teamID).First(&orgTeam)
	if result.Error != nil {
		utils.SugarLogger.Errorln(result.Error.Error())
	}
	return orgTeam.OrganizationID != ""
}

func AddTeamToOrganization(organizationID string, teamID string) error {
	if result := DB.Create(&model.TeamOrganization{
		OrganizationID: organizationID,
		TeamID:         teamID,
	}); result.Error != nil {
		return result.Error
	}
	return nil
}

func RemoveTeamFromOrganization(organizationID string, teamID string) error {
	if result := DB.Where("organization_id = ? AND team_id = ?", organizationID, teamID).Delete(&model.TeamOrganization{}); result.Error != nil {
		return result.Error
	}
	return nil
}
