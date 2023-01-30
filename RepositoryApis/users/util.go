package users

import (
	"deputy/data/smallRepository"
	"deputy/models"
)

func getAllUsers() ([]models.User, error) {
	return smallRepository.Users, nil
}

func roleInRoles(roleIds []int, id int) bool {
	for _, roleId := range roleIds {
		if roleId == id {
			return true
		}
	}
	return false
}
