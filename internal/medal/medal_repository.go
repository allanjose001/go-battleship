package medal

import (
	"time"

	"github.com/allanjose001/go-battleship/internal/entity"
)

// MedalsList lista de todas as medalhas do jogo com os caminhos atualizados
var MedalsList = []*Medal{
	{
		Name:         "Almirante",
		Description:  "Venceu sem perder navios",
		IconPath:     "assets/medals/Medalha1.png",
		GrayIconPath: "assets/medals/Interrogação.png",
		Verification: func(stats entity.PlayerStats) bool {
			return stats.WinWithoutLosses
		},
	},
	{
		Name:         "Capitão",
		Description:  "Acertou 7 tiros seguidos",
		IconPath:     "assets/medals/Medalha2.png",
		GrayIconPath: "assets/medals/Interrogação.png",
		Verification: func(stats entity.PlayerStats) bool {
			return stats.HigherHitSequence >= 7
		},
	},
	{
		Name:         "Capitão de Mar e Guerra",
		Description:  "Acertou 8 tiros seguidos",
		IconPath:     "assets/medals/Medalha3.png",
		GrayIconPath: "assets/medals/Interrogação.png",
		Verification: func(stats entity.PlayerStats) bool {
			return stats.HigherHitSequence >= 8
		},
	},
	{
		Name:         "Marinheiro",
		Description:  "Venceu em 1 minuto",
		IconPath:     "assets/medals/Medalha4.png",
		GrayIconPath: "assets/medals/Interrogação.png",
		Verification: func(stats entity.PlayerStats) bool {
			// Evita erro se o tempo for 0 (nunca jogou)
			return stats.FasterTime > 0 && stats.FasterTime <= time.Minute.Milliseconds()
		},
	},
}

// MedalsMap Map para acesso rápido pelo nome
var MedalsMap = make(map[string]*Medal)

// init inicializa map para facilitar load profile do json com medalhas
func init() {
	for _, m := range MedalsList {
		MedalsMap[m.Name] = m
	}
}

// GetMedals serve para pegar os objetos medal pelo nome
func GetMedals(names []string) []*Medal {
	var result []*Medal
	for _, n := range names {
		if m, ok := MedalsMap[n]; ok {
			result = append(result, m)
		}
	}
	return result
}
