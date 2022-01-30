package areas

import "testing"

func TestRectangle(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"zero", args{0, 0}, 0},
		{"negative a", args{-1, 0}, 0},
		{"negative b", args{0, -1}, 0},
		{"negative", args{-1, -1}, 0},
		{"a zero", args{0, 1}, 0},
		{"b zero", args{1, 0}, 0},
		{"arbitrary", args{11, 12}, 132},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rectangle(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Rectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
