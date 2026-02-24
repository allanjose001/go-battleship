package state

import "github.com/allanjose001/go-battleship/internal/entity"


type GameContext struct {
	Profile *entity.Profile
    //Match *Match
}

type contextAware interface {
	SetContext(*GameContext)
}