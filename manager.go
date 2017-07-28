package main

import (
	"fmt"
	"log"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type manager struct {
	game   *game
	meta   interface{}
	paused bool
}

func (m *manager) reset() {
	log.Println("[Game manager] game reseted")
	m.paused = false
}

func (m *manager) run(r *sdl.Renderer) error {
	log.Println("[Game manager] Game started")
	m.paused = false
	select {
	case <-m.game.run(r):
		return fmt.Errorf("Game loop problem")
	case <-time.After(time.Second * 5):
		return nil
	}
}

func (m *manager) destroy() {
	defer log.Println("[Game manager] Game manager destroyed")
	m.game.destroy()
}

func (m *manager) togglePause() {
	m.paused = !m.paused
	// no ternaries ? wat ?
	if m.paused {
		log.Println("[Game manager] Game paused")
		return
	}
	log.Println("[Game manager] Game resumed")
}

func newManager(r *sdl.Renderer, w, h, metaWidth int32, eventsChan interface{}) (*manager, error) {
	game, err := newGame(r, w-metaWidth, h, eventsChan)
	if err != nil {
		return nil, fmt.Errorf("Error creating game: %v", err)
	}
	return &manager{
		game: game,
		meta: true,
	}, nil
}
