package sortedArray

import (
	"deputy/data/smallRepository"
	"deputy/models"
	"reflect"
	"sort"
	"testing"
)

func TestRoleSortedArray_InitSortedById(t *testing.T) {
	type args struct {
		roles []models.Role
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"sorted source",
			args{
				smallRepository.Roles,
			},
		},
		{
			"not sorted source",
			args{
				append(smallRepository.Roles[2:], smallRepository.Roles[:2]...),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RoleSortedArray{}
			r.InitSortedById(tt.args.roles)
			if !sort.SliceIsSorted(r.SortedById, func(i, j int) bool { return r.SortedById[i].Id < r.SortedById[j].Id }) {
				t.Errorf("Strcut RoleSortedArray is not sorted")
			}

		})
	}
}

func TestRoleSortedArray_GetAllSubordinateIdsById(t *testing.T) {
	type args struct {
		roleId int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			"System Administrator, id - 1",
			args{
				1,
			},
			[]int{2, 3, 4, 5},
			false,
		},
		{
			"Supervisor, id - 3",
			args{
				3,
			},
			[]int{4, 5},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RoleSortedArray{}
			r.InitSortedById(smallRepository.Roles)
			got, err := r.GetAllSubordinateIdsById(tt.args.roleId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllSubordinateIdsById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllSubordinateIdsById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoleSortedArray_GetRoleById(t *testing.T) {
	type args struct {
		roleId int
	}
	tests := []struct {
		name    string
		args    args
		want    models.Role
		wantErr bool
	}{
		{
			"Trainee, id 5",
			args{
				5,
			},
			smallRepository.Roles[4],
			false,
		},
		{
			"Employee, id 4",
			args{
				4,
			},
			smallRepository.Roles[3],
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RoleSortedArray{}
			r.InitSortedById(smallRepository.Roles)
			got, err := r.GetRoleById(tt.args.roleId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRoleById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRoleById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoleSortedArray_InitSortedByParent(t *testing.T) {
	tests := []struct {
		name    string
		sourse  []models.Role
		want    []models.Role
		wantErr bool
	}{
		{
			"two elements",
			[]models.Role{
				{1, "A", 1000},
				{2, "B", 10},
			},
			[]models.Role{
				{2, "B", 10},
				{1, "A", 1000},
			},
			false,
		},
		{
			"three elements",
			[]models.Role{
				{1, "A", 1000},
				{2, "B", 10},
				{3, "C", 100},
			},
			[]models.Role{
				{2, "B", 10},
				{3, "C", 100},
				{1, "A", 1000},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RoleSortedArray{}
			r.InitSortedById(tt.sourse)
			r.InitSortedByParent()
			if err := r.InitSortedByParent(); (err != nil) != tt.wantErr {
				t.Errorf("InitSortedByParent() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(r.SortedByParent, tt.want) {
				t.Errorf("GetAllSubordinateIdsById() got = %v, want %v", r.SortedByParent, tt.want)
			}
		})
	}
}

func TestRoleSortedArray_getDirectSubordinateIdsById(t *testing.T) {
	type args struct {
		parentId int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"System Administrator, id 1",
			args{
				1,
			},
			[]int{2},
		},
		{
			"Location Manager, id 2",
			args{
				2,
			},
			[]int{3},
		},
		{
			"Supervisor, id 3",
			args{
				3,
			},
			[]int{4, 5},
		},
		{
			"trainee, id 4",
			args{
				4,
			},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RoleSortedArray{}
			r.InitSortedById(smallRepository.Roles)
			r.InitSortedByParent()
			if got := r.getDirectSubordinateIdsById(tt.args.parentId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDirectSubordinateIdsById() = %v, want %v", got, tt.want)
			}
		})
	}
}
