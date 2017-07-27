package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type scene struct {
	w          int
	h          int
	player     interface{}
	ennemies   []interface{}
	textures   map[string]*sdl.Texture
	eventsChan interface{}
}

func newScene(r *sdl.Renderer, w, h int, eventsChan interface{}) (*scene, error) {
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

	return &scene{
		w:          w,
		h:          h,
		eventsChan: eventsChan,
		textures:   textures,
	}, nil
}
