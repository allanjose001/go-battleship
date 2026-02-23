package service

import (
	"fmt"
	"github.com/allanjose001/go-battleship/internal/ai"
	"github.com/allanjose001/go-battleship/internal/entity"
)

type CampaignService struct {
	matchService *MatchService 
}

func NewCampaignService(ms *MatchService) *CampaignService {
	return &CampaignService{
		matchService: ms,
	}
}

// Deve ser chamado no "continuar campanha" do front
func (cs *CampaignService) StartCampaignMatch(username string, fleet *entity.Fleet) error {
	profile, err := FindProfile(username)
	if err != nil {
		return err
	}

	// Inicializa campanha se não existir
	if profile.CurrentCampaign == nil {
		profile.CurrentCampaign = &entity.Campaign{
			ID:             fmt.Sprintf("camp_%s", username),
			DifficultyStep: make(map[string]entity.MatchResult),
			IsActive:       true,
		}
	}

	diff, finished := cs.GetNextDifficulty(profile.CurrentCampaign)
	if finished {
		return fmt.Errorf("campanha para %s já foi concluída", username)
	}

	opponent := cs.selectAI(diff, fleet)
	_ = opponent // remover quando descomentar a linha 44

	// INTEGRAÇÃO 
	// return cs.matchService.StartMatch(profile, opponent, diff) 

	fmt.Printf("Campanha: %s enfrentando IA %s\n", username, diff)
	return nil
}

// HandleCampaignResult será chamado automaticamente pelo MatchService ao fim da partida
func (cs *CampaignService) HandleCampaignResult(username string, diff string, result entity.MatchResult) error {
	profile, err := FindProfile(username)
	if err != nil {
		return err
	}

	if profile.CurrentCampaign == nil {
		return fmt.Errorf("nenhuma campanha ativa para %s", username)
	}

	// 1. Registra o resultado
	profile.CurrentCampaign.DifficultyStep[diff] = result

	// 2. Verifica se a campanha inteira terminou
	_, finished := cs.GetNextDifficulty(profile.CurrentCampaign)
	if finished && result.Win {
		profile.CurrentCampaign.IsActive = false
		profile.Campaigns = append(profile.Campaigns, *profile.CurrentCampaign)
		profile.CurrentCampaign = nil
	}

	// 3. Persiste no JSON e atualiza estatísticas gerais
	// Usamos o AddMatchToProfile para que a partida de campanha também conte no histórico global
	_, err = AddMatchToProfile(profile, result)
	return err
}

// GetNextDifficulty (Lógica interna)
func (cs *CampaignService) GetNextDifficulty(c *entity.Campaign) (string, bool) {
	if c.DifficultyStep == nil {
		return "easy", false
	}
	steps := []string{"easy", "medium", "hard"}
	for _, step := range steps {
		if res, ok := c.DifficultyStep[step]; !ok || !res.Win {
			return step, false
		}
	}
	return "", true
}

// selectAI (Lógica interna)
func (cs *CampaignService) selectAI(diff string, fleet *entity.Fleet) *ai.AIPlayer {
	switch diff {
	case "medium":
		return ai.NewMediumAIPlayer(fleet)
	case "hard":
		return ai.NewHardAIPlayer(fleet)
	default:
		return ai.NewEasyAIPlayer()
	}
}