package simulation

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	eb "github.com/hajimehoshi/ebiten/v2"
	ebt "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/shmarlovsky/sim1/pkg/gravity"
	"github.com/shmarlovsky/sim1/pkg/ui"
	"golang.org/x/image/colornames"
)

type GravitySim struct {
	// contains all
	Particles []*gravity.Particle
	// same particles, divided by groups
	Green  []*gravity.Particle
	Red    []*gravity.Particle
	Yellow []*gravity.Particle
}

func NewGravitySim(n int) *GravitySim {
	yelow := gravity.ColouredParticles(300, colornames.Yellow)
	red := gravity.ColouredParticles(200, colornames.Red)
	green := gravity.ColouredParticles(200, colornames.Green)
	all := make([]*gravity.Particle, 0, len(red)+len(green)+len(green))
	all = append(all, yelow...)
	all = append(all, red...)
	all = append(all, green...)
	return &GravitySim{
		Particles: all,
		Yellow:    yelow,
		Green:     green,
		Red:       red,
	}
}

func (g *GravitySim) Update() error {
	go gravity.Interaction1(g.Red, g.Red, 0.1)
	go gravity.Interaction1(g.Yellow, g.Red, 0.15)
	go gravity.Interaction1(g.Green, g.Green, -0.7)
	go gravity.Interaction1(g.Green, g.Red, -0.2)
	go gravity.Interaction1(g.Red, g.Green, -0.1)

	return nil
}

func (g *GravitySim) Draw(screen *eb.Image) {
	ebt.DebugPrint(screen, fmt.Sprintf("FPS: %v", ebiten.ActualFPS()))
	for _, p := range g.Particles {
		ui.DrawParticle(screen, p.X, p.Y, p.Color)
	}
}

func (g *GravitySim) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
