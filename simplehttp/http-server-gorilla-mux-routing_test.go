package simplehttp

import "testing"

func TestGorilla(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestGorilla"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Gorilla()
		})
	}
}
