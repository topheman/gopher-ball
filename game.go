package main

import (
	"fmt"
	"log"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type game struct {
	time       int
	w          int32
	h          int32
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
			if err := g.renderBackground(r); err != nil {
				errChannel <- err
			}
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

func newGame(r *sdl.Renderer, w, h int32, eventsChan interface{}) (*game, error) {
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

// background

func (g *game) renderBackground(r *sdl.Renderer) error {
	bgRect := &sdl.Rect{X: 0, Y: 0, W: g.w, H: 1400}
	if err := r.Copy(g.textures["bg"], nil, bgRect); err != nil {
		return fmt.Errorf("could not copy background: %v", err)
	}
	leftWallRect := &sdl.Rect{X: 0, Y: 0, W: 80, H: 1400}
	if err := r.CopyEx(g.textures["wall"], nil, leftWallRect, 0, nil, sdl.FLIP_HORIZONTAL); err != nil {
		return fmt.Errorf("could not copy leftWallRect: %v", err)
	}
	rightWallRect := &sdl.Rect{X: g.w - 80, Y: 0, W: 80, H: 1400}
	if err := r.Copy(g.textures["wall"], nil, rightWallRect); err != nil {
		return fmt.Errorf("could not copy rightWallRect: %v", err)
	}
	return nil
}
