package simulation

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	eb "github.com/hajimehoshi/ebiten/v2"
	ebt "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/shmarlovsky/sim1/pkg/gravity"
)

const (
	SIZE = 5
)

type GravitySim struct {
	gravity.ParticleGroup
}

func NewGravitySim(n int) *GravitySim {
	return &GravitySim{
		ParticleGroup: *gravity.RandomParticleGroup(n),
	}
}

func (g *GravitySim) Update() error {
	// gravity.Interaction1(g.Particles, g.Particles, -1)
	gravity.Interaction1(g.Particles, g.Particles, 1)
	// g.ParticleGroup.Step()
	return nil
}

func (g *GravitySim) Draw(screen *eb.Image) {
	ebt.DebugPrint(screen, fmt.Sprintf("FPS: %v", ebiten.ActualFPS()))
	for _, p := range g.ParticleGroup.Particles {
		// TODO: move to `ui`
		img := eb.NewImage(SIZE, SIZE)
		img.Fill(p.Color)
		geoM := eb.GeoM{}
		geoM.Translate(float64(p.X), float64(p.Y))
		screen.DrawImage(img, &eb.DrawImageOptions{
			GeoM: geoM,
		})
	}
}

func (g *GravitySim) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
