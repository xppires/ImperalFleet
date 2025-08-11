package repository

import (
	"app/internal/models"
	"reflect"
	"testing"
)

func TestAuthRepositoryLocal_Authenticate(t *testing.T) {
	type fields struct {
		users []models.User
	}
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		want1   string
		want2   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AuthRepositoryLocal{
				users: tt.fields.users,
			}
			got, got1, got2, err := a.Authenticate(tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Authenticate() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Authenticate() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("Authenticate() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestNewAuthRepositoryLocal(t *testing.T) {
	tests := []struct {
		name string
		want *AuthRepositoryLocal
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthRepositoryLocal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthRepositoryLocal() = %v, want %v", got, tt.want)
			}
		})
	}
}
