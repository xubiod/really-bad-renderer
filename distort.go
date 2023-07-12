package reallybadrenderer

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

func Rotate(model []Point3, yaw, pitch, roll float64) []Point3 {
	trans_model := make([]Point3, 0)

	a_s := math.Sin(yaw)
	a_c := math.Cos(yaw)
	b_s := math.Sin(pitch)
	b_c := math.Cos(pitch)
	y_s := math.Sin(roll)
	y_c := math.Cos(roll)

	rot_slice := []float64{
		a_c * b_c, (a_c * b_s * y_s) - (a_s * y_c), (a_c * b_s * y_c) + (a_s * y_s),
		a_s * b_c, (a_s * b_s * y_s) + (a_c * y_c), (a_s * b_s * y_c) - (a_c * y_s),
		-b_s, b_s * y_s, b_c * y_c,
	}

	rot := mat.NewDense(3, 3, rot_slice)

	pt := mat.NewDense(1, 3, nil)
	var trans_pt mat.Dense

	for _, point := range model {
		pt.Set(0, 0, point.X)
		pt.Set(0, 1, point.Y)
		pt.Set(0, 2, point.Z)

		trans_pt.Mul(pt, rot)

		trans_model = append(trans_model, Point3{X: trans_pt.At(0, 0), Y: trans_pt.At(0, 1), Z: trans_pt.At(0, 2)})
	}

	return trans_model
}

func Translate(model []Point3, x, y, z float64) []Point3 {
	trans_model := make([]Point3, 0)

	trn_slice := []float64{
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	}

	trn := mat.NewDense(4, 4, trn_slice)

	pt := mat.NewDense(1, 4, nil)
	var trans_pt mat.Dense

	for _, point := range model {
		pt.Set(0, 0, point.X)
		pt.Set(0, 1, point.Y)
		pt.Set(0, 2, point.Z)
		pt.Set(0, 3, 1)

		trans_pt.Mul(pt, trn)

		trans_model = append(trans_model, Point3{X: trans_pt.At(0, 0), Y: trans_pt.At(0, 1), Z: trans_pt.At(0, 2)})
	}

	return trans_model
}

func Scale(model []Point3, x, y, z float64) []Point3 {
	trans_model := make([]Point3, 0)

	trn_slice := []float64{
		x, 0, 0,
		0, y, 0,
		0, 0, z,
	}

	trn := mat.NewDense(3, 3, trn_slice)

	pt := mat.NewDense(1, 3, nil)
	var trans_pt mat.Dense

	for _, point := range model {
		pt.Set(0, 0, point.X)
		pt.Set(0, 1, point.Y)
		pt.Set(0, 2, point.Z)

		trans_pt.Mul(pt, trn)

		trans_model = append(trans_model, Point3{X: trans_pt.At(0, 0), Y: trans_pt.At(0, 1), Z: trans_pt.At(0, 2)})
	}

	return trans_model
}
