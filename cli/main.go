package main

import (
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"

	g "xubiod/really-bad-renderer"
	p "xubiod/really-bad-renderer/projectors"
)

const width float64 = 180.0
const height float64 = 50.0

const leng float64 = 15

var model []g.Point3 = []g.Point3{
	// 0							1								2								3
	{X: -leng, Y: -leng, Z: -leng}, {X: -leng, Y: -leng, Z: +leng}, {X: -leng, Y: +leng, Z: +leng}, {X: -leng, Y: +leng, Z: -leng},
	// 4							5								6								7
	{X: +leng, Y: -leng, Z: -leng}, {X: +leng, Y: -leng, Z: +leng}, {X: +leng, Y: +leng, Z: +leng}, {X: +leng, Y: +leng, Z: -leng},
}

var model_lines []([]int) = []([]int){
	[]int{1, 3, 4}, []int{2, 5}, []int{3}, []int{},
	[]int{5, 7}, []int{}, []int{2, 5, 7}, []int{3},
}

var ux, uy, uz, grain float64

func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	quit := func() {
		s.Fini()
		os.Exit(0)
	}

	grain = 10

	// event handling
	go func(s tcell.Screen) {
		for {
			event := s.PollEvent()

			switch ev := event.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyESC:
					quit()

				case tcell.KeyLeft:
					ux -= 0.2 / grain

				case tcell.KeyRight:
					ux += 0.2 / grain

				case tcell.KeyUp:
					uy -= 0.2 / grain

				case tcell.KeyDown:
					uy += 0.2 / grain

				}
			}
		}
	}(s)

	for {
		s.Clear()
		trans_model := g.Rotate(model, ux, uy, uz)
		projected := p.WeakPerspective(trans_model, 100)

		for li, lrd := range model_lines {
			for _, ld := range lrd {
				for i := 0.0; i < 1.0; i += 0.02 {
					rx := (projected[ld].X * i) + (projected[li].X * (1 - i))
					ry := (projected[ld].Y * i) + (projected[li].Y * (1 - i))
					rz := (trans_model[ld].Z * i) + (trans_model[li].Z * (1 - i))
					rx *= 16. / 9.
					s.SetContent(int(rx+width/2), int(ry+height/2), ' ', nil, tcell.StyleDefault.Background(tcell.PaletteColor(int(rz+1024)%256)).Foreground(tcell.ColorBlack))
				}
			}
		}

		for idx, pt := range projected {
			s.SetContent(int((pt.X*(16.0/9.0))+width/2), int(pt.Y+height/2), ' ', nil, tcell.StyleDefault.Background(tcell.PaletteColor(int(trans_model[idx].Z+1024)%256)).Foreground(tcell.ColorBlack))
		}

		s.Show()

		s.Fill(' ', tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorBlack))

		ux += 0.01
		uy += 0.02
		uy += 0.03

		time.Sleep(16 * time.Millisecond)
	}
}
