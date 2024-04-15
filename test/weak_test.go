package reallybadrenderer_test

import (
	"errors"
	"fmt"
	"io/fs"
	"math"
	"os"
	"testing"

	g "xubiod/really-bad-renderer"
	p "xubiod/really-bad-renderer/projectors"
)

func TestWeak(t *testing.T) {
	c := 0
	model := []g.Point3{
		// 0				1				2				3
		{X: -30, Y: -30, Z: -30}, {X: -30, Y: -30, Z: +30}, {X: -30, Y: +30, Z: +30}, {X: -30, Y: +30, Z: -30},
		// 4				5				6				7
		{X: +30, Y: -30, Z: -30}, {X: +30, Y: -30, Z: +30}, {X: +30, Y: +30, Z: +30}, {X: +30, Y: +30, Z: -30},
	}

	model_lines := []([]int){
		[]int{1, 3, 4}, []int{2, 5}, []int{3}, []int{},
		[]int{5, 7}, []int{}, []int{2, 5, 7}, []int{3},
	}

	_, err := os.Stat("./weak-out")
	if errors.Is(err, fs.ErrNotExist) {
		os.Mkdir("./weak-out", 0666)
	}

	for time := 0.0; time < 4*math.Pi; time += (math.Pi / 32) {
		screen := p.WeakPerspective(g.Rotate(model, 0, time, 0), 100)

		err = g.GeneratePng(screen, model_lines, fmt.Sprintf("weak-out/weak-test_%d.png", c))

		if err != nil {
			t.Error(err)
		}

		c++
	}
}
