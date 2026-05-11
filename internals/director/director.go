package director

import (
	"image/color"

	"github.com/cohekoma/the-legend-coach/internals/scene"
	"github.com/cohekoma/the-legend-coach/internals/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	GameScreenWidth  = 1280
	GameScreenHeight = 800
)

type Game struct {
	CurrentScreen int
	MenuItems     []ui.MenuItem
}

func CreateNewGame() *Game {
	return &Game{
		CurrentScreen: scene.ScreenInbox,
		MenuItems: []ui.MenuItem{
			{Label: "Inbox", Target: scene.ScreenInbox, X: 20, Y: 80, W: 160, H: 36},
			{Label: "Squad", Target: scene.ScreenSquad, X: 20, Y: 126, W: 160, H: 36},
			{Label: "Tactics", Target: scene.ScreenTactics, X: 20, Y: 172, W: 160, H: 36},
			{Label: "Schedule", Target: scene.ScreenSchedule, X: 20, Y: 218, W: 160, H: 36},
			{Label: "Club", Target: scene.ScreenClub, X: 20, Y: 264, W: 160, H: 36},
		},
	}
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mouseX, mouseY := ebiten.CursorPosition()

		for _, item := range g.MenuItems {
			if item.Contains(mouseX, mouseY) {
				g.CurrentScreen = item.Target
				break
			}
		}
	}

	return nil
}

func (gameState *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 20, 24, 255})
	scene.DrawSidebar(screen, int(gameState.CurrentScreen), gameState.MenuItems)
	scene.DrawContent(screen, gameState.CurrentScreen)
}

func (g *Game) Layout(w, h int) (int, int) {
	return GameScreenWidth, GameScreenHeight
}
