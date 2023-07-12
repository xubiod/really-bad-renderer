package projector

import (
	g "xubiod/really-bad-renderer"

	"gonum.org/v1/gonum/mat"
)

func Ortho(points []g.Point3, left, right, bottom, top, near, far float64) []g.Point2 {
	screen := make([]g.Point2, 0)

	proj_slice := []float64{
		2 / (right - left), 0, 0, -((right + left) / (right - left)),
		0, 2 / (top - bottom), 0, -((top + bottom) / (top - bottom)),
		0, 0, -2 / (far - near), -((far + near) / (far - near)),
		0, 0, 0, 1,
	}

	proj := mat.NewDense(4, 4, proj_slice)

	pt := mat.NewDense(1, 4, nil)
	var trans_pt mat.Dense

	for _, point := range points {
		pt.Set(0, 0, point.X)
		pt.Set(0, 1, point.Y)
		pt.Set(0, 2, point.Z)
		pt.Set(0, 3, 1)

		trans_pt.Mul(pt, proj)

		screen = append(screen, g.Point2{X: trans_pt.At(0, 0), Y: trans_pt.At(0, 1)})
	}

	return screen
}
