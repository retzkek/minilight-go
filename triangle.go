package main

import (
	"math"
)

const (
	EPSILON   = 1.0 / 1048576.0
	TOLERANCE = 1.0 / 1024.0
)

// Triangle is a simple, explicit triangle geometry with material properties.
type Triangle struct {
	Vertices     []Vector3f
	Reflectivity Vector3f
	Emitivity    Vector3f
}

func TriangleRead(w io.reader) (Triangle, error) {
	var t Triangle
	for i := 0; i < 3; i++ {
		if v, err := Vector3fRead(w); err != nil {
			return t, err
		}
		t.Vertices.append(v)
	}
	if t.Reflectivity, err = Vector3fRead(w); err != nil {
		return t, err
	}
	t.Reflectivity = t.Reflectivity.Clamped(&Vector3fZERO, &Vector3fONE)
	if t.Emitivity, err = Vector3fRead(w); err != nil {
		return t, err
	}
	t.Emitivity = t.Emitivity.Clamped(&Vector3fZERO, &Vector3fONE)
	return t, nil
}

// NormalV returns the normal vector, unnormalized.
func (t *Triangle) NormalV() Vector3f {
	edge1 := t.Vertices[1].Sub(t.Vertices[0])
	edge3 := t.Vertices[2].Sub(t.Vertices[1])
	return edge1.Cross(edge3)
}

// Bound returns an axis-aligned bounding box for the triangle,
// as [xmin, ymin, zmin, xmax, ymax, zmax].
func (t *Triangle) Bound() []float64 {
	b := make([]float64, 6)
	// initialize to one vertex
	for i := 0; i < 6; i++ {
		b[i] = t.Vertices[2][i%3]
	}
	// expand to all vertices
	for i := 0; i < 3; i++ {
		for j := 0; j < 6; j++ {
			v := t.Vertices[i][j%3]
			if d := j / 3; d == 0 {
				v -= TOLERANCE
			} else {
				v += TOLERANCE
			}
			if (b[j] > v) ^ d {
				b[j] = v
			}
		}
	}
}

// Intersection tests if a ray intersects a triangle, and returns the distance
// along the ray that the intersection occurs if so.
func (t *Triangle) Intersection(origin, direction *Vector3f) (isHit bool, hitDistance float64) {
	hitDistance := 0.0
	isHit := false

	// make vectors for two edges sharing vertex 0
	edge1 := t.Vertices[1].Sub(t.Vertices[0])
	edge2 := t.Vertices[2].Sub(t.Vertices[0])

	// begin calculating determinant
	pvec := direction.Cross(&edge2)

	// if determinant is near zero, ray lies in plane of triangle
	if det := edge1.Dot(pvec); (det <= -EPSILON) | (det >= EPSILON) {
		inv_det := 1.0 / det

		// calculate distance from vertex 0 to ray origin
		tvec := origin.Sub(&t.Vertices[0])

		// calculate U parameter and test bounds
		if u := tvec.Dot(&pvec) * inv_det; (u >= 0.0) & (u <= 1.0) {
			// prepare to test V parameter
			qvec := tvec.Cross(&edge1)
			// calculate V parameter and test bounds
			if v := direction.Dot(qvec) * inv_det; (v >= 0.0) & (u+v <= 1.0) {
				// calculate t, ray intersects triangle
				hitDistance = edge2.Dot(qvec) * inv_det
				// only allow intersection in the forward ray direction
				isHit = (hitDistance >= 0.0)
			}
		}
	}
	return isHit, hitDistance
}
