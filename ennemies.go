package main

import (
	"fmt"
	"log"
	"math"
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

const (
	ennemyWidth float32 = 50
)

func createEnnemy(x float32) *ennemy {
	return &ennemy{
		w:  ennemyWidth,
		h:  ennemyWidth,
		x:  x,
		y:  -50,
		dy: 1,
	}
}

type ennemies struct {
	list            []*ennemy
	randomX         func(hole float32) float32
	isEnnemyOutside func(e *ennemy) bool
	textures        map[string]*sdl.Texture
}

func (e *ennemies) reset() {
	// reset slice + release memory
	e.list = nil
}

func (e *ennemies) checkCollision(p *player) bool {
	for _, hole := range e.list {
		dx := math.Pow(math.Abs(float64(hole.x-p.x)), 2)
		dy := math.Pow(math.Abs(float64(hole.y-p.y)), 2)
		if math.Sqrt(dx+dy) < float64(hole.w/2) {
			return true
		}
	}
	return false
}

func (e *ennemies) update(frameNumber int32) {
	// periodically add ennemy
	if frameNumber%500 == 0 {
		// todo process ennemyWidth based on frameNumber (the bigger hole, the bigger frame number)
		e.list = append(e.list, createEnnemy(e.randomX(ennemyWidth)))
	}
	// update y + remove hole outside of viewport
	for i := len(e.list) - 1; i >= 0; i-- {
		hole := e.list[i]
		hole.y += hole.dy
		if e.isEnnemyOutside(hole) {
			e.list = append(e.list[:i], e.list[i+1:]...)
		}
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

func createEnnemies(r *sdl.Renderer, randomX func(hole float32) float32, isEnnemyOutside func(e *ennemy) bool) (*ennemies, error) {
	textures := make(map[string]*sdl.Texture)
	ballTexture, err := img.LoadTexture(r, "assets/imgs/ball-hole.png")
	if err != nil {
		return nil, fmt.Errorf("Error loading hole texture: %v", err)
	}
	textures["hole"] = ballTexture

	list := make([]*ennemy, 0)

	return &ennemies{
		list:            list,
		randomX:         randomX,
		isEnnemyOutside: isEnnemyOutside,
		textures:        textures,
	}, nil
}
