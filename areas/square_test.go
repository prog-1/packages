package areas

import "testing"

func TestSquare(t *testing.T) {
	type args struct {
		side float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"zero", args{0}, 0},
		{"negative", args{-1}, 0},
		{"arbitrary", args{11}, 121},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Square(tt.args.side); got != tt.want {
				t.Errorf("Square() = %v, want %v", got, tt.want)
			}
		})
	}
}
