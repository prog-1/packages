package areas

import (
	"math"
	"testing"
)

func TestCircle(t *testing.T) {
	type args struct {
		r float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"zero", args{0}, 0},
		{"negative", args{-1}, 0},
		{"one", args{1}, math.Pi},
		{"arbitrary", args{12.5}, math.Pi * 12.5 * 12.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Circle(tt.args.r); got != tt.want {
				t.Errorf("Circle() = %v, want %v", got, tt.want)
			}
		})
	}
}
