#if defined(__WIN32)
	#include <SDL2/SDL.h>
	#include <SDL2/SDL_image.h>
	#include <SDL2/SDL_ttf.h>
	#include <stdlib.h>
#else
	#include <SDL.h>
	#include <SDL_image.h>
	#include <SDL_ttf.h>
#endif

#if !defined(SDL_WINDOW_ALLOW_HIGHDPI)
	#define SDL_WINDOW_ALLOW_HIGHDPI (0x00002000)
#endif
