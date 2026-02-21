package scenes

import (
	"github.com/allanjose001/go-battleship/UI"
	"github.com/allanjose001/go-battleship/game/components"
	"github.com/allanjose001/go-battleship/game/components/basic"
	"github.com/allanjose001/go-battleship/game/components/basic/colors"
	"github.com/hajimehoshi/ebiten/v2"
)

type DifficultyMenu struct {
	layout components.Widget
}

func NewDifficultyMenu() *DifficultyMenu {
	// Tamanho dos botões baseado no seu rascunho
	btnSize := basic.Size{W: 220, H: 60}

	// Criando os botões com a nomenclatura intuitiva
	btnRecruta := components.NewButton(basic.Point{}, btnSize, "Recruta", colors.Blue, colors.White, func(b *components.Button) {
		// Callback para nível Fácil
	})

	btnImediato := components.NewButton(basic.Point{}, btnSize, "Imediato", colors.Blue, colors.White, func(b *components.Button) {
		// Callback para nível Médio
	})

	btnAlmirante := components.NewButton(basic.Point{}, btnSize, "Almirante", colors.Blue, colors.White, func(b *components.Button) {
		// Callback para nível Difícil
	})

	// Referenciando as constantes exportadas (ScreenWidth/ScreenHeight) do pacote UI
	screenSize := basic.Size{W: float32(UI.ScreenWidth), H: float32(UI.ScreenHeight)}

	// Organização vertical centralizada
	column := components.NewColumn(
		basic.Point{X: 0, Y: 0},
		25, // Spacing
		screenSize,
		basic.Center,
		basic.Center,
		[]components.Widget{
			components.NewText(basic.Point{}, "SELECIONE SUA PATENTE", colors.White, 28),
			btnRecruta,
			btnImediato,
			btnAlmirante,
		},
	)

	return &DifficultyMenu{layout: column}
}

func (m *DifficultyMenu) Update() error {
	m.layout.Update(basic.Point{X: 0, Y: 0})
	return nil
}

func (m *DifficultyMenu) Draw(screen *ebiten.Image) {
	screen.Fill(colors.DarkBlue)
	m.layout.Draw(screen)
}

func (m *DifficultyMenu) Layout(w, h int) (int, int) {
	return w, h
}
