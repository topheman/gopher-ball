package main

// #cgo windows LDFLAGS: -lSDL2 -lSDL2_image -lSDL2_ttf
// #cgo linux freebsd darwin pkg-config: sdl2
// #cgo linux freebsd darwin LDFLAGS: -lSDL2_image -lSDL2_ttf
// #include <sdl_wrapper.h>
import "C"

import (
	"fmt"
	"os"

	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const gameWidth = 600
const gameHeight = 768

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}
}

func run() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return fmt.Errorf("Error initializing SDL: %v", err)
	}
	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		return fmt.Errorf("Error initilizing ttf: %v", err)
	}
	defer ttf.Quit()

	w, r, err := sdl.CreateWindowAndRenderer(gameWidth, gameHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		return fmt.Errorf("Error creating window: %v", err)
	}
	defer w.Destroy()

	_ = r

	if err := drawWelcomeScreen(r); err != nil {
		return fmt.Errorf("Error drawing welcome screen: %v", err)
	}

	time.Sleep(time.Second * 4)

	game, err := newGame(r, gameWidth, gameHeight)
	defer game.destroy()
	if err != nil {
		return fmt.Errorf("Error creating Game: %v", err)
	}

	events := make(chan sdl.Event)
	errorChannel := game.run(r, events)
	defer close(events)
	// wait for events and push them into "events" channel
	for {
		select {
		case events <- sdl.WaitEvent():
		case err := <-errorChannel:
			return err
		}
	}
}

func drawWelcomeScreen(r *sdl.Renderer) error {
	r.Clear()

	splashScreen, err := img.LoadTexture(r, "assets/imgs/splashScreen.jpg")
	if err != nil {
		return fmt.Errorf("Error loading splashScreen texture: %v", err)
	}
	defer splashScreen.Destroy()

	r.Copy(splashScreen, nil, nil)

	r.Present()

	return nil
}
