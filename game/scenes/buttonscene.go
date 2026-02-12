package scenes

import (
	"github.com/allanjose001/go-battleship/game/components"
	"github.com/allanjose001/go-battleship/game/components/basic"
	"github.com/allanjose001/go-battleship/game/components/basic/colors"
	"github.com/hajimehoshi/ebiten/v2"
)

// ButtonScene esse teste adiciona o botão voltar á trela, porem infinitamente rsrsrs, clickar em voltar remove ele da group
// (remove e add não devem ter mtos usos, foi só para mostrar o botão funcionando)
type ButtonScene struct { //teste porco mas ok
	group components.Widget
}

func (b *ButtonScene) OnEnter(_ Scene, screenSize basic.Size) {

	b.group = components.NewRow(
		basic.Point{},
		100,
		screenSize,
		basic.Center,
		basic.Center,
		[]components.Widget{
			components.NewButton(
				basic.Point{},
				basic.Size{W: 200.0, H: 70.0},
				"pressione",
				colors.Blue,
				nil,
				func(bt *components.Button) {
				},
			),
		},
	)
}

func (b *ButtonScene) OnExit(_ Scene) {

}

func (b *ButtonScene) Update() error {
	b.group.Update(basic.Point{})
	return nil
}

func (b *ButtonScene) Draw(screen *ebiten.Image) {
	b.group.Draw(screen)
}
