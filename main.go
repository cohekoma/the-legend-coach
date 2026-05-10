package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	ScreenWidth  = 1280
	ScreenHeight = 800
)

type Screen int

const (
	ScreenInbox Screen = iota
	ScreenSquad
	ScreenTactics
	ScreenSchedule
	ScreenClub
)

type MenuItem struct {
	Label  string
	Target Screen
	X, Y   int
	W, H   int
}

func (m MenuItem) Contains(px, py int) bool {
	return px >= m.X &&
		px <= m.X+m.W &&
		py >= m.Y &&
		py <= m.Y+m.H
}

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

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 20, 24, 255})
	g.drawSidebar(screen)
	g.drawContent(screen)
}

func (g *Game) drawSidebar(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, 0, 0, 220, ScreenHeight, color.RGBA{R: 28, G: 28, B: 36, A: 255})

	ebitenutil.DebugPrintAt(screen, "THE LEGEND COACH", 20, 30)

	for _, item := range g.menuItems {
		buttonColor := color.RGBA{R: 45, G: 45, B: 55, A: 255}

		if g.currentScreen == item.Target {
			buttonColor = color.RGBA{R: 80, G: 80, B: 110, A: 255}
		}

		ebitenutil.DrawRect(
			screen,
			float64(item.X),
			float64(item.Y),
			float64(item.W),
			float64(item.H),
			buttonColor,
		)

		ebitenutil.DebugPrintAt(screen, item.Label, item.X+12, item.Y+10)
	}
}

func (g *Game) drawContent(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, 240, 40, 680, 460, color.RGBA{R: 34, G: 34, B: 42, A: 255})

	switch g.currentScreen {
	case ScreenInbox:
		ebitenutil.DebugPrintAt(screen, "Inbox", 270, 70)
		ebitenutil.DebugPrintAt(screen, "- Assistant Manager: Training report is ready.", 270, 120)
		ebitenutil.DebugPrintAt(screen, "- Scout: New striker recommendation available.", 270, 150)

	case ScreenSquad:
		ebitenutil.DebugPrintAt(screen, "Squad", 270, 70)
		ebitenutil.DebugPrintAt(screen, "GK  Daniel Ward", 270, 120)
		ebitenutil.DebugPrintAt(screen, "ST  Marco Silva", 270, 150)

	case ScreenTactics:
		ebitenutil.DebugPrintAt(screen, "Tactics", 270, 70)
		ebitenutil.DebugPrintAt(screen, "Formation: 4-3-3", 270, 120)
		ebitenutil.DebugPrintAt(screen, "Mentality: Balanced", 270, 150)

	case ScreenSchedule:
		ebitenutil.DebugPrintAt(screen, "Schedule", 270, 70)
		ebitenutil.DebugPrintAt(screen, "Next Match: Sunday vs Red City", 270, 120)

	case ScreenClub:
		ebitenutil.DebugPrintAt(screen, "Club", 270, 70)
		ebitenutil.DebugPrintAt(screen, "Board Confidence: Stable", 270, 120)
	}
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
