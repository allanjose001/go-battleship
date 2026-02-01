package ai

import ( "math/rand" 
		"github.com/allanjose001/go-battleship/internal/entity"
)

type RandomStrategy struct {}

const boardSize = 10;

func (s *RandomStrategy) TryAttack(board *entity.Board) bool {
	for {
		x := rand.Intn(boardSize);
		y := rand.Intn(boardSize);

		if board.AttackPosition(x, y) {
			return true;
		}
	}
}