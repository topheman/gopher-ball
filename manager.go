package main

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

type manager struct {
	game   interface{}
	meta   interface{}
	paused bool
}

func (m *manager) reset() {
	log.Println("[Game manager] game reseted")
	m.paused = false
}

func (m *manager) start() {
	log.Println("[Game manager] Game started")
	m.paused = false
}

func (m *manager) destroy() {
	log.Println("[Game manager] Game started")
	m.paused = false
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

func newManager(r *sdl.Renderer, w, h int, eventsChan interface{}) (*manager, error) {
	game, err := newGame(r, w, h/2, eventsChan)
	if err != nil {
		return nil, fmt.Errorf("Error creating game: %v", err)
	}
	return &manager{
		game: game,
		meta: true,
	}, nil
}
