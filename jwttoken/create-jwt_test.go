package jwttoken

import "testing"

func TestCreateJWt(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestCreateJWt"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateJWt()
		})
	}
}
