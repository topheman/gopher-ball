package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	mu           sync.RWMutex
	w            float32
	h            float32
	x            float32
	y            float32
	dx           float32
	dy           float32
	acceleration float32
	textures     map[string]*sdl.Texture
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
		w:            50,
		h:            50,
		x:            150,
		y:            500,
		dx:           0.0,
		dy:           0.0,
		acceleration: 0.05,
		textures:     textures,
	}, nil
}

func (p *player) reset(x, y float32) {
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
	p.x += p.dx
	p.y += p.dy
}

func (p *player) bumpAcceleration() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.acceleration += 0.1
}

func (p *player) updateDirection(ddx, ddy float32) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.dx += p.acceleration * ddx
	p.dy += p.acceleration * ddy
}

func (p *player) render(r *sdl.Renderer) error {
	p.mu.RLock()
	defer p.mu.RUnlock()
	bgRect := &sdl.Rect{X: int32(p.x - p.w/2), Y: int32(p.y - p.h/2), W: int32(p.w), H: int32(p.h * 256 / 218)}
	if err := r.Copy(p.textures["shadow"], nil, bgRect); err != nil {
		return fmt.Errorf("could not copy player shadow: %v", err)
	}
	rect := &sdl.Rect{X: int32(p.x - p.w/2), Y: int32(p.y - p.h/2), W: int32(p.w), H: int32(p.h)}
	if err := r.Copy(p.textures["ball"], nil, rect); err != nil {
		return fmt.Errorf("could not copy player ball: %v", err)
	}
	return nil
}
