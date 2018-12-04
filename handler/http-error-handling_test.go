package handler

import "testing"

func TestHttpErrorHandling(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestHttpErrorHandling"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HttpErrorHandling()
		})
	}
}
