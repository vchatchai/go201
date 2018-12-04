package cache

import "testing"

func TestHttpCache(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestHttpCache"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HttpCache()
		})
	}
}
