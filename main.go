package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
)

func run() {
	monitorX, monitorY := pixelgl.PrimaryMonitor().Size()
	cfg := pixelgl.WindowConfig{
		Bounds:                 pixel.R(0, 0, monitorX, monitorY),
		VSync:                  true,
		Undecorated:            true,
		TransparentFramebuffer: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(color.RGBA{
		R: 0,
		G: 0,
		B: 0,
		A: 10,
	})

	var u = pixel.V(0.0, 0.0)
	var v = pixel.V(0.0, 0.0)
	drawing := false
	imd := imdraw.New(nil)
	for !win.Closed() {
		if win.Pressed(pixelgl.MouseButtonLeft) {
			win.Clear(color.RGBA{
				R: 0,
				G: 0,
				B: 0,
				A: 10,
			})

			if !drawing {
				u = win.MousePosition()
				drawing = true
			} else {
				v = win.MousePosition()

				rect := pixel.R(u.X, u.Y, v.X, v.Y)

				imd.Clear()
				imd.Color = color.RGBA{
					R: 225,
					G: 225,
					B: 225,
					A: 10,
				}
				imd.Push(rect.Min, rect.Max)
				imd.Rectangle(1)

			}
		}

		if win.JustReleased(pixelgl.MouseButtonLeft) {
			drawing = false
		}
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
