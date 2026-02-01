package ai

import "github.com/allanjose001/go-battleship/internal/entity"

type Strategy interface {
	TryAttack(board *entity.Board) bool 
}