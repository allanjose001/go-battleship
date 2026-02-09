package ai

import "github.com/allanjose001/go-battleship/internal/entity"

type DiscoveryStrategy struct{}

func (s *DiscoveryStrategy) TryAttack(ai *AIPlayer, board *entity.Board) bool {
	if ai.IsChasing() {
		return false
	}
	if len(ai.priorityQueue) == 0 {
		return false
	}

	// Pega a primeira posição da fila de prioridade
	x, y := ai.PopPriority()
	ship := board.AttackPositionB(x, y)
	ai.AdjustStrategy(board, x, y, ship)
	return true
}
