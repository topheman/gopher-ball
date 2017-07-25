package main

import (
	"fmt"
	"os"

	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing %v", err)
		os.Exit(2)
	}
	defer sdl.Quit()

	window, renderer, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating window %v", err)
		os.Exit(3)
	}
	defer window.Destroy()

	_ = renderer

	time.Sleep(time.Second * 3)
}
