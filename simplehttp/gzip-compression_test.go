package simplehttp

import "testing"

func TestGZipCompression(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestGZipCompression"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GZipCompression()
		})
	}
}
