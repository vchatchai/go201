package mongo

import "testing"

func TestDelete(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestDelete"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Delete()
		})
	}
}
