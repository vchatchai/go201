package https

import "testing"

func TestHttps(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestHttps"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Https()
		})
	}
}
