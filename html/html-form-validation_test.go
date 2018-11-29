package html

import "testing"

func TestValidator(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestValidator"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Validator()
		})
	}
}
