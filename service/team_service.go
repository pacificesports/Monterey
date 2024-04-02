package service

import (
	"monterey/model"
	"monterey/utils"
)

func GetAllTeams() []model.Team {
	var teams []model.Team
	result := DB.Find(&teams)
	if result.Error != nil {
		utils.SugarLogger.Errorln(result.Error.Error())
	}
	for i := range teams {
		teams[i].BannerURL = GetBannerForTeam(teams[i].ID)
	}
	return teams
}

func GetTeamByID(teamID string) model.Team {
	var team model.Team
	result := DB.Where("id = ?", teamID).First(&team)
	if result.Error != nil {
		utils.SugarLogger.Errorln(result.Error.Error())
	}
	team.BannerURL = GetBannerForTeam(teamID)
	return team
}

func CreateTeam(team model.Team) error {
	if DB.Where("id = ?", team.ID).Select("*").Updates(&team).RowsAffected == 0 {
		utils.SugarLogger.Infoln("New team created with id: " + team.ID)
		if result := DB.Create(&team); result.Error != nil {
			return result.Error
		}
	} else {
		utils.SugarLogger.Infoln("Team with id: " + team.ID + " has been updated!")
	}
	return nil
}

func DeleteTeam(teamID string) error {
	if result := DB.Where("id = ?", teamID).Delete(&model.Team{}); result.Error != nil {
		return result.Error
	}
	return nil
}

func GetBannerForTeam(teamID string) string {
	var teamOrg model.TeamOrganization
	result := DB.Where("team_id = ?", teamID).First(&teamOrg)
	if result.Error != nil {
		utils.SugarLogger.Errorln(result.Error.Error())
	}
	if teamOrg.OrganizationID != "" {
		return GetOrganizationByID(teamOrg.OrganizationID).BannerURL
	}
	return ""
}
