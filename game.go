package main

import (
	"fmt"
	"log"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type game struct {
	w          int
	h          int
	player     *player
	ennemies   []interface{}
	textures   map[string]*sdl.Texture
	eventsChan interface{}
}

func (g *game) reset() {
	log.Println("[Game] game reseted")
}

func (g *game) run(r *sdl.Renderer) <-chan error {
	log.Println("[Game] Game started")
	errChannel := make(chan error)
	go func() {
		defer close(errChannel)
		for range time.Tick(time.Millisecond) {
			r.Clear()
			if err := g.player.render(r); err != nil {
				errChannel <- err
			}
			r.Present()
		}
	}()
	return errChannel
}

func (g *game) destroy() {
	defer log.Println("[Game] Game destroyed")
	g.player.destroy()
	for _, t := range g.textures {
		t.Destroy()
	}
}

func (g *game) update() {
	log.Println("[Game] Game update")
}

func (g *game) render() {
	log.Println("[Game] Game render")
}

func newGame(r *sdl.Renderer, w, h int, eventsChan interface{}) (*game, error) {
	textures := make(map[string]*sdl.Texture)
	// destroy ?
	bgTexture, err := img.LoadTexture(r, "assets/imgs/wood-background.jpg")
	if err != nil {
		return nil, fmt.Errorf("Error loading bg texture: %v", err)
	}
	textures["bg"] = bgTexture

	wallTexture, err := img.LoadTexture(r, "assets/imgs/wall-wood.png")
	if err != nil {
		return nil, fmt.Errorf("Error loading wall texture: %v", err)
	}
	textures["wall"] = wallTexture

	player, err := newPlayer(r)
	if err != nil {
		return nil, fmt.Errorf("Error creating player: %v", err)
	}

	return &game{
		w:          w,
		h:          h,
		player:     player,
		eventsChan: eventsChan,
		textures:   textures,
	}, nil
}
