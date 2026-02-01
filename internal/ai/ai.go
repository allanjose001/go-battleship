package ai

import ( "github.com/allanjose001/go-battleship/internal/entity" )

type AIPlayer struct { 
	Strategies []Strategy
}

func (ai *AIPlayer) Attack(enemyBoard *entity.Board) {
	for _, strat := range ai.Strategies { // verifica estrategias disponiveis
		if strat.TryAttack(enemyBoard) {
			return;
		}
	}
}