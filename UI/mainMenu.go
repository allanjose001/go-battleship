package UI

import (
	"github.com/allanjose001/go-battleship/game/ai"
	"github.com/allanjose001/go-battleship/game/scenes"
	"github.com/hajimehoshi/ebiten/v2"
)

type App struct {
	current Screen
}

func NewApp() *App {
	a := &App{}
	// Inicia no menu de dificuldade
	a.current = scenes.NewDifficultyMenu(ScreenWidth, ScreenHeight, func(p *ai.AIPlayer) {
		a.StartGame(p)
	})
	return a
}

func (a *App) StartGame(p *ai.AIPlayer) {
	// Transição para o tabuleiro de batalha
	a.current = NewDualBoardUI(10, 10)
}

func (a *App) Update() error              { return a.current.Update() }
func (a *App) Draw(screen *ebiten.Image)  { a.current.Draw(screen) }
func (a *App) Layout(w, h int) (int, int) { return ScreenWidth, ScreenHeight }

func Run() error {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Battleship - Fúria dos Mares")
	return ebiten.RunGame(NewApp())
}
