package scenes

import (
	"github.com/allanjose001/go-battleship/game/components/basic"
	"github.com/allanjose001/go-battleship/game/state"
	"github.com/hajimehoshi/ebiten/v2"
)

// SceneStack struct que gerencia rotas (para scenes que necessitam de compartilhar estado e/ou
// partilham de um fluxo)
type SceneStack struct {
	stack      []Scene
	ctx        *state.GameContext
	screenSize basic.Size
}

// stackAware como Interface interna: cenas que aceitam injeção da stack (para identificar as que usam)
type stackAware interface {
	SetStack(*SceneStack)
}

// contextAware tambem serve para passar context adiante
type contextAware interface {
	SetContext(*state.GameContext)
}

func NewSceneStack(size basic.Size, first Scene, ctx *state.GameContext) *SceneStack {
	s := &SceneStack{
		stack:      []Scene{},
		screenSize: size,
		ctx:        ctx,
	}

	s.Push(first)
	return s
}

func (s *SceneStack) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *SceneStack) Current() Scene {
	if len(s.stack) == 0 {
		return nil
	}
	return s.stack[len(s.stack)-1]
}

// Push adiciona scene à pilha e chama OnExit na anterior + OnEnter na scene do parâmetro
func (s *SceneStack) Push(next Scene) {
	// injeta stack se a cena suportar
	if aware, ok := next.(stackAware); ok {
		aware.SetStack(s)
	}

	//mesmo com context
	if aware, ok := next.(contextAware); ok {
		aware.SetContext(s.ctx)
	}

	var prev Scene
	if len(s.stack) > 0 {
		prev = s.stack[len(s.stack)-1]
		prev.OnExit(next)
	}

	s.stack = append(s.stack, next)
	next.OnEnter(prev, s.screenSize)
}

// Pop remove última scene e chama a anterior da pilha caso exista
func (s *SceneStack) Pop() {
	if len(s.stack) == 0 {
		return
	}

	top := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	var next Scene
	if len(s.stack) > 0 {
		next = s.stack[len(s.stack)-1]
	}

	top.OnExit(next)

	if next != nil {
		next.OnEnter(top, s.screenSize)
	}
}

// Replace troca sem passar estado
func (s *SceneStack) Replace(next Scene) {
	s.Pop()
	s.Push(next)
}

func (s *SceneStack) Update() error {
	current := s.Current()
	if current == nil {
		return nil
	}
	return current.Update()
}

func (s *SceneStack) Draw(screen *ebiten.Image) {
	current := s.Current()
	if current == nil {
		return
	}
	current.Draw(screen)
}
