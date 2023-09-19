package service

import "monterey/model"

func GetRolesForUserForOrganization(organizationID string, userID string) []string {
	var roles []string
	DB.Table("organization_user_role").Where("organization_id = ? AND user_id = ?", organizationID, userID).Pluck("role", &roles)
	return roles
}

func SetRolesForUserForOrganization(organizationID string, userID string, roles []string) error {
	DB.Table("organization_user_role").Where("organization_id = ? AND user_id = ?", organizationID, userID).Delete(&model.OrganizationUserRole{})
	for _, role := range roles {
		if result := DB.Create(&model.OrganizationUserRole{
			OrganizationID: organizationID,
			UserID:         userID,
			Role:           role,
		}); result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func DeleteRolesForUserForOrganization(organizationID string, userID string) {
	DB.Table("organization_user_role").Where("organization_id = ? AND user_id = ?", organizationID, userID).Delete(&model.OrganizationUserRole{})
}

func GetRolesForUserForTeam(teamID string, userID string) []string {
	var roles []string
	DB.Table("team_user_role").Where("team_id = ? AND user_id = ?", teamID, userID).Pluck("role", &roles)
	return roles
}

func SetRolesForUserForTeam(teamID string, userID string, roles []string) error {
	DB.Table("team_user_role").Where("team_id = ? AND user_id = ?", teamID, userID).Delete(&model.TeamUserRole{})
	for _, role := range roles {
		if result := DB.Create(&model.TeamUserRole{
			TeamID: teamID,
			UserID: userID,
			Role:   role,
		}); result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func DeleteRolesForUserForTeam(teamID string, userID string) {
	DB.Table("team_user_role").Where("team_id = ? AND user_id = ?", teamID, userID).Delete(&model.TeamUserRole{})
}
