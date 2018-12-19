package restful

import "testing"

func TestDeleteRest(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestDeleteRest"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteRest()
		})
	}
}
