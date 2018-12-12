package restful

import "testing"

func TestUpdate(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestUpdate"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Update()
		})
	}
}
