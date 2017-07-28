package main

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type game struct {
	w          int
	h          int
	player     interface{}
	ennemies   []interface{}
	textures   map[string]*sdl.Texture
	eventsChan interface{}
}

func (g *game) reset() {
	log.Println("[Game] game reseted")
}

func (g *game) start() {
	log.Println("[Game] Game started")
}

func (g *game) destroy() {
	log.Println("[Game] Game destroyed")
}

func newGame(r *sdl.Renderer, w, h int, eventsChan interface{}) (*game, error) {
	textures := make(map[string]*sdl.Texture)
	// destroy ?
	bgTexture, err := img.LoadTexture(r, "assets/imgs/wood-background.png")
	if err != nil {
		return nil, fmt.Errorf("Error loading bg texture: %v", err)
	}
	textures["bg"] = bgTexture

	wallTexture, err := img.LoadTexture(r, "assets/imgs/wall-wood.png")
	if err != nil {
		return nil, fmt.Errorf("Error loading wall texture: %v", err)
	}
	textures["wall"] = wallTexture

	return &game{
		w:          w,
		h:          h,
		eventsChan: eventsChan,
		textures:   textures,
	}, nil
}
