package tuples

import "math"

type Tuple struct {
	X, Y, Z, W float64
}

func Add(t1, t2 Tuple) Tuple {
	return Tuple{
		X: t1.X + t2.X,
		Y: t1.Y + t2.Y,
		Z: t1.Z + t2.Z,
		W: t1.W + t2.W,
	}
}

const epsilon = 0.00001

func Equal(t1, t2 Tuple) bool {
	return math.Abs(t1.X-t2.X) < epsilon &&
		math.Abs(t1.Y-t2.Y) < epsilon &&
		math.Abs(t1.Z-t2.Z) < epsilon &&
		math.Abs(t1.W-t2.W) < epsilon
}

func Point(x, y, z float64) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: 1.0}
}

func Vector(x, y, z float64) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: 0.0}
}

func Subtract(t1, t2 Tuple) Tuple {
	return Tuple{
		X: t1.X - t2.X,
		Y: t1.Y - t2.Y,
		Z: t1.Z - t2.Z,
		W: t1.W - t2.W,
	}
}

func Negate(t Tuple) Tuple {
	return Tuple{
		X: -t.X,
		Y: -t.Y,
		Z: -t.Z,
		W: -t.W,
	}
}

func Multiply(t Tuple, scalar float64) Tuple {
	return Tuple{
		X: t.X * scalar,
		Y: t.Y * scalar,
		Z: t.Z * scalar,
		W: t.W * scalar,
	}
}

func Divide(t Tuple, scalar float64) Tuple {
	return Tuple{
		X: t.X / scalar,
		Y: t.Y / scalar,
		Z: t.Z / scalar,
		W: t.W / scalar,
	}
}

func Magnitude(t Tuple) float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z + t.W*t.W)
}

func Normalize(t Tuple) Tuple {
	magnitude := Magnitude(t)
	return Tuple{
		X: t.X / magnitude,
		Y: t.Y / magnitude,
		Z: t.Z / magnitude,
		W: t.W / magnitude,
	}
}

func Dot(t1, t2 Tuple) float64 {
	return t1.X*t2.X +
		t1.Y*t2.Y +
		t1.Z*t2.Z +
		t1.W*t2.W
}

func Cross(t1, t2 Tuple) Tuple {
	return Vector(
		t1.Y*t2.Z-t1.Z*t2.Y,
		t1.Z*t2.X-t1.X*t2.Z,
		t1.X*t2.Y-t1.Y*t2.X,
	)
}
