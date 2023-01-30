package linear

import (
	"deputy/RepositoryApis/users"
	"deputy/models"
	"errors"
)

func GetRoleFromSimpleRepository(id int) (models.Role, error) {
	roles, err := getAllRoles()
	if err != nil {
		return models.Role{}, err
	}
	for _, role := range roles {
		if role.Id == id {
			return role, nil
		}
	}
	return models.Role{}, errors.New("Role is not found")
}

func GetRolesFromSimpleRepository() ([]models.Role, error) {
	return getAllRoles()
}

func GetSubordinatesForRoleId(RoleId int) ([]models.User, error) {
	subordinateRoleIds, err := getSubordinateRoleIdsForRoleId(RoleId)
	if err != nil {
		return nil, err
	}
	subordinates, err := users.GetUsersByRoleIdsFromSimpleRepository(subordinateRoleIds)
	if err != nil {
		return nil, err
	}
	return subordinates, nil
}
