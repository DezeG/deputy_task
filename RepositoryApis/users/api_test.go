package users

import (
	"deputy/data/smallRepository"
	"deputy/models"
	"reflect"
	"testing"
)

func TestGetUserFromSimpleRepository(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    models.User
		wantErr bool
	}{
		{
			"Get Amily, id  = 2",
			args{
				2,
			},
			smallRepository.Users[1],
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserFromSimpleRepository(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserFromSimpleRepository() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserFromSimpleRepository() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUsersFromSimpleRepository(t *testing.T) {
	tests := []struct {
		name    string
		want    []models.User
		wantErr bool
	}{
		{
			"Get all default users",
			smallRepository.Users,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUsersFromSimpleRepository()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsersFromSimpleRepository() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsersFromSimpleRepository() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUsersByRoleIdsFromSimpleRepository(t *testing.T) {
	type args struct {
		roleIds []int
	}
	tests := []struct {
		name    string
		args    args
		want    []models.User
		wantErr bool
	}{
		{
			"Get users with role ids 3, 4 and 5",
			args{
				[]int{3, 4, 5},
			},
			[]models.User{smallRepository.Users[1], smallRepository.Users[2], smallRepository.Users[4]},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUsersByRoleIdsFromSimpleRepository(tt.args.roleIds)
			if (err != nil) != tt.wantErr {
				t.Errorf("getSubordinates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSubordinates() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRoleIdForUserId(t *testing.T) {
	type args struct {
		userId int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"User Id 1",
			args{
				1,
			},
			1,
			false,
		},
		{
			"User Id 4",
			args{
				4,
			},
			2,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRoleIdForUserId(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRoleIdForUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetRoleIdForUserId() got = %v, want %v", got, tt.want)
			}
		})
	}
}
