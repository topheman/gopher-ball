package main

import (
	"fmt"
	"log"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type game struct {
	time     int
	w        int32
	h        int32
	player   *player
	floor    *floor
	ennemies []interface{}
}

func (g *game) reset() {
	log.Println("[Game] game reseted")
}

func (g *game) run(r *sdl.Renderer, events <-chan sdl.Event) <-chan error {
	log.Println("[Game] Game started")
	errChannel := make(chan error)
	// render loop
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
	// update loop
	go func() {
		for {
			select {
			case e := <-events:
				if g.handleEvent(e) {
					// in case of a quit event, push to the error channel, that will qui
					errChannel <- fmt.Errorf("Done")
				}
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

// returns true if this is a quit event
func (g *game) handleEvent(event sdl.Event) bool {
	switch event.(type) {
	case *sdl.QuitEvent:
		log.Printf("[Event] %T", event)
		return true
	default:
		log.Printf("[Event] %T", event)
		return false
	}
}

func newGame(r *sdl.Renderer, w, h int32) (*game, error) {

	player, err := newPlayer(r)
	if err != nil {
		return nil, fmt.Errorf("Error creating player: %v", err)
	}

	floor, err := newFloor(r, w, h)
	if err != nil {
		return nil, fmt.Errorf("Error creating floor: %v", err)
	}

	return &game{
		w:      w,
		h:      h,
		player: player,
		floor:  floor,
	}, nil
}
