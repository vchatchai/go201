package html

import "testing"

func TestUploadFile(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestUploadFile"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UploadFile()
		})
	}
}
