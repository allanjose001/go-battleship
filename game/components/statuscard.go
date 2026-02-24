package components

import (
	"fmt"
	"image/color"

	"github.com/allanjose001/go-battleship/game/components/basic"
	"github.com/allanjose001/go-battleship/game/components/basic/colors"
	"github.com/hajimehoshi/ebiten/v2"
)

// StatusCard representa um painel de estatísticas composto por vários widgets internos.
// Ele encapsula layout, tamanho e renderização de um bloco de stats.
// PS: Pedi pro chatgpt documentar
type StatusCard struct {
	pos        basic.Point    // posição base lógica do componente
	currentPos basic.Point    // posição final após aplicação de offsets
	size       basic.Size     // tamanho total do container
	body       StylableWidget // widget raiz que desenha todo o conteúdo
}

// SetPos define a posição base do StatusCard.
func (s *StatusCard) SetPos(point basic.Point) {
	s.pos = point
}

// NewStatCard constrói um StatusCard completo.
// O layout muda dependendo de omnde a tela for usada (ranking x profile).
// - isRanking: ativa versão compacta
// - playerName: nome exibido no topo do card
func NewStatCard(
	pos basic.Point,
	screenSize basic.Size,
	matches, wins, score, higherHitSequence int,
	winrate, mediumHitRate float32,
	isRanking bool,
	playerName string,
	rankingPosition int,
) *StatusCard {

	contSize, cardSize := switchSizes(isRanking, screenSize)

	// Gera lista de widgets de estatísticas conforme modo
	statsList := initWidgets(matches, wins, winrate, mediumHitRate, score, higherHitSequence, isRanking, cardSize)

	var titleWidget Widget

	if isRanking {
		var circleColor color.Color
		switch rankingPosition {
		case 1:
			circleColor = color.RGBA{R: 255, G: 215, B: 0, A: 255} // Dourado
		case 2:
			circleColor = color.RGBA{R: 192, G: 192, B: 192, A: 255} // Prata
		case 3:
			circleColor = color.RGBA{R: 205, G: 127, B: 50, A: 255} // Bronze
		default:
			circleColor = colors.White // Branco para os demais
		}

		rankText := NewText(
			basic.Point{},
			fmt.Sprintf("%d°", rankingPosition),
			colors.Black,
			20,
		)

		rankCircle := NewContainer(
			basic.Point{},
			basic.Size{W: 40, H: 40},
			20,
			circleColor,
			basic.Center,
			basic.Center,
			rankText,
		)

		nameText := NewText(
			basic.Point{},
			playerName,
			colors.White,
			30,
		)

		titleWidget = NewRow(
			basic.Point{},
			15,
			basic.Size{W: contSize.W, H: 40},
			basic.Center,
			basic.Center,
			[]Widget{rankCircle, nameText},
		)
	} else {
		titleWidget = NewText(
			basic.Point{},
			playerName,
			colors.White,
			30,
		)
	}

	return &StatusCard{
		pos:        pos,
		currentPos: pos,
		size:       contSize,
		body: NewContainer( //pai de todos
			basic.Point{},
			contSize, 12,
			colors.Dark,
			basic.Start, //não precisa mexer aqui pois os filhos são Layout
			basic.Start, //, não há como fazer sobrecarga de metodo
			NewColumn(
				basic.Point{}, 10, contSize,
				basic.Center, basic.Center,
				[]Widget{
					// Nome do jogador no topo
					titleWidget,
					// Container intermediário que organiza a Row de stats
					NewContainer(
						basic.Point{},
						basic.Size{
							W: contSize.W,
							H: cardSize.H,
						}, 0,
						nil,
						basic.Start,
						basic.Start,
						NewRow(
							basic.Point{}, 20,
							basic.Size{
								W: contSize.W,
								H: cardSize.H,
							},
							basic.Center,
							basic.Center,
							statsList,
						),
					),
				},
			),
		),
	}
}

// switchSizes decide layout que o componente vai assumir
func switchSizes(isRanking bool, screenSize basic.Size) (basic.Size, basic.Size) {
	var contSize, cardSize basic.Size

	// Define tamanhos diferentes para ranking (compacto) vs tela de profile
	if isRanking {
		contSize = basic.Size{
			W: 0.9 * screenSize.W,
			H: 120,
		}
		cardSize = basic.Size{W: 300, H: 60}
	} else {
		contSize = basic.Size{
			W: 0.9 * screenSize.W,
			H: 220,
		}
		cardSize = basic.Size{W: 210, H: 100}
	}
	return contSize, cardSize
}

// GetPos retorna a posição base do StatusCard.
func (s *StatusCard) GetPos() basic.Point {
	return s.pos
}

// GetSize retorna o tamanho total do StatusCard.
func (s *StatusCard) GetSize() basic.Size {
	return s.size
}

// Update recalcula posição final com offset e propaga para o corpo.
func (s *StatusCard) Update(offset basic.Point) {
	s.currentPos = s.pos.Add(offset)
	s.body.Update(s.currentPos)
}

// Draw delega o desenho para o widget interno.
func (s *StatusCard) Draw(screen *ebiten.Image) {
	s.body.Draw(screen)
}

// initWidgets cria os cartões individuais de estatística.
// Retorna versão completa ou compacta dependendo do uso.
func initWidgets(matches int, wins int, winrate float32, mediumHitRate float32, score int, higherHitSequence int, ranking bool, size basic.Size) []Widget {
	// versão compacta para ranking
	if ranking {
		return []Widget{
			createStatCard("Pontuação Total", fmt.Sprintf("%d", score), size),
			createStatCard("Vitórias", fmt.Sprintf("%d", wins), size),
			createStatCard("Maior Sequência de Acertos", fmt.Sprintf("%d", higherHitSequence), size),
		}
	}

	// versão completa
	return []Widget{
		createStatCard("Partidas", fmt.Sprintf("%d", matches), size),
		createStatCard("Vitórias", fmt.Sprintf("%d", wins), size),
		createStatCard("% de Vitória", fmt.Sprintf("%.2f", winrate)+" %", size),
		createStatCard("% de Acertos Média", fmt.Sprintf("%.2f", mediumHitRate)+" %", size),
		createStatCard("Maior Score", fmt.Sprintf("%d", score), size),
	}
}

// createStatCard cria um card visual individual com label + valor.
// Encapsula layout interno (texto + container branco).
func createStatCard(label, value string, size basic.Size) *Container {
	labelTxt := NewText(basic.Point{}, label, colors.Black, 20)
	valueTxt := NewText(basic.Point{}, value, colors.Black, 25)

	// coluna centralizada com label em cima e valor embaixo
	content := NewColumn(
		basic.Point{X: 0, Y: 0}, 5,
		size,
		basic.Center,
		basic.Center,
		[]Widget{labelTxt, valueTxt},
	)

	// container visual do card
	return NewContainer(
		basic.Point{},
		size,
		12,
		colors.White,
		basic.Center,
		basic.Center,
		content,
	)
}
