package projector

import (
	g "xubiod/really-bad-renderer"

	"gonum.org/v1/gonum/mat"
)

func WeakPerspective(points []g.Point3, focal_length float64) []g.Point2 {
	screen := make([]g.Point2, 0)

	pt := mat.NewDense(1, 3, nil)

	for _, point := range points {
		pt.Set(0, 0, point.X)
		pt.Set(0, 1, point.Y)
		pt.Set(0, 2, point.Z)

		fpt := g.Point3{X: pt.At(0, 0), Y: pt.At(0, 1), Z: pt.At(0, 2)}

		screen = append(screen, g.Point2{X: (fpt.X * focal_length) / (focal_length + fpt.Z), Y: (fpt.Y * focal_length) / (focal_length + fpt.Z)})
	}

	return screen
}
