package reallybadrenderer

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func GeneratePng(screen []Point2, model_lines []([]int), fn string) error {
	r := image.Rect(0, 0, 400, 400)
	m := image.NewNRGBA(r)

	for li, lrd := range model_lines {
		for _, ld := range lrd {
			for i := 0.0; i < 1.0; i += 0.002 {
				rx := (screen[ld].X * i) + (screen[li].X * (1 - i))
				ry := (screen[ld].Y * i) + (screen[li].Y * (1 - i))
				m.SetNRGBA(int(rx)+200, int(ry)+200, color.NRGBA{0, 0xFF, 0xFF, 0xFF})
			}
		}
	}

	for _, pt := range screen {
		m.SetNRGBA(int(pt.X)+200, int(pt.Y)+200, color.NRGBA{0xFF, 0xFF, 0xFF,
			/*uint8((float64(idx) / float64(len(screen))) * float64(len(screen))),
			uint8((float64(idx) / float64(len(screen))) * float64(len(screen))),
			uint8((float64(idx) / float64(len(screen))) * float64(len(screen))),*/
			0xFF})
	}

	f, err := os.Create(fn)
	if err != nil {
		return err
	}

	if err := png.Encode(f, m); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}
