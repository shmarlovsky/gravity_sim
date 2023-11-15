package main

import (
	"log"
	"math/rand"
	"time"

	eb "github.com/hajimehoshi/ebiten/v2"
	"github.com/shmarlovsky/sim1/pkg/simulation"
)

const (
	SCREEN_WIDTH  = 600
	SCREEN_HEIGHT = 600
	ENTITY_H      = 20
	ENTITY_W      = 20
	ENTITIES_N    = 20
	SPEED_X       = 3
	SPEED_Y       = 3
)

func main() {
	rand.Seed(time.Now().UnixNano())
	eb.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	eb.SetWindowTitle("Gravity")
	s := simulation.NewGravitySim(40)
	if err := eb.RunGame(s); err != nil {
		log.Fatal(err)
	}
}
