package scenes

import (
	"log"

	"github.com/allanjose001/go-battleship/game/components"
	"github.com/allanjose001/go-battleship/game/components/basic"
	"github.com/allanjose001/go-battleship/game/components/basic/colors"
	"github.com/allanjose001/go-battleship/internal/service"
	"github.com/hajimehoshi/ebiten/v2"
)

type RankingScene struct {
	layout       components.Widget
	currentPage  int
}

func (m *RankingScene) OnExit(_ Scene) {}

func (m *RankingScene) OnEnter(_ Scene, screenSize basic.Size) {
	err := m.init(screenSize)
	
	if err != nil {
		log.Fatal("Erro ao carregar componentes na tela inicial: ", err)
	}
}

func (m *RankingScene) Update() error {
	if m.layout != nil {
		m.layout.Update(basic.Point{X: 0, Y: 0})
	}
	return nil
}

func (m *RankingScene) Draw(screen *ebiten.Image) {
	if m.layout != nil {
		m.layout.Draw(screen)
	}
}

func (m *RankingScene) init(screenSize basic.Size) error {
	
	itemsPerPage := 3
	start := m.currentPage * itemsPerPage
	end := start + itemsPerPage

	allPlayers := service.GetTopScores(9)

	if start > len(allPlayers) {
		start = len(allPlayers)
	}
	if end > len(allPlayers) {
		end = len(allPlayers)
	}

	pagePlayers := allPlayers[start:end]

	var sceneWidgets []components.Widget

	sceneWidgets = append(
		sceneWidgets, 
		components.NewText(
			basic.Point{},
			"Ranking",
			colors.White,
			42,
		),
	)

	for i, player := range pagePlayers {
		card := components.NewStatCard(
			basic.Point{},
			screenSize,
			player.Stats.Matches,
			player.Stats.Wins,
			player.Stats.TotalScore,
			player.Stats.HigherHitSequence,
			player.Stats.WinRate(),
			player.Stats.Accuracy(),
			true,
			player.Username,
			start + i + 1,
		)
		sceneWidgets = append(sceneWidgets, card)
	}

	var paginationButtons []components.Widget

	if m.currentPage > 0 {
		paginationButtons = append(
			paginationButtons, 
			components.NewButton(
				basic.Point{},
				basic.Size{W: 150, H: 40},
				"Anterior",
				colors.Dark,
				nil,
				func(bt *components.Button) {
					m.currentPage--
					m.init(screenSize)
				},
			),
		)
	}

	if end < len(allPlayers) {
		paginationButtons = append(
			paginationButtons, 
			components.NewButton(
				basic.Point{},
				basic.Size{W: 150, H: 40},
				"Próximo",
				colors.Dark,
				nil,
				func(bt *components.Button) {
					m.currentPage++
					m.init(screenSize)
				},
			),
		)
	}

	if len(paginationButtons) > 0 {
		pagRow := components.NewRow(
			basic.Point{},
			20,
			basic.Size{},
			basic.Start,
			basic.Start,
			paginationButtons,
		)

		pagContainer := components.NewContainer(
			basic.Point{},
			pagRow.GetSize(),
			0,
			nil,
			basic.Center,
			basic.Center,
			pagRow,
		)

		sceneWidgets = append(sceneWidgets, pagContainer)
	}

	sceneWidgets = append(
		sceneWidgets,
		components.NewButton(
			basic.Point{},
			basic.Size{W: 400, H: 50},
			"Voltar",
			colors.Dark,
			nil,
			func(bt *components.Button) {
				SwitchTo(&HomeScreen{})
			},
		),
	)

	m.layout = components.NewColumn(
		basic.Point{},
		15,
		basic.Size{W: screenSize.W, H: screenSize.H},
		basic.Center,
		basic.Center,
		sceneWidgets,
	)

	return nil
}