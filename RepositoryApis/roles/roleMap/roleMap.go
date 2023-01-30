package roleMap

import (
	"deputy/models"
	"errors"
)

type RoleMap struct {
	RoleMapById      map[int]models.Role
	SubordinatesById map[int]*SubordinateField
}

type SubordinateField struct {
	Subordinates     []int
	IndirectIncluded bool
}

func (r *RoleMap) InitRoleMapById(roles []models.Role) {
	r.RoleMapById = make(map[int]models.Role)
	for _, role := range roles {
		r.RoleMapById[role.Id] = role
	}
}

func (r *RoleMap) InitSubordinatesById() error {
	if r.RoleMapById == nil {
		return errors.New("Call function InitRoleMapById() first")
	}
	r.SubordinatesById = make(map[int]*SubordinateField)
	for _, role := range r.RoleMapById {
		subordinates, ok := r.SubordinatesById[role.Parent]
		if !ok {
			r.SubordinatesById[role.Parent] = &SubordinateField{
				Subordinates:     []int{role.Id},
				IndirectIncluded: false,
			}
			continue
		}
		(*subordinates).Subordinates = append((*subordinates).Subordinates, role.Id)
	}

	return nil
}

func (r *RoleMap) GetRoleById(roleId int) models.Role {
	return r.RoleMapById[roleId]
}

func (r *RoleMap) GetAllSubordinateIdsById(roleId int) []int {
	if r.SubordinatesById == nil {
		r.InitSubordinatesById()
	}
	subordinates, ok := r.SubordinatesById[roleId]
	if !ok {
		return []int{}
	} else if subordinates.IndirectIncluded {
		return subordinates.Subordinates
	}
	r.getIndirectSubordinateIdsById(roleId)

	return r.SubordinatesById[roleId].Subordinates
}

func (r *RoleMap) getIndirectSubordinateIdsById(roleId int) {
	for i := 0; i < len(r.SubordinatesById[roleId].Subordinates); i++ {
		subordinateRoleId := r.SubordinatesById[roleId].Subordinates[i]
		newSubordinates, ok := r.SubordinatesById[subordinateRoleId]
		if !ok {
			continue
		}
		currentSubordinates := r.SubordinatesById[roleId]
		(*currentSubordinates).Subordinates = append((*currentSubordinates).Subordinates, (*newSubordinates).Subordinates...)
	}

	r.SubordinatesById[roleId].IndirectIncluded = true
}
