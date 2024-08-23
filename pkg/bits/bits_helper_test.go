package bits

import (
	"fmt"
	"testing"
)

func TestGetMask(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		args args
		want byte
	}{
		{args{0}, 0},
		{args{1}, 0b1},
		{args{2}, 0b11},
		{args{3}, 0b11},
		{args{4}, 0b111},
		{args{5}, 0b111},
		{args{8}, 0b1111},
		{args{10}, 0b1111},
		{args{15}, 0b1111},
		{args{16}, 0b11111},
		{args{20}, 0b11111},
		{args{32}, 0b111111},
		{args{33}, 0b111111},
		{args{63}, 0b111111},
		{args{64}, 0b1111111},
		{args{65}, 0b1111111},
		{args{120}, 0b1111111},
		{args{128}, 0b11111111},
		{args{130}, 0b11111111},
		{args{253}, 0b11111111},
		{args{254}, 0b11111111},
		{args{255}, 0b11111111},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("v=%d", tt.args.v), func(t *testing.T) {
			if got := GetMaskByValue(tt.args.v); got != tt.want {
				t.Errorf("GetMaskByValue() = %b, want %b", got, tt.want)
			}
		})
	}
}
