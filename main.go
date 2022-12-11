package main

import (
	"log"

	"github.com/abemaru/ebiten_demo/snake"
	"github.com/hajimehoshi/ebiten/v2"
)


func main() {
	game := snake.NewGame()

	ebiten.SetWindowSize(600, 600)
	ebiten.SetWindowTitle("Snake demo")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}


