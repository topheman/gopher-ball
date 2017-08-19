gopher-ball
===========

<table>
    <tr>
        <td>
            <a href="https://github.com/topheman/gopher-ball/releases" style="text-decoration:none;">
                <img src="https://raw.githubusercontent.com/topheman/gopher-ball/master/assets/originals/icon.png" width="75" />
                <strong>DOWNLOAD DEMOS</strong> in the releases section
            </a>
            <br />
            Choose your build by platform - currently only supporting MacOS (darwin).
        </td>
        <td style="width:30%; text-align: right;">
            <a href="http://i.imgur.com/Y1bT6Du.gif">
                <img src="http://i.imgur.com/G064PZD.gif">
            </a>
        </td>
    </tr>
</table>

## Goal

I have a JavaScript background (both frontend and NodeJS). I started go a few weeks ago (really enjoy it - just like a NodeJS with pointers and threads 😜).

I needed some project to test on. For the last years, I made a few [video games in JavaScript](http://dev.topheman.com/my-projects/) and I think it's a good way to learn a new programming language (some people will tell that it's a little weird to learn go that way which might be more server/data oriented).

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

You can build a dev version of the game via a simple `go build`.

You can also create the same build via `make darwin-dev` (same as `go build`, though it will name the binary `gopher-ball.app`, so that when you open it from the finder it doesn't open a terminal first).

## Build

### TLDR;

This part is still in progress (for the moment, only MacOS packaging is supported).

* `make darwin`: will create a bundle for MacOS, in the `./dist` folder
* `make clean`: cleans up the `./dist` folder
* `make` (alias of `make help`): prints out the documentation of the [Makefile](https://github.com/topheman/gopher-ball/blob/master/Makefile) 

### Explanation

There were a lot of apps made in golang with sdl2 (or other golang bridge with c) but none of them implement a **release step** (generate a standalone binary that you could share).

Since, there are C libraries involved, it implies that you link them in some way in the bundle you will generate.

Here is my solution (please share yours):

#### On MacOS:

I create a bundle with the same folder structure as any MacOS `.app`. Then I copy copy the shared libraries of SDL2 to `Contents/Frameworks` and link them via `install_name_tool`.

Some infos:

* To list the libs used by your binary: `otool -L <binary_name>` (same as Linux's `ldd`)
* To link those libraries, use `install_name_tool -change <lib_name> @executable_path/../Frameworks/<lib_name> <binary_name>`

Checkout the [Makefile](https://github.com/topheman/gopher-ball/blob/master/Makefile) for the whole build steps.

Thanks to [veeableful](https://github.com/veeableful) for her help [on this issue](https://github.com/veandco/go-sdl2/issues/234).

## Credits

- assets/imgs/wood-background.png - [source](https://fr.vecteezy.com/art-vectoriel/133727-vector-wood-planks-background) - copyright by [carterart](https://fr.vecteezy.com/membres/carterart)
- assets/imgs/gopher.png - [The Go gopher was designed by Renee French](http://reneefrench.blogspot.com/) / [gopher.png was created by Takuya Ueda](https://twitter.com/tenntenn) - [source](https://github.com/golang-samples)
- assets/fonts/UbuntuMono-B.ttf - [from fontsquirrel.com](https://www.fontsquirrel.com/fonts/ubuntu-mono) - [under open license](http://font.ubuntu.com/ufl/)

## Preview

[![Preview](http://i.imgur.com/Y1bT6Du.gif)](http://i.imgur.com/Y1bT6Du.gif)

[![Preview](https://raw.githubusercontent.com/topheman/gopher-ball/master/assets/imgs/splashScreen.jpg)](http://i.imgur.com/Y1bT6Du.gif)