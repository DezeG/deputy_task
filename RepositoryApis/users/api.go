package users

import (
	"deputy/models"
	"errors"
)

func GetUserFromSimpleRepository(id int) (models.User, error) {
	users, err := getAllUsers()
	if err != nil {
		return models.User{}, err
	}
	for _, user := range users {
		if user.Id == id {
			return user, nil
		}
	}
	return models.User{}, errors.New("User is not found")
}

func GetUsersFromSimpleRepository() ([]models.User, error) {
	return getAllUsers()
}

func GetUsersByRoleIdsFromSimpleRepository(roleIds []int) ([]models.User, error) {
	users, err := GetUsersFromSimpleRepository()
	if err != nil {
		return []models.User{}, err
	}

	var subordinates []models.User
	for _, user := range users {
		if roleInRoles(roleIds, user.Role) {
			subordinates = append(subordinates, user)
		}
	}
	return subordinates, nil
}

func GetRoleIdForUserId(userId int) (int, error) {
	user, err := GetUserFromSimpleRepository(userId)
	if err != nil {
		return 0, err
	}
	return user.Role, nil
}
