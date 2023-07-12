package projector

import (
	g "xubiod/really-bad-renderer"
)

func Perspective(points []g.Point3, camera g.Point3, camera_orientation g.Point3, display_surface g.Point3) []g.Point2 {
	r2c := make([]g.Point3, 0)

	for _, vtx := range points {
		r2c = append(r2c, g.Point3{X: vtx.X - camera.X, Y: vtx.Y - camera.Y, Z: vtx.Z - camera.Z})
	}

	camera_transform := g.Rotate(r2c, -camera_orientation.X, -camera_orientation.Y, -camera_orientation.Z)

	screen := make([]g.Point2, 0)

	for _, vtx := range camera_transform {
		bx := (display_surface.Z/vtx.Z)*vtx.X + display_surface.X
		by := (display_surface.Z/vtx.Z)*vtx.Y + display_surface.Y

		screen = append(screen, g.Point2{X: bx, Y: by})
	}

	return screen
}
