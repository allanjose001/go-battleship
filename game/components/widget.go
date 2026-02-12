package components

import (
	"github.com/allanjose001/go-battleship/game/components/basic"
	"github.com/hajimehoshi/ebiten/v2"
)

// Widget serve para facilitar posicionamento e tratamento de componentes
type Widget interface {

	// GetPos Retorna a posição e tamanho do widget
	GetPos() basic.Point
	//SetPos para atualizar quando necessário
	SetPos(basic.Point)
	GetSize() basic.Size

	//Update Atualiza o widget//seus filhos no padrão ebiten
	Update(offset basic.Point)

	// Draw Desenha o widget/seus filhos chamando Draw
	Draw(screen *ebiten.Image)
}
