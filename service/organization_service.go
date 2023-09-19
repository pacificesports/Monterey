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
	result := DB.Where("organization_id = ?", organizationID).First(&organization)
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
