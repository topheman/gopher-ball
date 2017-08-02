gopher-ball
===========

## Goal

I have a JavaScript background (both frontend and NodeJS). I started go a few weeks ago (really enjoy it - just like a NodeJS with pointers and threads 😜).

I needed some project to test on. For the last years, I made a few [video games in JavaScript](http://dev.topheman.com/my-projects/) and I think it's a good way to learn a new programming language (some will tell that it's a little weird to learn go which might be more server/data oriented).

Anyway here is my first golang project. I must thank [Francesc Campoy](https://github.com/campoy) for his great [Youtube Videos JustForFunc](https://youtu.be/aYkxFbd6luY?list=PL64wiCrrxh4Jisi7OcCJIUpguV_f5jGnZ).

As you will see, **this is still a work in progress**. If some of you have infos about the build part, please share them via the [issues](https://github.com/topheman/gopher-ball/issues), [twitter](https://twitter.com/topheman) or any other way ...

## Install

You need to install first SDL2 and the SDL2 bindings for Go.

### Install SDL2

You need to install first SDL2 and the SDL2 bindings for Go. To do so, follow the instructions [here](https://github.com/veandco/go-sdl2).
It is quite easy to install on basically any platform.

You will also need to install [pkg-config](https://en.wikipedia.org/wiki/Pkg-config).

Example (on Mac OS X - with `pkg-config` already present):

```shell
$ which pkg-config
/usr/local/bin/pkg-config
$ brew install sdl2
$ brew install sdl2_image
$ brew install sdl2_ttf
$ brew install sdl2_mixer
$ go get -v github.com/veandco/go-sdl2/sdl
$ go get -v github.com/veandco/go-sdl2/mix
$ go get -v github.com/veandco/go-sdl2/img
$ go get -v github.com/veandco/go-sdl2/ttf
```

## Development

You will run two different commands:

* [`./build.sh`](https://github.com/topheman/gopher-ball/blob/master/build.sh): generates a `gopher-ball` and a `gopher-ball.app` via `go build` script with proper flags - **There might be a better way**.
* Then run `./gopher-ball` to start the app.
* If you are on Mac OS X, you can open `./gopher-ball.app` from the finder (this isn't a real packaged app, I'm still working on that)

## Build - Help wanted

If you install sdl2 and the go binding for sdl2, you can create a binary to test the game. I started go 3 weeks ago and this game last week, I still haven't figured how to properly package such an app.

How would you package this kind of app:

- cross platforms (for Mac OS X, Windows and Linux) - for starters, on Mac OS X
- embedding the assets
- statically linking the c++ library sdl2

## Credits

- assets/imgs/wood-background.png - [source](https://fr.vecteezy.com/art-vectoriel/133727-vector-wood-planks-background) - copyright by [carterart](https://fr.vecteezy.com/membres/carterart)
- assets/imgs/gopher.png - [The Go gopher was designed by Renee French](http://reneefrench.blogspot.com/) / [gopher.png was created by Takuya Ueda](https://twitter.com/tenntenn) - [source](https://github.com/golang-samples)
- assets/fonts/UbuntuMono-B.ttf - [from fontsquirrel.com]https://www.fontsquirrel.com/fonts/ubuntu-mono() - [under open license](http://font.ubuntu.com/ufl/)