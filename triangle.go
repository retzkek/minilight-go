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
