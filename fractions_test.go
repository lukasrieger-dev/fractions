package fractions

import (
	"testing"
)

func TestNew(t *testing.T) {
	want := fraction{3, 5}
	got := New(3, 5)

	if !EqualFractions(got, want) {
		t.Errorf("New() = %v, want %v", got, want)
	}
}

func TestDiv(t *testing.T) {
	type args struct {
		fract1 fraction
		fract2 fraction
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "division1",
			args: args{
				fract1: fraction{1, 2},
				fract2: fraction{5, 2},
			},
			want: 0.2,
		},
		{
			name: "division2",
			args: args{
				fract1: fraction{5, 1},
				fract2: fraction{2, 2},
			},
			want: 5.0,
		},
		{
			name: "division3",
			args: args{
				fract1: fraction{0, 1},
				fract2: fraction{3, 2},
			},
			want: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.fract1.Div(tt.args.fract2).toFloat()

			if got != tt.want {
				t.Errorf("Div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFraction_toFloat(t *testing.T) {
	type fields struct {
		numerator   int
		denominator int
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"toFloat1", fields{1, 2}, 0.5},
		{"toFloat2", fields{3, -4}, -0.75},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fraction{
				numerator:   tt.fields.numerator,
				denominator: tt.fields.denominator,
			}
			if got := f.toFloat(); got != tt.want {
				t.Errorf("Fraction.toFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMult(t *testing.T) {
	type args struct {
		fract1 fraction
		fract2 fraction
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "mult1",
			args: args{
				fract1: fraction{1, 2},
				fract2: fraction{5, 2},
			},
			want: 1.25,
		},
		{
			name: "mult2",
			args: args{
				fract1: fraction{5, 1},
				fract2: fraction{2, -1},
			},
			want: -10.0,
		},
		{
			name: "mult3",
			args: args{
				fract1: fraction{0, 1},
				fract2: fraction{3, 2},
			},
			want: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.fract1.Mult(tt.args.fract2).toFloat()

			if got != tt.want {
				t.Errorf("Mult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFraction_normalize(t *testing.T) {
	type fields struct {
		numerator   int
		denominator int
	}
	type result struct {
		numerator   int
		denominator int
	}
	tests := []struct {
		name   string
		fields fields
		want   result
	}{
		{"normalize1", fields{10, 2}, result{5, 1}},
		{"normalize2", fields{11, 7}, result{11, 7}},
		{"normalize3", fields{48, 24}, result{2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fraction{
				numerator:   tt.fields.numerator,
				denominator: tt.fields.denominator,
			}
			f.Normalize()

			if !(f.numerator == tt.want.numerator && f.denominator == tt.want.denominator) {
				t.Errorf("Normalize() = %v, want %v", []int{f.numerator, f.denominator}, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		fract1 fraction
		fract2 fraction
	}
	tests := []struct {
		name string
		args args
		want fraction
	}{
		{"add1", args{fraction{10, 3}, fraction{5, 4}}, fraction{55, 12}},
		{"add2", args{fraction{7, 5}, fraction{9, 10}}, fraction{23, 10}},
		{"add3", args{fraction{-1, 3}, fraction{2, 6}}, fraction{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.fract1.Add(tt.args.fract2); !(got.toFloat() == tt.want.toFloat()) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	type args struct {
		fract1 fraction
		fract2 fraction
	}
	tests := []struct {
		name string
		args args
		want fraction
	}{
		{"sub1", args{fraction{10, 3}, fraction{5, 4}}, fraction{25, 12}},
		{"sub2", args{fraction{7, 5}, fraction{9, 10}}, fraction{5, 10}},
		{"sub3", args{fraction{-1, 3}, fraction{2, 6}}, fraction{-2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.fract1.Sub(tt.args.fract2); !(got.toFloat() == tt.want.toFloat()) {
				t.Errorf("Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}
