package repository

import (
	"app/internal/interfaces"
	"app/internal/models"
	"context"
	"reflect"
	"testing"
)

func TestSpacecraftRepositoryMysql_Create(t *testing.T) {
	type fields struct {
		conn interfaces.DBStore
	}
	type args struct {
		ctx   context.Context
		craft *models.SpacecraftRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SpacecraftRepositoryMysql{
				conn: tt.fields.conn,
			}
			got, err := r.Create(tt.args.ctx, tt.args.craft)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpacecraftRepositoryMysql_Delete(t *testing.T) {
	type fields struct {
		conn interfaces.DBStore
	}
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SpacecraftRepositoryMysql{
				conn: tt.fields.conn,
			}
			if err := r.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSpacecraftRepositoryMysql_Get(t *testing.T) {
	type fields struct {
		conn interfaces.DBStore
	}
	type args struct {
		ctx context.Context
		in1 *map[string][]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Spacecraft
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SpacecraftRepositoryMysql{
				conn: tt.fields.conn,
			}
			got, err := r.Get(tt.args.ctx, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpacecraftRepositoryMysql_GetByID(t *testing.T) {
	type fields struct {
		conn interfaces.DBStore
	}
	type args struct {
		ctx context.Context
		id  int
		in2 *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Spacecraft
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SpacecraftRepositoryMysql{
				conn: tt.fields.conn,
			}
			got, err := r.GetByID(tt.args.ctx, tt.args.id, tt.args.in2)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpacecraftRepositoryMysql_Update(t *testing.T) {
	type fields struct {
		conn interfaces.DBStore
	}
	type args struct {
		ctx   context.Context
		id    string
		craft *models.SpacecraftRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SpacecraftRepositoryMysql{
				conn: tt.fields.conn,
			}
			if err := r.Update(tt.args.ctx, tt.args.id, tt.args.craft); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
