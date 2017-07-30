package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type game struct {
	mu       sync.RWMutex
	time     int
	w        int32
	h        int32
	player   *player
	floor    *floor
	ennemies []interface{}
	level    int
}

func (g *game) reset() {
	log.Println("[Game] game reseted")
}

func (g *game) run(r *sdl.Renderer, events <-chan sdl.Event) <-chan error {
	log.Println("[Game] Game started")
	errChannel := make(chan error)
	// update / render loop
	go func() {
		defer close(errChannel)
		tick := time.Tick(5 * time.Millisecond)
		for {
			select {
			case <-tick:
				// update coordinates part
				g.player.update()
				// manage collision part
				g.handleCollisions()
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
	// event loop
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

func (g *game) bumpLevel() {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.level++
	g.player.bumpAcceleration()
}

func (g *game) handleCollisions() {
	// player vs floor
	managePlayerFloorCollision(g.player, g.floor)
	// player vs ennemies

	// ennemies vs floor

	// ennemie vs ennemies
}

// returns true if this is a quit event
func (g *game) handleEvent(event sdl.Event) bool {
	switch event.(type) {
	case *sdl.QuitEvent:
		log.Printf("[Event] %T", event)
		return true
	case *sdl.KeyDownEvent:
		left := event.(*sdl.KeyDownEvent).Keysym.Sym == sdl.K_LEFT
		right := event.(*sdl.KeyDownEvent).Keysym.Sym == sdl.K_RIGHT
		up := event.(*sdl.KeyDownEvent).Keysym.Sym == sdl.K_UP
		down := event.(*sdl.KeyDownEvent).Keysym.Sym == sdl.K_DOWN
		// sdl.GetKeyboardState() doesn't seem to be working - can't handle two keys at the same time for the moment
		switch event.(*sdl.KeyDownEvent).Keysym.Sym {
		case sdl.K_UP:
			g.player.updateDirection(0, -1)
		case sdl.K_DOWN:
			g.player.updateDirection(0, 1)
		case sdl.K_LEFT:
			g.player.updateDirection(-1, 0)
		case sdl.K_RIGHT:
			g.player.updateDirection(1, 0)
		}
		log.Printf("[Event] %T | up: %v | right: %v | down: %v | left: %v", event, up, right, down, left)
		return false
	default:
		log.Printf("[Event] %T", event)
		return false
	}
}

func newGame(r *sdl.Renderer, w, h int32) (*game, error) {

	level := 1

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
		level:  level,
	}, nil
}
