package gravity

import (
	"image/color"
	"math"
	"math/rand"

	"golang.org/x/image/colornames"
)

const (
	MAX_X          = 800
	MAX_Y          = 800
	GRAVITY_RADIUS = 100
	G              = 9.83
)

type Particle struct {
	M     float64
	X     float64
	Y     float64
	Vx    float64
	Vy    float64
	Color color.Color
}

func NewParticle(m, x, y, vx, vy float64, color color.Color) *Particle {
	return &Particle{
		M:     m,
		X:     x,
		Y:     y,
		Vx:    vx,
		Vy:    vy,
		Color: color,
	}
}

type ParticleGroup struct {
	Particles []*Particle
}

func RandomParticleGroup(n int) *ParticleGroup {
	return &ParticleGroup{
		Particles: RandomParticles(n),
	}
}

func (pg *ParticleGroup) Step() {
}

func RandomParticles(number int) []*Particle {
	particles := make([]*Particle, 0, number)
	for i := 0; i < number; i++ {
		x, y := RandomPosition()
		particles = append(particles, NewParticle(3, x, y, 2, 2, colornames.Lightyellow))
		// particles = append(particles, NewParticle(3, x, y, 3, 3, randomColor()))
	}
	return particles
}

func RandomPosition() (x, y float64) {
	return float64(rand.Intn(MAX_X)), float64(rand.Intn(MAX_Y))
}

func RandomColor() color.Color {
	colorInd := rand.Intn(len(colornames.Names))
	return colornames.Map[colornames.Names[colorInd]]
}

// negative g - attraction force, positive - repulsion
func Interaction1(p1, p2 []*Particle, g float64) {
	var a *Particle
	var b *Particle

	for i := 0; i < len(p1); i++ {
		fx := 0.0
		fy := 0.0
		for j := 0; j < len(p2); j++ {
			a = p1[i]
			b = p2[j]
			dx := a.X - b.X
			dy := a.Y - b.Y
			distance := math.Sqrt(dx*dx + dy*dy)

			if distance > 0 && distance < GRAVITY_RADIUS {
				f := g * 1 / distance
				fx += (f * dx)
				fy += (f * dy)
			}
		}
		a.Vx = (a.Vx + fx) * 0.5
		a.Vy = (a.Vy + fy) * 0.5
		a.X += a.Vx
		a.Y += a.Vy
		// not go beyond screen
		if a.X <= 0 || a.X >= MAX_X {
			a.Vx *= -1
		}
		if a.Y <= 0 || a.Y >= MAX_Y {
			a.Vy *= -1
		}
	}
}
