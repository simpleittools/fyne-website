package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	w.SetContent(makeSign())
	w.SetPadded(false)
	w.Resize(fyne.NewSize(120, 120))
	w.ShowAndRun()
}

func makeSign() fyne.CanvasObject {
	bg := canvas.NewCircle(color.NRGBA{255, 0, 0, 255})
	bg.StrokeColor = color.White
	bg.StrokeWidth = 5

	bar := canvas.NewRectangle(color.White)

	// manual layouts, as used here, do not automatically resize.
	c := container.NewWithoutLayout(bg, bar)

	bg.Resize(fyne.NewSize(100, 100))
	bg.Move(fyne.NewPos(10, 10))
	bar.Resize(fyne.NewSize(80, 20))
	bar.Move(fyne.NewPos(20, 50))

	return c
}
