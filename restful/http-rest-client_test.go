package restful

import (
	"testing"
)

func TestClient(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestClient"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Client()
		})
	}
}

func Test_getEmployeesClient(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test_getEmployeesClient"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getEmployeesClient()
		})
	}
}

func Test_postEmployeeClient(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test_postEmployeeClient"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postEmployeeClient()
		})
	}
}

func Test_updateEmployeeClient(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test_updateEmployeeClient"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateEmployeeClient()
		})
	}
}

func Test_deleteEmployeeClient(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test_deleteEmployeeClient"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteEmployeeClient()
		})
	}
}
