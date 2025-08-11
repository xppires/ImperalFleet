package handlers

import (
	"app/internal/services"
	"github.com/gorilla/mux"
	"net/http"
	"reflect"
	"testing"
)

func TestNewSpacecrafHandlers(t *testing.T) {
	type args struct {
		spacecraftService *services.SpacecraftService
	}
	tests := []struct {
		name string
		args args
		want *SpacecraftHandlers
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSpacecrafHandlers(tt.args.spacecraftService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSpacecrafHandlers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpacecraftHandlers_Create(t *testing.T) {
	type fields struct {
		spacecraftService *services.SpacecraftService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SpacecraftHandlers{
				spacecraftService: tt.fields.spacecraftService,
			}
			s.Create(tt.args.w, tt.args.r)
		})
	}
}

func TestSpacecraftHandlers_Delete(t *testing.T) {
	type fields struct {
		spacecraftService *services.SpacecraftService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SpacecraftHandlers{
				spacecraftService: tt.fields.spacecraftService,
			}
			s.Delete(tt.args.w, tt.args.r)
		})
	}
}

func TestSpacecraftHandlers_GetByID(t *testing.T) {
	type fields struct {
		spacecraftService *services.SpacecraftService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SpacecraftHandlers{
				spacecraftService: tt.fields.spacecraftService,
			}
			s.GetByID(tt.args.w, tt.args.r)
		})
	}
}

func TestSpacecraftHandlers_RegisteRoutes(t *testing.T) {
	type fields struct {
		spacecraftService *services.SpacecraftService
	}
	type args struct {
		router *mux.Router
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SpacecraftHandlers{
				spacecraftService: tt.fields.spacecraftService,
			}
			s.RegisteRoutes(tt.args.router)
		})
	}
}

func TestSpacecraftHandlers_Update(t *testing.T) {
	type fields struct {
		spacecraftService *services.SpacecraftService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SpacecraftHandlers{
				spacecraftService: tt.fields.spacecraftService,
			}
			s.Update(tt.args.w, tt.args.r)
		})
	}
}

func TestSpacecraftHandlers_get(t *testing.T) {
	type fields struct {
		spacecraftService *services.SpacecraftService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SpacecraftHandlers{
				spacecraftService: tt.fields.spacecraftService,
			}
			s.get(tt.args.w, tt.args.r)
		})
	}
}
