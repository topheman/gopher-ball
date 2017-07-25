package main

import (
	"fmt"
	"os"

	"time"

	"github.com/veandco/go-sdl2/sdl"
)

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

	window, renderer, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return fmt.Errorf("Error creating window: %v", err)
	}
	defer window.Destroy()

	_ = renderer

	time.Sleep(time.Second * 3)

	return nil
}
