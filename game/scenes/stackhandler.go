package scenes

type StackHandler struct {
	stack *SceneStack
	ctx *GameContext
}

// SetStack serve para "passar a stack adiante" sem precisar explicitamente passar de scene em scene
func (b *StackHandler) SetStack(s *SceneStack) {
	b.stack = s
}


func (b *StackHandler) SetContext(ctx *GameContext) {
	b.ctx = ctx
}