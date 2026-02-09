package main

import (
	//"gobattleship/UI"
	//"fmt"
	//"log"

	"github.com/allanjose001/go-battleship/internal/ai"
	"github.com/allanjose001/go-battleship/internal/entity"
	"github.com/allanjose001/go-battleship/internal/service"
 	"github.com/allanjose001/go-battleship/game"
	"github.com/allanjose001/go-battleship/game/components"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	//========= teste inserção no tabuleiro ============
	board1 := new(entity.Board);

	fleet1 := entity.NewFleet();

	board1.PlaceShip(fleet1.GetShipByIndex(0), 1, 1)
	board1.PlaceShip(fleet1.GetShipByIndex(1), 4, 4)
	fleet1.GetShipByIndex(2).Rotate()
	board1.PlaceShip(fleet1.GetShipByIndex(2), 7, 7)
	fleet1.GetShipByIndex(3).Rotate()
	board1.PlaceShip(fleet1.GetShipByIndex(3), 5, 1)
	board1.PlaceShip(fleet1.GetShipByIndex(4), 6, 4)

	entity.PrintBoard(board1);
	
	//========= teste AI ===========
	
	//aiPlayer := ai.NewEasyAIPlayer();
	//aiPlayer := ai.NewMediumAIPlayer(fleet1);
	aiPlayer := ai.NewHardAIPlayer(fleet1);

	for i := 0; i < 50; i++ {
	aiPlayer.Attack(board1);
	entity.PrintBoard(board1);
	}

	//========== teste de profile ===========

	profile1 := new(service.Profile);
	profile1.Username = "Player1";
	profile1.TotalScore = 200
	profile1.HighestScore = 50
	profile1.GamesPlayed = 5
	profile1.MedalsEarned = 2

	service.SaveProfile(*profile1);
	//err1 := service.SaveProfile(*profile1)
  
  //profile2, err := service.FindProfile("Player2");
	//if err != nil {
	//	log.Fatal(err)
	//}

	//fmt.Printf("perfil encontrado: %+v\n", profile2);

	//service.RemoveProfile("Player1");
  
  //============= teste do front ========================

	components.InitFonts() //carrega a fonte apenas uma vez
	g := game.NewGame()
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}

}
