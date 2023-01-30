package roleMap

import (
	"deputy/data/smallRepository"
	"deputy/models"
	"reflect"
	"sort"
	"testing"
)

func TestRoleMap_InitRoleMapById(t *testing.T) {
	type args struct {
		roles []models.Role
	}
	tests := []struct {
		name        string
		args        args
		want_length int
		want_key_3  models.Role
	}{
		{
			"create roleMap and check role with id 3",
			args{
				smallRepository.Roles,
			},
			len(smallRepository.Roles),
			smallRepository.Roles[2],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RoleMap{}
			r.InitRoleMapById(tt.args.roles)
			if len(r.RoleMapById) != tt.want_length {
				t.Errorf("Role roleMap length got = %v, want %v", len(r.RoleMapById), tt.want_length)
			}
			if !reflect.DeepEqual(r.RoleMapById[3], tt.want_key_3) {
				t.Errorf("GetSubordinatesForRoleId() got = %v, want %v", r.RoleMapById[3], tt.want_key_3)
			}
		})
	}
}

func TestRoleMap_GetAllSubordinateIdsById(t *testing.T) {
	type args struct {
		roleId int
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
			[]int{2, 3, 4, 5},
		},
		{
			"Supervisor, id 3",
			args{
				3,
			},
			[]int{4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RoleMap{}
			r.InitRoleMapById(smallRepository.Roles)
			got := r.GetAllSubordinateIdsById(tt.args.roleId)
			// maps do not guarantee order
			sort.Slice(got, func(i, j int) bool {
				return got[i] < got[j]
			})
			if got := r.GetAllSubordinateIdsById(tt.args.roleId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllSubordinateIdsById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoleMap_GetRoleById(t *testing.T) {
	type args struct {
		roleId int
	}
	tests := []struct {
		name string
		args args
		want models.Role
	}{
		{
			"Employee, id 4",
			args{
				4,
			},
			smallRepository.Roles[3],
		},
		{
			"Location manager, id 2",
			args{
				2,
			},
			smallRepository.Roles[1],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RoleMap{}
			r.InitRoleMapById(smallRepository.Roles)
			if got := r.GetRoleById(tt.args.roleId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRoleById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoleMap_InitSubordinatesById(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			"Init from small repository",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RoleMap{}
			r.InitRoleMapById(smallRepository.Roles)
			if err := r.InitSubordinatesById(); (err != nil) != tt.wantErr {
				t.Errorf("InitSubordinatesById() error = %v, wantErr %v", err, tt.wantErr)
			}
			// Root id 0 has one child. Employee and trainee have no children
			if len(r.SubordinatesById) != 4 {
				t.Errorf("InitSubordinatesById() len = %v, want len %v", len(r.SubordinatesById), 4)
			}
			if len(r.SubordinatesById[1].Subordinates) != 1 {
				t.Errorf("InitSubordinatesById() subordinate [1] len = %v, want len %v", len(r.SubordinatesById[1].Subordinates), 1)
			}
			if len(r.SubordinatesById[3].Subordinates) != 2 {
				t.Errorf("InitSubordinatesById() subordinate [3] len = %v, want len %v", len(r.SubordinatesById[3].Subordinates), 2)
			}
		})
	}
}

func TestRoleMap_getIndirectSubordinateIdsById(t *testing.T) {
	type args struct {
		roleId int
	}
	tests := []struct {
		name             string
		args             args
		wantSubordiantes int
	}{
		{
			"System Administrator, id 1",
			args{
				1,
			},
			4,
		},
		{
			"Location Manager, id 2",
			args{
				2,
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RoleMap{}
			r.InitRoleMapById(smallRepository.Roles)
			r.InitSubordinatesById()
			r.GetAllSubordinateIdsById(tt.args.roleId)
			if len(r.SubordinatesById[tt.args.roleId].Subordinates) != tt.wantSubordiantes {
				t.Errorf("getIndirectSubordinateIdsById() len all subordinaes = %v, want len %v", len(r.SubordinatesById[tt.args.roleId].Subordinates), tt.wantSubordiantes)
			}
		})
	}
}
