package scenes

import (
	"github.com/allanjose001/go-battleship/game/components"
	"github.com/allanjose001/go-battleship/game/components/basic"
	"github.com/allanjose001/go-battleship/game/components/basic/colors"
	"github.com/allanjose001/go-battleship/game/state"
	"github.com/allanjose001/go-battleship/internal/medal"
	"github.com/hajimehoshi/ebiten/v2"
)

// ProfileScene representa a tela de perfil do jogador.
type ProfileScene struct {
	state *state.GameState
	root  *components.Column // O container pai que envolve toda a cena.
	StackHandler
}

// init Função que inicializa componentes
func (p *ProfileScene) init(size basic.Size) {
	playerName := p.ctx.Profile.Username

	// Chamamos o método agora vinculado à struct
	medals := p.loadMedals()

	// Coluna principal que centraliza verticalmente
	p.root = components.NewColumn(
		basic.Point{},
		40,
		size,
		basic.Start,
		basic.Center,
		[]components.Widget{
			// Title
			components.NewText(basic.Point{},
				"PERFIL DE JOGADOR",
				colors.White,
				42),

			// Container com Row para estatisticas
			components.NewStatCard(
				basic.Point{},
				size,
				&p.ctx.Profile.Stats,
				false,
				playerName,
				0,
			),

			// Título da seção de medalhas
			components.NewText(basic.Point{}, "MURAL DE MEDALHAS", colors.White, 28),

			// Container com Row para medalhas reais
			components.NewContainer(
				basic.Point{},
				basic.Size{W: 750, H: 100},
				0, nil,
				basic.Center, basic.Center,
				components.NewRow(
					basic.Point{},
					40,
					basic.Size{W: 750, H: 100},
					basic.Center, basic.Center,
					*medals,
				),
			),

			// Botão Voltar
			components.NewButton(
				basic.Point{},
				basic.Size{W: 220, H: 55},
				"Voltar",
				colors.Dark,
				colors.White,
				func(b *components.Button) {
					p.stack.Pop()
				},
			),
		},
	)
}

// loadMedals agora é um método de ProfileScene para acessar p.ctx.Profile.Stats
func (p *ProfileScene) loadMedals() *[]components.Widget {
	var widgets = []components.Widget{}

	// Iteramos sobre a MedalsList oficial do seu medal_repository.go
	for _, m := range medal.MedalsList {

		// Lógica pedida pelo Cauã:
		// 1. Verifica se o player atingiu os requisitos
		isUnlocked := m.Verification(p.ctx.Profile.Stats)

		displayIcon := m.GrayIconPath // Ícone cinza por padrão
		displayTitle := "BLOQUEADA"
		displayDesc := "???" // Descrição oculta

		// 2. Se conquistou, libera as informações reais
		if isUnlocked {
			displayIcon = m.IconPath
			displayTitle = m.Name
			displayDesc = m.Description
		}

		// 3. Adiciona o componente visual real
		widgets = append(widgets, components.NewMedal(
			displayIcon,
			displayTitle,
			displayDesc,
			basic.Size{W: 230, H: 90},
		))
	}

	return &widgets
}

// Implementações do contrato Scene
func (p *ProfileScene) OnEnter(prev Scene, size basic.Size) {
	p.init(size)
}

func (p *ProfileScene) OnExit(next Scene) {}

func (p *ProfileScene) Update() error {
	p.root.Update(basic.Point{X: 0, Y: 0})
	return nil
}

func (p *ProfileScene) Draw(screen *ebiten.Image) {
	p.root.Draw(screen)
}
