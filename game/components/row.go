package components

import (
	"github.com/allanjose001/go-battleship/game/components/basic"
	"github.com/hajimehoshi/ebiten/v2"
)

// Row organiza widgets em uma linha, com espaçamento e altura opcional fixa
// <USAR A ALTURA DO PAI PARA ALINHAMENTO NO EIXO SECUNDARIO (Y)>
type Row struct {
	Pos        basic.Point //posição inicial
	Spacing    float32     //espaçamento horizontal entre elementos
	Children   []Widget
	MainAlign  basic.Align // alinhamento dos elementos no eixo principal (center, start, end)
	CrossAlign basic.Align //eixo cruzado
	size       basic.Size  //para calculo de tamanho caso necessario
	currentPos basic.Point
}

// NewRow cria uma linha e já calcula a posição de todos os widgets,
// alinhando verticalmente e no eixo secundario de acordo com o alinhamento dado
// construtor
func NewRow(
	pos basic.Point,
	spacing float32,
	parentSize basic.Size,
	mainAlign basic.Align,
	crossAlign basic.Align,
	children []Widget,
) *Row {

	r := &Row{
		Pos:        pos,
		Spacing:    spacing,
		Children:   children,
		MainAlign:  mainAlign,
		CrossAlign: crossAlign,
	}

	// 1) posiciona como Start/Start
	r.init()

	// 2) se ambos Start, não faz nada
	if mainAlign == basic.Start && crossAlign == basic.Start {
		return r
	}

	// 3) aplica alinhamentos relativos ao retângulo do pai iniciado em r.Pos
	if mainAlign != basic.Start {
		r.alignMain(parentSize)
	}
	if crossAlign != basic.Start {
		r.alignCross(parentSize)
	}

	return r
}

// posiciona filhos como Start/Start
func (r *Row) init() {
	cursorX := float32(0) // posição local relativa

	for _, w := range r.Children {
		size := w.GetSize()
		w.SetPos(basic.Point{
			X: cursorX,
			Y: 0, // local Y relativo
		})
		cursorX += size.W + r.Spacing
	}
	r.calcSize()
}

// Update chama Update de todos os filhos
func (r *Row) Update(offset basic.Point) {
	r.currentPos = r.Pos.Add(offset)
	for _, w := range r.Children {
		w.Update(r.currentPos)
	}
}

// Draw chama Draw de todos os filhos fazendo recursão com posição na arvore
func (r *Row) Draw(screen *ebiten.Image) {
	for _, w := range r.Children {
		w.Draw(screen)
	}
}

// alinhamento no eixo principal (horizontal)
func (r *Row) alignMain(parentSize basic.Size) {
	content := r.GetSize()

	var offsetX float32
	switch r.MainAlign {
	case basic.Start:
		return
	case basic.Center:
		offsetX = (parentSize.W - content.W) / 2
	case basic.End:
		offsetX = parentSize.W - content.W
	}

	for _, w := range r.Children {
		p := w.GetPos()
		// mantém o deslocamento relativo original dentro da Row e adiciona o offset calculado
		p.X = p.X + offsetX
		w.SetPos(p)
	}
}

// alinhamento no eixo cruzado (vertical)
func (r *Row) alignCross(parentSize basic.Size) {
	for _, w := range r.Children {
		size := w.GetSize()
		p := w.GetPos()

		switch r.CrossAlign {
		case basic.Start:
			p.Y = 0
		case basic.Center:
			// centraliza verticalmente relativo à Row
			p.Y = (parentSize.H - size.H) / 2
		case basic.End:
			p.Y = parentSize.H - size.H
		}

		w.SetPos(p)
	}
}

// tamanho ocupado pela row, calculado no init
func (r *Row) calcSize() {
	var totalW float32
	var maxH float32

	for i, w := range r.Children {
		s := w.GetSize()

		if i > 0 {
			totalW += r.Spacing
		}
		totalW += s.W

		if s.H > maxH {
			maxH = s.H
		}
	}

	r.size = basic.Size{W: totalW, H: maxH}
}

func (r *Row) GetPos() basic.Point {
	return r.Pos
}

func (r *Row) GetSize() basic.Size {
	return r.size
}

func (r *Row) SetPos(basic.Point) {
}
