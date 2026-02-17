package entity

type Profile struct {
	Username    string        `json:"username"`
	Stats       PlayerStats   `json:"player_stats"` //evitei field promotion para facilitar jason
	MedalsNames []string      `json:"medals"`       //armazena apenas nomes
	History     []MatchResult `json:"history"`
}

// AddMatch adiciona partida ao histórico e atualiza stats
func (p *Profile) AddMatch(r MatchResult) {
	p.History = append(p.History, r)
	p.Stats.ApplyMatch(r)
}
