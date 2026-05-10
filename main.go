package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	ScreenWidth  = 1280
	ScreenHeight = 800
)

type Game struct {
	currentScreen Screen
	menuItems     []MenuItem
}

func CreateNewGame() *Game {
	return &Game{
		currentScreen: ScreenInbox,
		menuItems: []MenuItem{
			{Label: "Inbox", Target: ScreenInbox, X: 20, Y: 80, W: 160, H: 36},
			{Label: "Squad", Target: ScreenSquad, X: 20, Y: 126, W: 160, H: 36},
			{Label: "Tactics", Target: ScreenTactics, X: 20, Y: 172, W: 160, H: 36},
			{Label: "Schedule", Target: ScreenSchedule, X: 20, Y: 218, W: 160, H: 36},
			{Label: "Club", Target: ScreenClub, X: 20, Y: 264, W: 160, H: 36},
		},
	}
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mouseX, mouseY := ebiten.CursorPosition()

		for _, item := range g.menuItems {
			if item.Contains(mouseX, mouseY) {
				g.currentScreen = item.Target
				break
			}
		}
	}

	return nil
}

func (gameState *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 20, 24, 255})
	drawSidebar(screen, gameState)
	drawContent(screen, gameState)
}

func (g *Game) Layout(w, h int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("The Legend Coach")

	if err := ebiten.RunGame(CreateNewGame()); err != nil {
		log.Fatal("Error")
	}
}
