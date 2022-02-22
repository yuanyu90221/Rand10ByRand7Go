package random_util

import "testing"

func TestRand10(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "check rand10 range",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rand10(); (got <= 10 && got >= 1) != tt.want {
				t.Errorf("Rand10() = %v , want %v", got, tt.want)
			}
		})
	}
}
