package main

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

type game struct {
	w             int
	h             int
	scene         *scene
	score         int
	playerNbLives int
	eventsChan    interface{}
	paused        bool
}

func (g *game) start() {
	log.Println("Game started")
	g.paused = false
}

func (g *game) togglePause() {
	g.paused = !g.paused
	// no ternaries ? wat ?
	if g.paused {
		log.Println("Game paused")
		return
	}
	log.Println("Game resumed")
}

func newGame(r *sdl.Renderer, w, h, sceneWith int, eventsChan interface{}) (*game, error) {
	scene, err := newScene(r, sceneWith, h, eventsChan)
	if err != nil {
		return nil, fmt.Errorf("Error creating scene: %v", err)
	}
	return &game{
		w:          w,
		h:          h,
		scene:      scene,
		eventsChan: eventsChan,
		paused:     true,
	}, nil
}
