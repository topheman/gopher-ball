package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	mu       sync.RWMutex
	w        int32
	h        int32
	x        int32
	y        int32
	dx       float32
	dy       float32
	textures map[string]*sdl.Texture
}

func newPlayer(r *sdl.Renderer) (*player, error) {
	textures := make(map[string]*sdl.Texture)
	ballTexture, err := img.LoadTexture(r, "assets/imgs/ball-steel-no-shadow.png")
	if err != nil {
		return nil, fmt.Errorf("Error loading ball texture: %v", err)
	}
	textures["ball"] = ballTexture

	shadowTexture, err := img.LoadTexture(r, "assets/imgs/ball-steel-only-shadow.png")
	if err != nil {
		return nil, fmt.Errorf("Error loading ball shadow texture: %v", err)
	}
	textures["shadow"] = shadowTexture

	return &player{
		w:        50,
		h:        50,
		x:        150,
		y:        500,
		dx:       1.5,
		textures: textures,
	}, nil
}

func (p *player) reset(x, y int32) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.x = x
	p.y = y
}

func (p *player) destroy() {
	p.mu.Lock()
	defer p.mu.Unlock()
	defer log.Println("[Player] Player destroyed")
	for _, t := range p.textures {
		t.Destroy()
	}
}

func (p *player) update() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.x = int32((float32(p.x*100) + p.dx*100) / 100)
	p.y = int32((float32(p.y*100) + p.dy*100) / 100)
}

func (p *player) render(r *sdl.Renderer) error {
	p.mu.RLock()
	defer p.mu.RUnlock()
	bgRect := &sdl.Rect{X: p.x - p.w/2, Y: p.y - p.h/2, W: p.w, H: p.h * 256 / 218}
	if err := r.Copy(p.textures["shadow"], nil, bgRect); err != nil {
		return fmt.Errorf("could not copy player shadow: %v", err)
	}
	rect := &sdl.Rect{X: p.x - p.w/2, Y: p.y - p.h/2, W: p.w, H: p.h}
	if err := r.Copy(p.textures["ball"], nil, rect); err != nil {
		return fmt.Errorf("could not copy player ball: %v", err)
	}
	return nil
}
