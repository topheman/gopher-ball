package main

import (
	"fmt"
	"log"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type game struct {
	time       int
	w          int32
	h          int32
	player     *player
	floor      *floor
	ennemies   []interface{}
	eventsChan interface{}
}

func (g *game) reset() {
	log.Println("[Game] game reseted")
}

func (g *game) run(r *sdl.Renderer) <-chan error {
	log.Println("[Game] Game started")
	errChannel := make(chan error)
	// update + render loop
	go func() {
		defer close(errChannel)
		tick := time.Tick(5 * time.Millisecond)
		for {
			select {
			case <-tick:
				// update part
				g.player.update()
				// render part
				r.Clear()
				if err := g.floor.render(r); err != nil {
					errChannel <- err
				}
				if err := g.player.render(r); err != nil {
					errChannel <- err
				}
				r.Present()
			}
		}
	}()
	return errChannel
}

func (g *game) destroy() {
	defer log.Println("[Game] Game destroyed")
	g.player.destroy()
	g.floor.destroy()
}

func (g *game) update() {
	log.Println("[Game] Game update")
}

func (g *game) render() {
	log.Println("[Game] Game render")
}

func newGame(r *sdl.Renderer, w, h int32, eventsChan interface{}) (*game, error) {

	player, err := newPlayer(r)
	if err != nil {
		return nil, fmt.Errorf("Error creating player: %v", err)
	}

	floor, err := newFloor(r, w, h)
	if err != nil {
		return nil, fmt.Errorf("Error creating floor: %v", err)
	}

	return &game{
		w:          w,
		h:          h,
		player:     player,
		floor:      floor,
		eventsChan: eventsChan,
	}, nil
}
