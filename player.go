package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	w       int32
	h       int32
	x       int32
	y       int32
	dx      int32
	dy      int32
	texture *sdl.Texture
}

func newPlayer(r *sdl.Renderer) (*player, error) {
	texture, err := img.LoadTexture(r, "assets/imgs/ball-steel.png")
	if err != nil {
		return nil, fmt.Errorf("Error loading ball texture: %v", err)
	}
	return &player{
		w:       50,
		h:       50,
		x:       150,
		y:       500,
		texture: texture,
	}, nil
}

func (p *player) reset(x, y int32) {
	p.x = x
	p.y = y
}

func (p *player) render(r *sdl.Renderer) error {
	rect := &sdl.Rect{X: p.x - p.w/2, Y: p.y - p.h/2 + 10, W: p.w, H: p.h}
	if err := r.Copy(p.texture, nil, rect); err != nil {
		return fmt.Errorf("could not copy player background: %v", err)
	}
	return nil
}
