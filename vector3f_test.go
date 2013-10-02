package main

import (
	"math"
	"testing"
)

const (
	EPS = 1e-14 // criterium for comparing reals
)

func IsAlmostZero(x float64) bool {
	return math.Abs(x) < EPS
}

func AreAlmostEqual(x, y float64) bool {
	return math.Abs(x-y) < EPS
}

func TestDot(t *testing.T) {
	v := Vector3f{[]float64{1.0, 2.0, 3.0}}
	u := Vector3f{[]float64{4.2, 5.1, -2.3}}
	const r = 7.5
	if x := v.Dot(&u); !AreAlmostEqual(x, r) {
		t.Errorf("%v Dot %v = %v, want %v", v, u, x, r)
	}
}

//func (v *Vector3f) Negative() Vector3f {
//func (v *Vector3f) Unitized() Vector3f {
//func (v *Vector3f) Cross(u *Vector3f) Vector3f {
//func (v *Vector3f) Add(u *Vector3f) Vector3f {
//func (v *Vector3f) Sub(u *Vector3f) Vector3f {
//func (v *Vector3f) MulV(u *Vector3f) Vector3f {
//func (v *Vector3f) MulF(f float64) Vector3f {
//func (v *Vector3f) IsZero() bool {
//func (v *Vector3f) Clamped(min, max *Vector3f) {
//func Vector3fRead(r io.Reader) (Vector3f, error) {
