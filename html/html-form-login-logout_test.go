package html

import "testing"

func TestHttpForm(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestHttpForm"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HttpForm()
		})
	}
}
