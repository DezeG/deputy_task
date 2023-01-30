package linear

import (
	"deputy/data/smallRepository"
	"deputy/models"
	"reflect"
	"testing"
)

func TestGetRoleFromSimpleRepository(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    models.Role
		wantErr bool
	}{
		{
			"Get Location Manager, id = 2",
			args{
				2,
			},
			models.Role{
				2,
				"Location Manager",
				1,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRoleFromSimpleRepository(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRoleFromSimpleRepository() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRoleFromSimpleRepository() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRolesFromSimpleRepository(t *testing.T) {
	tests := []struct {
		name    string
		want    []models.Role
		wantErr bool
	}{
		{
			"Get all default roles",
			smallRepository.Roles,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRolesFromSimpleRepository()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRolesFromSimpleRepository() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRolesFromSimpleRepository() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSubordinatesForRoleId(t *testing.T) {
	type args struct {
		RoleId int
	}
	tests := []struct {
		name    string
		args    args
		want    []models.User
		wantErr bool
	}{
		{
			"Get subordinated for a System Administrator, id = 1",
			args{
				1,
			},
			smallRepository.Users[1:],
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSubordinatesForRoleId(tt.args.RoleId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSubordinatesForRoleId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSubordinatesForRoleId() got = %v, want %v", got, tt.want)
			}
		})
	}
}
