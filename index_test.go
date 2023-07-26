package rbac

import "testing"

func TestRbac(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Add action",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rbac()
		})
	}
}
