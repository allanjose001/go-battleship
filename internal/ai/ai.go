package ai

import ( "github.com/allanjose001/go-battleship/internal/entity" )

type AIPlayer struct { 
	virtualBoard [10][10]int
	priorityQueue []Pair
	Strategies []Strategy
	enemyFleet *entity.Fleet
}

func (ai *AIPlayer) Attack(enemyBoard *entity.Board) {
	for _, strat := range ai.Strategies { // verifica estrategias disponiveis
		if strat.TryAttack(enemyBoard) {
			return;
		}
	}
}

// retorna o tamanho do proximo navio da enemyFleet
func (ai *AIPlayer) SizeOfNextShip() int {
	for _, ship := range ai.enemyFleet.Ships {
		if ship != nil && !ship.IsDestroyed() {
			return ship.Size
		}
	}

	return 0
}

func (ai *AIPlayer) FleetShipDestroyed(size int) {
	for _, ship := range ai.enemyFleet.Ships {
		if ship != nil && ship.Size == size && !ship.IsDestroyed() {
			ship.HitCount = ship.Size
			return;
		}
	}
}