package state

import "github.com/allanjose001/go-battleship/internal/entity"

// GameContext possui dados de interesse do jogo (tela de jogo, perfis, etc)
type GameContext struct {
	Profile *entity.Profile
	Match   *entity.Match
}
