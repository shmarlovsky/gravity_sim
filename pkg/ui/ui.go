package ui

import (
	"image/color"

	eb "github.com/hajimehoshi/ebiten/v2"
)

const (
	DEFAULT_SIZE = 5
)

func DrawParticle(dst *eb.Image, x, y float64, color color.Color) {
	img := eb.NewImage(DEFAULT_SIZE, DEFAULT_SIZE)
	img.Fill(color)
	geoM := eb.GeoM{}
	geoM.Translate(x, y)
	dst.DrawImage(img, &eb.DrawImageOptions{
		GeoM: geoM,
	})
}
