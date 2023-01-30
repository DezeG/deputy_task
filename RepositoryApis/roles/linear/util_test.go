package linear

import (
	"deputy/data/smallRepository"
	"deputy/models"
	"reflect"
	"testing"
)

func Test_getAllRoles(t *testing.T) {
	tests := []struct {
		name    string
		want    []models.Role
		wantErr bool
	}{
		{
			"Default five roles",
			smallRepository.Roles,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getAllRoles()
			if (err != nil) != tt.wantErr {
				t.Errorf("getAllRoles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllRoles() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAllSubordinateIdsForRoleId(t *testing.T) {
	type args struct {
		roles  *[]models.Role
		roleId int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Role ID 3",
			args{
				&smallRepository.Roles,
				3,
			},
			[]int{4, 5},
		},
		{
			"Role ID 1",
			args{
				&smallRepository.Roles,
				1,
			},
			[]int{2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAllSubordinateIdsForRoleId(tt.args.roles, tt.args.roleId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllSubordinateIdsForRoleId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDirectSubordinateIdsForRoleId(t *testing.T) {
	type args struct {
		roles  *[]models.Role
		roleId int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Role ID 3 Supervisor",
			args{
				&smallRepository.Roles,
				3,
			},
			[]int{4, 5},
		},
		{
			"Role ID 1 System Administrator",
			args{
				&smallRepository.Roles,
				1,
			},
			[]int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDirectSubordinateIdsForRoleId(tt.args.roles, tt.args.roleId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDirectSubordinateIdsForRoleId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getIndirectSubordinateIds(t *testing.T) {
	type args struct {
		roles              *[]models.Role
		subordinateRoleIds *[]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"All subordinates for Location Manager, id 2",
			args{
				&smallRepository.Roles,
				&[]int{2},
			},
			[]int{2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if getIndirectSubordinateIds(tt.args.roles, tt.args.subordinateRoleIds); !reflect.DeepEqual(*tt.args.subordinateRoleIds, tt.want) {
				t.Errorf("getIndirectSubordinateIds() = %v, want %v", tt.args.subordinateRoleIds, tt.want)
			}
		})
	}
}

func Test_getSubordinateRoleIdsForRoleId(t *testing.T) {
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
			"Subordinates for Supervisor, id = 3",
			args{
				3,
			},
			[]int{4, 5},
			false,
		},
		{
			"Subordinates for Location Manager, id = 2",
			args{
				2,
			},
			[]int{3, 4, 5},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getSubordinateRoleIdsForRoleId(tt.args.roleId)
			if (err != nil) != tt.wantErr {
				t.Errorf("getSubordinateRoleIdsForRoleId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSubordinateRoleIdsForRoleId() got = %v, want %v", got, tt.want)
			}
		})
	}
}
