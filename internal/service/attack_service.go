package service

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/allanjose001/go-battleship/game/shared/board"
	"github.com/allanjose001/go-battleship/internal/ai"
	"github.com/allanjose001/go-battleship/internal/entity"
)

type AttackService struct{}

func NewAttackService() *AttackService {
	// Inicializa a semente do randomizador para os sons mudarem a cada partida
	rand.Seed(time.Now().UnixNano())
	return &AttackService{}
}

// PlayerAttack com randomização de áudio
func (s *AttackService) PlayerAttack(aiBoard *board.Board, row, col int, attempts, hits, totalShipCells int) (int, int, bool, bool) {
	cell := &aiBoard.Cells[row][col]

	if cell.State == board.Hit || cell.State == board.Miss {
		return attempts, hits, false, false
	}

	attempts++

	if cell.State == board.Ship {
		cell.State = board.Hit
		hits++

		// RANDOMIZER DE ATAQUE: Sorteia um número entre 1 e 4
		num := rand.Intn(4) + 1
		soundPath := fmt.Sprintf("assets/audio/sfx/attack%d.wav", num)
		fmt.Printf("🎰 [Randomizer] Sorteado: %s\n", soundPath)

		// Nota: Aqui você precisa ter acesso ao seu SoundService global ou contexto
		// Se o audioService não estiver acessível, este código dará erro de compilação.
		// O ideal é disparar o som na Scene, mas se quiser aqui:
		// audioService.PlaySFX(soundPath)

		if hits >= totalShipCells {
			return attempts, hits, true, true
		}
		return attempts, hits, true, false
	}

	if cell.State == board.Empty {
		cell.State = board.Miss

		// RANDOMIZER DE ÁGUA: Sorteia um número entre 1 e 5
		num := rand.Intn(5) + 1
		soundPath := fmt.Sprintf("assets/audio/sfx/waterSplash%d.wav", num)
		fmt.Printf("🎰 [Randomizer] Sorteado: %s\n", soundPath)

		// audioService.PlaySFX(soundPath)
	}

	return attempts, hits, false, false
}

// AITurn permanece com a lógica de sincronização
func (s *AttackService) AITurn(aiPlayer *ai.AIPlayer, entityBoard *entity.Board, playerBoard *board.Board, attempts, hits, totalShipCells int) (int, int, bool) {
	if aiPlayer == nil {
		return attempts, hits, false
	}

	attempts++
	aiPlayer.Attack(entityBoard)

	for r := 0; r < board.Rows; r++ {
		for c := 0; c < board.Cols; c++ {
			entPos := entityBoard.Positions[r][c]
			cell := &playerBoard.Cells[r][c]

			if entity.IsAttacked(entPos) && cell.State != board.Hit && cell.State != board.Miss {
				if cell.State == board.Ship {
					cell.State = board.Hit
					hits++
					if hits >= totalShipCells {
						return attempts, hits, true
					}
				} else {
					cell.State = board.Miss
				}
			}
		}
	}

	return attempts, hits, false
}
