package linear

import (
	"deputy/data/smallRepository"
	"deputy/models"
)

func getSubordinateRoleIdsForRoleId(roleId int) ([]int, error) {
	role, err := GetRoleFromSimpleRepository(roleId)
	if err != nil {
		return []int{}, err
	}
	roles, err := GetRolesFromSimpleRepository()
	if err != nil {
		return nil, err
	}

	subRoleIds := getAllSubordinateIdsForRoleId(&roles, role.Id)

	return subRoleIds, nil
}

func getAllSubordinateIdsForRoleId(roles *[]models.Role, roleId int) []int {
	subRoleIds := getDirectSubordinateIdsForRoleId(roles, roleId)
	getIndirectSubordinateIds(roles, &subRoleIds)
	return subRoleIds
}

func getDirectSubordinateIdsForRoleId(roles *[]models.Role, roleId int) []int {
	var subordinateRoleIds []int
	for _, r := range *roles {
		if r.Parent == roleId {
			subordinateRoleIds = append(subordinateRoleIds, r.Id)
		}
	}
	return subordinateRoleIds
}

func getIndirectSubordinateIds(roles *[]models.Role, subordinateRoleIds *[]int) {
	for i := 0; i < len(*subordinateRoleIds); i++ {
		id := (*subordinateRoleIds)[i]
		for _, r := range *roles {
			if r.Parent == id {
				*subordinateRoleIds = append(*subordinateRoleIds, r.Id)
			}
		}
	}
}

func getAllRoles() ([]models.Role, error) {
	return smallRepository.Roles, nil
}
