package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var frameNumber int32

type game struct {
	mu       sync.RWMutex
	w        int32
	h        int32
	player   *player
	floor    *floor
	ennemies *ennemies
}

func (g *game) reset() {
	log.Println("[Game] game reseted")
	frameNumber = 0
	g.ennemies.reset()
	g.player.reset()
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
				frameNumber++
				// update coordinates part
				g.player.update()
				g.ennemies.update(frameNumber)
				// g.floor.update() // no need
				// manage collision part
				g.handleCollisions()
				if g.player.isDead() {
					g.reset()
				}
				// render part
				r.Clear()
				if err := g.floor.render(r, frameNumber); err != nil {
					errChannel <- err
				}
				if err := g.ennemies.render(r); err != nil {
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

func (g *game) handleCollisions() {
	// player vs floor
	managePlayerFloorCollision(g.player, g.floor)
	// player vs ennemies
	if g.ennemies.checkCollision(g.player) {
		g.player.die()
	}
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

	player, err := newPlayer(r)
	if err != nil {
		return nil, fmt.Errorf("Error creating player: %v", err)
	}

	floor, err := newFloor(r, w, h)
	if err != nil {
		return nil, fmt.Errorf("Error creating floor: %v", err)
	}
	randomX := floor.compileRandomX()
	isEnnemyOutside := floor.compileIsEnnemyOutside()

	ennemies, err := createEnnemies(r, randomX, isEnnemyOutside)
	if err != nil {
		return nil, fmt.Errorf("Error creating ennemies: %v", err)
	}

	return &game{
		w:        w,
		h:        h,
		player:   player,
		floor:    floor,
		ennemies: ennemies,
	}, nil
}
