package jwttoken

import "testing"

func TestJWT(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestJWT"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			JWT()
		})
	}
}
