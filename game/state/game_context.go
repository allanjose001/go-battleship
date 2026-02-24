package state

import "github.com/allanjose001/go-battleship/internal/entity"


type GameContext struct {
	Profile *entity.Profile
    Match *entity.Match
}

type ContextAware interface {
	SetContext(*GameContext)
}

func NewGameContext() *GameContext {
    return &GameContext{}
}

func (c *GameContext) SetProfile(p *entity.Profile) {
    c.Profile = p
}

func (c *GameContext) SetMatch(m *entity.Match) {
    c.Match = m
}