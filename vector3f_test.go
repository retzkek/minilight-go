package main

import (
	"math"
	"testing"
)

const (
	EPS = 1e-10 // criterion for comparing reals
)

func AreAlmostEqual(x, y float64) bool {
	return math.Abs(x-y) < EPS
}

func VectorsAreAlmostEqual(v, u Vector3f) bool {
	for i := 0; i < 3; i++ {
		if !AreAlmostEqual(v[i], u[i]) {
			return false
		}
	}
	return true
}

var unarytests = []struct {
	v        Vector3f
	negative Vector3f
	unitized Vector3f
	isZero   bool
}{
	{Vector3fZERO,
		Vector3fZERO,
		Vector3fZERO,
		true},
	{Vector3fONE,
		Vector3f{-1.0, -1.0, -1.0},
		Vector3f{0.57735026919, 0.57735026919, 0.57735026919},
		false},
	{Vector3f{1.0, 2.0, 3.0},
		Vector3f{-1.0, -2.0, -3.0},
		Vector3f{0.267261241912, 0.534522483825, 0.801783725737},
		false},
}

func TestNegative(t *testing.T) {
	for _, tt := range unarytests {
		if x := tt.v.Negative(); !VectorsAreAlmostEqual(x, tt.negative) {
			t.Errorf("Negative(%v) = %v, want %v", tt.v, x, tt.negative)
		}
	}
}
func TestUnitized(t *testing.T) {
	for _, tt := range unarytests {
		if x := tt.v.Unitized(); !VectorsAreAlmostEqual(x, tt.unitized) {
			t.Errorf("Unitized(%v) = %v, want %v", tt.v, x, tt.unitized)
		}
	}
}
func TestIsZero(t *testing.T) {
	for _, tt := range unarytests {
		if x := tt.v.IsZero(); x != tt.isZero {
			t.Errorf("IsZero(%v) = %v, want %v", tt.v, x, tt.isZero)
		}
	}
}

var binarytests = []struct {
	v, u  Vector3f
	dot   float64
	cross Vector3f
	add   Vector3f
	sub   Vector3f
	mulv  Vector3f
}{
	{Vector3fZERO, Vector3fZERO,
		0.0,
		Vector3fZERO,
		Vector3fZERO,
		Vector3fZERO,
		Vector3fZERO},
	{Vector3fZERO, Vector3fONE,
		0.0,
		Vector3fZERO,
		Vector3fONE,
		Vector3f{-1.0, -1.0, -1.0},
		Vector3fZERO},
	{Vector3fONE, Vector3fONE,
		3.0,
		Vector3fZERO,
		Vector3f{2.0, 2.0, 2.0},
		Vector3fZERO,
		Vector3fONE},
	{Vector3f{1.0, 2.0, 3.0}, Vector3f{4.2, 5.1, -2.3},
		7.5,
		Vector3f{-19.9, 14.9, -3.3},
		Vector3f{5.2, 7.1, 0.7},
		Vector3f{-3.2, -3.1, 5.3},
		Vector3f{4.2, 10.2, -6.9}},
}

func TestDot(t *testing.T) {
	for _, tt := range binarytests {
		if x := tt.v.Dot(&tt.u); !AreAlmostEqual(x, tt.dot) {
			t.Errorf("%v Dot %v = %v, want %v", tt.v, tt.u, x, tt.dot)
		}
	}
}
func TestCross(t *testing.T) {
	for _, tt := range binarytests {
		if x := tt.v.Cross(&tt.u); !VectorsAreAlmostEqual(x, tt.cross) {
			t.Errorf("%v Cross %v = %v, want %v", tt.v, tt.u, x, tt.cross)
		}
	}
}
func TestAdd(t *testing.T) {
	for _, tt := range binarytests {
		if x := tt.v.Add(&tt.u); !VectorsAreAlmostEqual(x, tt.add) {
			t.Errorf("%v Add %v = %v, want %v", tt.v, tt.u, x, tt.add)
		}
	}
}
func TestSub(t *testing.T) {
	for _, tt := range binarytests {
		if x := tt.v.Sub(&tt.u); !VectorsAreAlmostEqual(x, tt.sub) {
			t.Errorf("%v Sub %v = %v, want %v", tt.v, tt.u, x, tt.sub)
		}
	}
}
func TestMulV(t *testing.T) {
	for _, tt := range binarytests {
		if x := tt.v.MulV(&tt.u); !VectorsAreAlmostEqual(x, tt.mulv) {
			t.Errorf("%v MulV %v = %v, want %v", tt.v, tt.u, x, tt.mulv)
		}
	}
}

var scalartests = []struct {
	v    Vector3f
	f    float64
	mulf Vector3f
}{
	{Vector3fZERO, 0.0, Vector3fZERO},
	{Vector3fZERO, 1.0, Vector3fZERO},
	{Vector3fONE, 0.0, Vector3fZERO},
	{Vector3fONE, 1.0, Vector3fONE},
	{Vector3fONE, -1.0, Vector3f{-1.0, -1.0, -1.0}},
	{Vector3f{-1.0, -1.0, -1.0}, 1.0, Vector3f{-1.0, -1.0, -1.0}},
	{Vector3f{1.0, 2.0, 3.0}, -1.4, Vector3f{-1.4, -2.8, -4.2}},
}

func TestMulF(t *testing.T) {
	for _, tt := range scalartests {
		if x := tt.v.MulF(tt.f); !VectorsAreAlmostEqual(x, tt.mulf) {
			t.Errorf("%v MulF %v = %v, want %v", tt.v, tt.f, x, tt.mulf)
		}
	}
}

var clampedtests = []struct {
	v   Vector3f
	min Vector3f
	max Vector3f
	out Vector3f
}{
	{Vector3fZERO, Vector3fZERO, Vector3fZERO, Vector3fZERO},
	{Vector3fONE, Vector3fONE, Vector3fONE, Vector3fONE},
	{Vector3fZERO, Vector3fONE, Vector3fONE, Vector3fONE},
	{Vector3fONE, Vector3fZERO, Vector3fZERO, Vector3fZERO},
	{Vector3f{1.0, 2.0, 3.0}, Vector3fZERO, Vector3fONE, Vector3fONE},
	{Vector3f{-1.0, 2.0, 3.0}, Vector3fZERO, Vector3fONE, Vector3f{0.0, 1.0, 1.0}},
	{Vector3f{-1.0, 2.0, 3.0}, Vector3f{-1.0, -1.0, -1.0}, Vector3fONE, Vector3f{-1.0, 1.0, 1.0}},
}

func TestClamped(t *testing.T) {
	for _, tt := range clampedtests {
		if x := tt.v.Clamped(&tt.min, &tt.max); !VectorsAreAlmostEqual(x, tt.out) {
			t.Errorf("%v Clamped(%v,%v) = %v, want %v", tt.v, tt.min, tt.max, x, tt.out)
		}
	}
}
