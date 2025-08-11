package repository

import (
	pb "app/internal/genproto/users"
	"app/internal/models"
	"context"
	"reflect"
	"testing"
)

func TestNewUmsRepositoryLocal(t *testing.T) {
	tests := []struct {
		name string
		want *UmsRepositoryLocal
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUmsRepositoryLocal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUmsRepositoryLocal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUmsRepositoryLocal_Authenticate(t *testing.T) {
	type fields struct {
		users []models.User
	}
	type args struct {
		ctx         context.Context
		authRequest *pb.AuthenticateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.AuthenticateResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &UmsRepositoryLocal{
				users: tt.fields.users,
			}
			got, err := a.Authenticate(tt.args.ctx, tt.args.authRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Authenticate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUmsRepositoryLocal_GetUserById(t *testing.T) {
	type fields struct {
		users []models.User
	}
	type args struct {
		ctx         context.Context
		userRequest *pb.GetUserByIdRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.GetUserByIdResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &UmsRepositoryLocal{
				users: tt.fields.users,
			}
			got, err := a.GetUserById(tt.args.ctx, tt.args.userRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUmsRepositoryLocal_GetUsers(t *testing.T) {
	type fields struct {
		users []models.User
	}
	type args struct {
		ctx         context.Context
		userRequest *pb.GetUsersRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.GetUsersResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &UmsRepositoryLocal{
				users: tt.fields.users,
			}
			got, err := a.GetUsers(tt.args.ctx, tt.args.userRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsers() got = %v, want %v", got, tt.want)
			}
		})
	}
}
