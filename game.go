package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var frameNumber int32

type game struct {
	mu        sync.RWMutex
	w         int32
	h         int32
	player    *player
	floor     *floor
	scoreFont *ttf.Font
	ennemies  *ennemies
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
				g.renderScore(r)
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
	g.scoreFont.Close()
}

func (g *game) update() {
	log.Println("[Game] Game update")
}

func (g *game) renderScore(r *sdl.Renderer) error {
	surface, err := g.scoreFont.RenderUTF8_Solid("Score: "+strconv.Itoa(int(frameNumber)), sdl.Color{
		R: 144,
		G: 0,
		B: 0,
		A: 1,
	})
	defer surface.Free()
	if err != nil {
		return fmt.Errorf("Error creating score surface: %v", err)
	}
	texture, err := r.CreateTextureFromSurface(surface)
	defer texture.Destroy()
	surfaceRect := &sdl.Rect{X: 200, Y: g.h - 80, W: 300, H: 60}
	if err := r.Copy(texture, nil, surfaceRect); err != nil {
		return fmt.Errorf("could not copy score texture: %v", err)
	}
	return nil
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
		return false
	default:
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

	scoreFont, err := ttf.OpenFont("./assets/fonts/UbuntuMono-B.ttf", 80)
	if err != nil {
		return nil, fmt.Errorf("Error opening font: %v", err)
	}

	return &game{
		w:         w,
		h:         h,
		player:    player,
		floor:     floor,
		ennemies:  ennemies,
		scoreFont: scoreFont,
	}, nil
}
