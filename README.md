# Gominesweeper

A simple minesweeper game build with Golang
![Screenshot](Screenshot.png?raw=true "Screenshot")


## Description

The minesweeper game can play in command prompt, the library add gocui for the UI control.
It suggest run in unicode command prompt but ASCII also will work.

## Getting Started

### Dependencies

OS:
* Windows 10 (tested)
* Linux (not test but should work)

libraries:
* go 1.16
* gocui v0.5.0

### Installing

```sh
git clone https://github.com/s2031215/Gominesweeper.git
```
### Executing program

```sh
cd Gominesweeper
go run . # start game
go build . # build executable file
```

## Authors

s2031215

## Version History

* 1.0
    * Initial Release
    
## Roadmap

- [x] Add Flag Function
- [x] Add Status table
- [ ] Add Difficulty Level 10x10 20x20
- [ ] Add Item can help when play (one more live/auto detect mines)

## Acknowledgments

Inspiration, code snippets, etc.
* [gocui](https://github.com/jroimartin/gocui)
