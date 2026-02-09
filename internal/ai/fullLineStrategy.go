package ai

import "github.com/allanjose001/go-battleship/internal/entity"

type FullLineStrategy struct {}

func (s *FullLineStrategy) TryAttack(ai *AIPlayer, board *entity.Board) bool {
	if len(ai.priorityQueue) == 0 {
		return false
	}
	x, y := ai.PopPriority()
	ship := board.AttackPositionB(x, y)
	ai.AdjustStrategy(board, x, y, ship)
	if ship == nil || ship.IsDestroyed() {
		return true
	}
	// Se acertou mas não destruiu, adiciona os próximos da linha na fila de prioridade

	horizontal := false
	vertical := false

    if ai.IsValidForTesting(x, y-1) && ai.virtualBoard[x][y-1] == 2 { horizontal = true }
    if ai.IsValidForTesting(x, y+1) && ai.virtualBoard[x][y+1] == 2 { horizontal = true }
    if ai.IsValidForTesting(x-1, y) && ai.virtualBoard[x-1][y] == 2 { vertical = true }
    if ai.IsValidForTesting(x+1, y) && ai.virtualBoard[x+1][y] == 2 { vertical = true }

	ai.ClearPriorityQueue()

    if horizontal {
        // varre à esquerda
        c := y // c de coluna(col)
        for {
            ny := c - 1
            if !ai.IsValidForTesting(x, ny) || ai.virtualBoard[x][ny] != 0 { break }
            ai.AddToPriorityQueue(x, ny)
            c = ny
        }
        // varre à direita
        c = y
        for {
            ny := c + 1
            if !ai.IsValidForTesting(x, ny) || ai.virtualBoard[x][ny] != 0 { break }
            ai.AddToPriorityQueue(x, ny)
            c = ny
        }
        ai.StartChase()
    } else if vertical {
        r := x // r de linha(row)
        for {
            nx := r - 1
            if !ai.IsValidForTesting(nx, y) || ai.virtualBoard[nx][y] != 0 { break }
            ai.AddToPriorityQueue(nx, y)
            r = nx
        }
        r = x
        for {
            nx := r + 1
            if !ai.IsValidForTesting(nx, y) || ai.virtualBoard[nx][y] != 0 { break }
            ai.AddToPriorityQueue(nx, y)
            r = nx
        }
        ai.StartChase()
    } else {
        ai.AttackNeighbors(x, y)
    }
    return true
 

}