package scenes

import (
	"log"

	"github.com/allanjose001/go-battleship/game/components"
	"github.com/allanjose001/go-battleship/game/components/basic"
	"github.com/allanjose001/go-battleship/game/components/basic/colors"
	"github.com/hajimehoshi/ebiten/v2"
)

// DifficultyScene representa a tela de seleção de dificuldade da IA.
type DifficultyScene struct {
	layout components.Widget
	StackHandler
}

func (d *DifficultyScene) OnExit(_ Scene) {}

func (d *DifficultyScene) OnEnter(_ Scene, screenSize basic.Size) {
	err := d.init(screenSize)
	if err != nil {
		log.Fatal("Erro ao carregar componentes na tela de dificuldade: ", err)
	}
}

func (d *DifficultyScene) Update() error {
	if d.layout != nil {
		d.layout.Update(basic.Point{X: 0, Y: 0})
	}
	return nil
}

func (d *DifficultyScene) Draw(screen *ebiten.Image) {
	if d.layout != nil {
		d.layout.Draw(screen)
	}
}

// init Inicializa componentes da tela de seleção de dificuldade
func (d *DifficultyScene) init(screenSize basic.Size) error {
	d.layout = components.NewColumn(
		basic.Point{},
		30,
		basic.Size{W: screenSize.W, H: screenSize.H},
		basic.Center,
		basic.Center,
		[]components.Widget{
			components.NewText(
				basic.Point{},
				"SELECIONAR DIFICULDADE",
				colors.White,
				36,
			),

			components.NewButton(
				basic.Point{},
				basic.Size{W: 400, H: 60},
				"Fácil",
				colors.Dark,
				nil,
				func(bt *components.Button) {
					log.Println("Dificuldade selecionada: Fácil")
					//TODO: iniciar jogo com NewEasyAIPlayer
				},
			),

			components.NewButton(
				basic.Point{},
				basic.Size{W: 400, H: 60},
				"Médio",
				colors.Dark,
				nil,
				func(bt *components.Button) {
					log.Println("Dificuldade selecionada: Médio")
					//TODO: iniciar jogo com NewMediumAIPlayer
				},
			),

			components.NewButton(
				basic.Point{},
				basic.Size{W: 400, H: 60},
				"Difícil",
				colors.Dark,
				nil,
				func(bt *components.Button) {
					log.Println("Dificuldade selecionada: Difícil")
					//TODO: iniciar jogo com NewHardAIPlayer
				},
			),

			components.NewButton(
				basic.Point{},
				basic.Size{W: 400, H: 50},
				"Voltar",
				colors.Dark,
				nil,
				func(bt *components.Button) {
					d.stack.Pop()
				},
			),
		},
	)
	return nil
}
