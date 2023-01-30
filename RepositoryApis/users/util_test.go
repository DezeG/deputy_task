package users

import (
	"deputy/data/smallRepository"
	"deputy/models"
	"reflect"
	"testing"
)

func Test_getAllUsers(t *testing.T) {
	tests := []struct {
		name    string
		want    []models.User
		wantErr bool
	}{
		{
			"All default users",
			smallRepository.Users,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getAllUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("getAllUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllUsers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_roleInRoles(t *testing.T) {
	type args struct {
		roleIds []int
		id      int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Role 7 in the array of roles 1, 2, 3",
			args{
				[]int{1, 2, 3},
				7,
			},
			false,
		},
		{
			"Role 2 in the array of roles 1, 2, 3",
			args{
				[]int{1, 2, 3},
				2,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := roleInRoles(tt.args.roleIds, tt.args.id); got != tt.want {
				t.Errorf("roleInRoles() = %v, want %v", got, tt.want)
			}
		})
	}
}
