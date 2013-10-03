package main

import (
	"fmt"
	"io"
	"math"
)

// Vector3f is an 3D vector/point.
type Vector3f [3]float64

// Dot computes the dot product of v and u.
func (v *Vector3f) Dot(u *Vector3f) float64 {
	return v[0]*u[0] + v[1]*u[1] + v[2]*u[2]
}

// Negative returns a new vector with all elements negated.
func (v *Vector3f) Negative() Vector3f {
	var u Vector3f
	u[0] = -v[0]
	u[1] = -v[1]
	u[2] = -v[2]
	return u
}

// Unitize returns a new vector in the same direction of v with length 1.
func (v *Vector3f) Unitized() Vector3f {
	length := math.Sqrt(v.Dot(v))
	var oneOverLength float64
	if length != 0.0 {
		oneOverLength = 1.0 / length
	}
	return v.MulF(oneOverLength)
}

// Cross returns the cross product of v and u.
func (v *Vector3f) Cross(u *Vector3f) Vector3f {
	var w Vector3f
	w[0] = (v[1] * u[2]) - (v[2] * u[1])
	w[1] = (v[2] * u[0]) - (v[0] * u[2])
	w[2] = (v[0] * u[1]) - (v[1] * u[0])
	return w
}

// Add returns the result of v+u.
func (v *Vector3f) Add(u *Vector3f) Vector3f {
	var w Vector3f
	w[0] = v[0] + u[0]
	w[1] = v[1] + u[1]
	w[2] = v[2] + u[2]
	return w
}

// Sub returns the result of v-u.
func (v *Vector3f) Sub(u *Vector3f) Vector3f {
	var w Vector3f
	w[0] = v[0] - u[0]
	w[1] = v[1] - u[1]
	w[2] = v[2] - u[2]
	return w
}

// MulV returns a new vector where each element w_i is the product v_i*u_i.
func (v *Vector3f) MulV(u *Vector3f) Vector3f {
	var w Vector3f
	w[0] = v[0] * u[0]
	w[1] = v[1] * u[1]
	w[2] = v[2] * u[2]
	return w
}

// MulF returns the scalar product of v and f.
func (v *Vector3f) MulF(f float64) Vector3f {
	var w Vector3f
	w[0] = v[0] * f
	w[1] = v[1] * f
	w[2] = v[2] * f
	return w
}

// IsZero returns true if v has zero length.
func (v *Vector3f) IsZero() bool {
	// TODO: should this have a tolerance?
	return (v[0] == 0.0) && (v[1] == 0.0) && (v[2] == 0.0)
}

// Clamped returns v adjusted to be within min and max.
func (v *Vector3f) Clamped(min, max *Vector3f) Vector3f {
	var w Vector3f
	for i := 0; i < 3; i++ {
		switch {
		case v[i] < min[i]:
			w[i] = min[i]
		case v[i] > max[i]:
			w[i] = max[i]
		default:
			w[i] = v[i]
		}
	}
	return w
}

var (
	Vector3fZERO = Vector3f{0.0, 0.0, 0.0}
	Vector3fONE  = Vector3f{1.0, 1.0, 1.0}
)

// Vector3fRead reads a single vector from r.
func Vector3fRead(r io.Reader) (Vector3f, error) {
	var v Vector3f
	var s0, s1 string
	if _, err := fmt.Fscanf(r, "%1s %g %g %g %1s", &s0, &v[0], &v[1],
		&v[2], &s1); err != nil {
		return v, err
	}
	if s0 != "(" || s1 != ")" {
		return v, fmt.Errorf("read error")
	}
	return v, nil
}
