package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"golang.org/x/image/colornames"

	eb "github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREEN_WIDTH  = 800
	SCREEN_HEIGHT = 600
	ENTITY_H      = 20
	ENTITY_W      = 20
	ENTITIES_N    = 20
	SPEED_X       = 3
	SPEED_Y       = 3
)

type Entity struct {
	Img *eb.Image
	X   int
	Y   int
	Dx  int
	Dy  int
}

type Population struct {
	Entities []*Entity
}

func NewPopulation() *Population {
	entities := make([]*Entity, 0, ENTITIES_N)
	for i := 0; i < ENTITIES_N; i++ {
		x := rand.Intn(SCREEN_WIDTH)
		y := rand.Intn(SCREEN_HEIGHT)
		colorN := rand.Intn(len(colornames.Names))
		e := NewEntity(x, y, colornames.Map[colornames.Names[colorN]])
		entities = append(entities, e)
	}
	return &Population{
		Entities: entities,
	}
}

func (p *Population) Step() {
	for _, e := range p.Entities {

		if e.X >= SCREEN_WIDTH-ENTITY_W-5 {
			e.Dx *= -1
		}
		if e.Y >= SCREEN_HEIGHT-ENTITY_H-5 {
			e.Dy *= -1
		}
		if e.X <= 0 {
			e.Dx *= -1
		}
		if e.Y <= 0 {
			e.Dy *= -1
		}
		e.X += e.Dx
		e.Y += e.Dy
	}
}

func (e Entity) String() string {
	return fmt.Sprintf("Entity(%v,%v)", e.X, e.Y)
}

func NewEntity(x, y int, color color.Color) *Entity {
	img := eb.NewImage(ENTITY_H, ENTITY_W)
	img.Fill(color)
	return &Entity{
		X:   x,
		Y:   y,
		Img: img,
		Dx:  SPEED_X * randomDirection(),
		Dy:  SPEED_Y * randomDirection(),
	}
}

type Game struct {
	Population *Population
}

func NewGame() *Game {
	return &Game{
		Population: NewPopulation(),
	}
}

func (g *Game) Update() error {
	g.Population.Step()
	return nil
}

func (g *Game) Draw(screen *eb.Image) {
	for _, e := range g.Population.Entities {
		geoM := eb.GeoM{}
		geoM.Translate(float64(e.X), float64(e.Y))
		screen.DrawImage(e.Img, &eb.DrawImageOptions{
			GeoM: geoM,
		})
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func randomDirection() int {
	r := rand.Intn(3)
	return r - 1
	// if r == 0 {
	// 	return -1
	// }
	// return 1
}

func main() {
	rand.Seed(time.Now().UnixNano())
	eb.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	eb.SetWindowTitle("Sim 1")
	g := NewGame()
	if err := eb.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
