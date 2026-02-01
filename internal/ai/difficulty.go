package ai

func NewEasyAIPlayer() *AIPlayer {
	return &AIPlayer{
		Strategies: []Strategy{
			&RandomStrategy{},
		},
	}
}

