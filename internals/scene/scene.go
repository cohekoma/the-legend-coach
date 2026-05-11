package scene

import (
	"image/color"

	"github.com/cohekoma/the-legend-coach/internals/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenInbox int = iota
	ScreenSquad
	ScreenTactics
	ScreenSchedule
	ScreenClub
)

func DrawSidebar(screen *ebiten.Image, currentScene int, menuItems []ui.MenuItem) {
	ebitenutil.DrawRect(screen, 0, 0, 220, 1800, color.RGBA{R: 28, G: 28, B: 36, A: 255})

	ebitenutil.DebugPrintAt(screen, "THE LEGEND COACH", 20, 30)

	for _, item := range menuItems {
		buttonColor := color.RGBA{R: 45, G: 45, B: 55, A: 255}

		if currentScene == item.Target {
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

func DrawContent(screen *ebiten.Image, currentScreen int) {
	ebitenutil.DrawRect(screen, 240, 40, 680, 460, color.RGBA{R: 34, G: 34, B: 42, A: 255})

	switch currentScreen {
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
