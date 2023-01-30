package sortedArray

import (
	"deputy/models"
	"errors"
	"sort"
)

type RoleSortedArray struct {
	SortedById     []models.Role
	SortedByParent []models.Role
}

func (r *RoleSortedArray) InitSortedById(roles []models.Role) {
	if sort.SliceIsSorted(roles, func(i, j int) bool { return roles[i].Id < roles[j].Id }) {
		r.SortedById = roles
		return
	}

	sort.Slice(roles, func(i, j int) bool {
		return roles[i].Id < roles[j].Id
	})

	r.SortedById = roles
}

func (r *RoleSortedArray) InitSortedByParent() error {
	if r.SortedById == nil {
		return errors.New("Call function InitRoleMapById() first")
	}
	r.SortedByParent = make([]models.Role, len(r.SortedById))
	numberOfCopiedElements := copy(r.SortedByParent, r.SortedById)
	if numberOfCopiedElements != len(r.SortedById) {
		return errors.New("Failed to copy slice")
	}
	sort.Slice(r.SortedByParent, func(i, j int) bool {
		return r.SortedByParent[i].Parent < r.SortedByParent[j].Parent
	})

	return nil
}

func (r *RoleSortedArray) GetRoleById(roleId int) (models.Role, error) {
	i := sort.Search(len(r.SortedById), func(i int) bool {
		return r.SortedById[i].Id >= roleId
	})

	if i >= len(r.SortedById) || r.SortedById[i].Id != roleId {
		return models.Role{}, errors.New("Cannot find role")
	}

	return r.SortedById[i], nil
}

func (r *RoleSortedArray) GetAllSubordinateIdsById(roleId int) ([]int, error) {
	if r.SortedByParent == nil {
		r.InitSortedByParent()
	}

	subordinates := r.getDirectSubordinateIdsById(roleId)
	for i := 0; i < len(subordinates); i++ {
		newSubordinates := r.getDirectSubordinateIdsById(subordinates[i])
		subordinates = append(subordinates, newSubordinates...)
	}

	return subordinates, nil
}

func (r *RoleSortedArray) getDirectSubordinateIdsById(parentId int) []int {
	index := sort.Search(len(r.SortedById), func(i int) bool {
		return r.SortedById[i].Parent >= parentId
	})

	if index >= len(r.SortedById) || r.SortedById[index].Parent != parentId {
		return []int{}
	}

	var start, end, k int
	for k = index - 1; k >= 0; k-- {
		if r.SortedByParent[k].Parent != r.SortedByParent[index].Parent {
			start = k + 1
			break
		}
	}
	if k == -1 {
		start = 0
	}
	for k = index + 1; k < len(r.SortedByParent); k++ {
		if r.SortedByParent[k].Parent != r.SortedByParent[index].Parent {
			end = k - 1
			break
		}
	}
	if k == len(r.SortedByParent) {
		end = len(r.SortedByParent) - 1
	}

	var subordinates []int
	for i := start; i <= end; i++ {
		subordinates = append(subordinates, r.SortedByParent[i].Id)
	}
	return subordinates
}
