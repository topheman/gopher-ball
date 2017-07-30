package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const imageHeight = 1400

type floor struct {
	mu       sync.RWMutex
	x        int32
	w        int32
	h        int32
	wall     int32
	textures map[string]*sdl.Texture
}

func (f *floor) destroy() {
	defer log.Println("[Floor] Floor destroyed")
	for _, t := range f.textures {
		t.Destroy()
	}
}

func (f *floor) update(dx int32) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.x += dx
}

func (f *floor) render(r *sdl.Renderer) error {
	f.mu.RLock()
	defer f.mu.RUnlock()
	bgRect := &sdl.Rect{X: 0, Y: 0, W: f.w, H: 1400}
	if err := r.Copy(f.textures["bg"], nil, bgRect); err != nil {
		return fmt.Errorf("[Floor] could not copy background: %v", err)
	}
	leftWallRect := &sdl.Rect{X: 0, Y: 0, W: 80, H: 1400}
	if err := r.CopyEx(f.textures["wall"], nil, leftWallRect, 0, nil, sdl.FLIP_HORIZONTAL); err != nil {
		return fmt.Errorf("[Floor] could not copy leftWallRect: %v", err)
	}
	rightWallRect := &sdl.Rect{X: f.w - 80, Y: 0, W: 80, H: 1400}
	if err := r.Copy(f.textures["wall"], nil, rightWallRect); err != nil {
		return fmt.Errorf("[Floor] could not copy rightWallRect: %v", err)
	}
	return nil
}

func newFloor(r *sdl.Renderer, gameWidth, gameHeight int32) (*floor, error) {
	textures := make(map[string]*sdl.Texture)
	bgTexture, err := img.LoadTexture(r, "assets/imgs/wood-background.jpg")
	if err != nil {
		return nil, fmt.Errorf("[Floor] Error loading bg texture: %v", err)
	}
	textures["bg"] = bgTexture

	wallTexture, err := img.LoadTexture(r, "assets/imgs/wall-wood.png")
	if err != nil {
		return nil, fmt.Errorf("[Floor] Error loading wall texture: %v", err)
	}
	textures["wall"] = wallTexture

	return &floor{
		w:        gameWidth,
		h:        gameHeight,
		wall:     60,
		textures: textures,
	}, nil
}
