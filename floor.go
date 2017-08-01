package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const imageHeight = 1400

type floor struct {
	mu       sync.RWMutex
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

func (f *floor) update() {
	f.mu.Lock()
	defer f.mu.Unlock()
}

func (f *floor) render(r *sdl.Renderer, frameNumber int32) error {
	f.mu.RLock()
	defer f.mu.RUnlock()
	var i int32
	for i <= 1 {
		// draw two tiles of floor
		y := (i-1)*f.h + frameNumber%(f.h)
		bgRect := &sdl.Rect{X: 0, Y: y, W: f.w, H: f.h}
		if err := r.Copy(f.textures["bg"], nil, bgRect); err != nil {
			return fmt.Errorf("[Floor] could not copy background: %v", err)
		}
		leftWallRect := &sdl.Rect{X: 0, Y: y, W: 80, H: f.h}
		if err := r.CopyEx(f.textures["wall"], nil, leftWallRect, 0, nil, sdl.FLIP_HORIZONTAL); err != nil {
			return fmt.Errorf("[Floor] could not copy leftWallRect: %v", err)
		}
		rightWallRect := &sdl.Rect{X: f.w - 80, Y: y, W: 80, H: f.h}
		if err := r.Copy(f.textures["wall"], nil, rightWallRect); err != nil {
			return fmt.Errorf("[Floor] could not copy rightWallRect: %v", err)
		}
		i++
	}
	return nil
}

// returns a function that return a random x coordinate, based on ennemy width (injected by ennemies class)
func (f *floor) compileRandomX() func(holeWidth float32) float32 {
	return func(holeWidth float32) float32 {
		return rand.Float32()*float32(f.w-f.wall*2-int32(holeWidth)) + float32(f.wall+int32(holeWidth)/2)
	}
}

// returns a function that returns true if the ennemy is out of the viewport (to remove it)
func (f *floor) compileIsEnnemyOutside() func(e *ennemy) bool {
	return func(e *ennemy) bool {
		return int32(e.y) > f.h
	}
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
