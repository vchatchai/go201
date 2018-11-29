package html

import "testing"

func TestHtmlForm(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestHtmlForm"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HtmlForm()
		})
	}
}
