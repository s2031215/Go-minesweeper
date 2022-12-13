# Gominesweeper - Go Console minesweeper game

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/s2031215/Go-minesweeper?style=for-the-badge)

A simple minesweeper game build with Golang

## Screenshots
![Screenshot](Screenshot.png?raw=true "Screenshot")

## Description

The minesweeper game can play in command prompt, the library add gocui for the UI control.
It suggest run in unicode command prompt but ASCII also will work.

## Getting Started

### Dependencies

OS:
* Windows 10 (tested)
* Linux(ARM) Armbian 5.4.2 (tested)
* Linxu(x86) to-do

libraries:
* go 1.16
* gocui v0.5.0

### Binary Install
- Download the executable files in [Releases](https://github.com/s2031215/Go-minesweeper/releases)
- Run and Fun!
### Executing program from sourse code

```sh
git clone https://github.com/s2031215/Gominesweeper.git
cd Gominesweeper
go run . # start game
#or
go build . # build executable file
./Gominesweeper #linux
Gominesweeper.exe #Windows
```

### install as modules 
```sh
#set go path
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
go install github.com/s2031215/Gominesweeper@latest
#run program
Gominesweeper
```

## Authors

s2031215

## Version History

* 1.0
    * Initial Release
* 1.1
    * Add Linux(Arm64) build support
    
## Roadmap

- [x] Add Flag Function
- [x] Add Status table
- [ ] Add Difficulty Level 10x10 20x20
- [ ] Add Item can help when play (one more life/auto detect mines)

## Acknowledgments

* [gocui](https://github.com/jroimartin/gocui)
