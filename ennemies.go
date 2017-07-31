package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type ennemy struct {
	mu sync.RWMutex
	w  float32
	h  float32
	x  float32
	y  float32
	dy float32
}

func createEnnemy(x float32) *ennemy {
	return &ennemy{
		w:  50,
		h:  50,
		x:  x,
		y:  -50,
		dy: 1,
	}
}

type ennemies struct {
	list     []*ennemy
	randomX  func() float32
	textures map[string]*sdl.Texture
}

func (e *ennemies) update(frameNumber int32) {
	// periodically add ennemy
	if frameNumber%500 == 0 {
		e.list = append(e.list, createEnnemy(e.randomX()))
	}
	// update y
	for _, hole := range e.list {
		hole.mu.RLock()
		hole.y += hole.dy
		hole.mu.RUnlock()
	}
}

func (e *ennemies) render(r *sdl.Renderer) error {
	for _, hole := range e.list {
		hole.mu.RLock()
		bgRect := &sdl.Rect{X: int32(hole.x - hole.w/2), Y: int32(hole.y - hole.h/2), W: int32(hole.w), H: int32(hole.h * 256 / 218)}
		if err := r.Copy(e.textures["hole"], nil, bgRect); err != nil {
			hole.mu.RUnlock()
			return fmt.Errorf("could not copy hole: %v", err)
		}
		hole.mu.RUnlock()
	}
	return nil
}

func (e *ennemies) destroy() {
	defer log.Println("[Ennemies] Ennemies destroyed")
	for _, t := range e.textures {
		t.Destroy()
	}
}

func createEnnemies(r *sdl.Renderer, randomX func() float32) (*ennemies, error) {
	textures := make(map[string]*sdl.Texture)
	ballTexture, err := img.LoadTexture(r, "assets/imgs/ball-hole.png")
	if err != nil {
		return nil, fmt.Errorf("Error loading hole texture: %v", err)
	}
	textures["hole"] = ballTexture

	list := make([]*ennemy, 0)

	return &ennemies{
		list:     list,
		randomX:  randomX,
		textures: textures,
	}, nil
}
