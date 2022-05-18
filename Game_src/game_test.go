package wordlegame

import (

	"testing"
)

func TestGameStartSate(t *testing.T) {
	game := *NewWordleGame()
	game.Start("start", 6)
	if game.answer != "start" {
		t.Fatalf("Game answer is not set")
	}
	if game.triesLeft != 6 {
		t.Fatalf("Game triesleft is not set")
	}

	if game.tries != 6 {
		t.Fatalf("Game tries is not set")
	}

}
func TestGameUpdateState(t *testing.T) {
	game := *NewWordleGame()
	game.Start("start", 6)
	game.Update("patti")
	game.Update("maker")
	game.Update("happy")
	if game.userInput != "happy" {
		t.Fatalf("userInput not updated")
	}
	if game.triesLeft != 3 {
		t.Fatalf("triesLeft not updated")
	}
	game.Update("stpak")

	if string(game.gameOutput) != "**_._" {
		t.Fatalf("gameoutput not updated")
	}
	game.Update("st")

	if string(game.gameOutput) != "**" {
		t.Fatalf("gameoutput not updated")
	}
	

}
func TestGameOver(t *testing.T) {
	game := *NewWordleGame()
	game.Start("start", 6)
	game.Update("patti")
	game.Update("patti")
	if game.IsGameOver() == true {
		t.Fatalf("gameover not working")
	}
	game.Update("patti")
	game.Update("patti")
	game.Update("patti")
	game.Update("patti")
	if game.triesLeft != 0 {
		t.Fatalf("gametries not working")
	}
	if game.IsGameOver() == false {
		t.Fatalf("gameover not working")
	}

}
func TestIsWinner(t *testing.T) {
	game := *NewWordleGame()
	game.Start("start", 6)
	game.Update("patti")
	if game.IsWinner(){
		t.Fatalf("Is winner not working")

	} 
	game.Update("start")
	if !game.IsWinner(){
		t.Fatalf("Is winner not working")

	} 

}
func TestGamesPlayed(t *testing.T){
	game := *NewWordleGame()
	if game.played!=0{
		t.Fatalf("games played not working")
	}
	game.Start("start", 6)
	if game.played!=1{
		t.Fatalf("games played not working")
	}

}
func TestWinscount(t *testing.T) {
	game := *NewWordleGame()
	if game.wins!=0{
		t.Fatalf("games wins not working")
	}
	game.Start("start", 6)
	game.Update("patti")
	if game.wins!=0{
		t.Fatalf("games wins not working")
	}
	game.Update("start")
	if game.wins!=1{
		t.Fatalf("games wins not working")
	}
	}
func TestWinStreak(t *testing.T){
	game := *NewWordleGame()
	game.Start("start", 6)
    game.Update("patti")
	if game.winstreak!=0{
		t.Fatalf("games winsstreak not working")
	}
	game.Update("start")
	if game.winstreak!=1{
		t.Fatalf("games winstreak not working")
	}
	game.Start("happy", 6)
	game.Update("happy")
	if game.winstreak!=2{
		t.Fatalf("games winstreak not working")
	}
	game.Start("makes", 3)
	game.Update("happy")
	game.Update("start")
	game.Update("other")
	if game.winstreak!=0{
		t.Fatalf("games winstreak not working")
	}

}
func TestMaxWinStreak(t *testing.T){
	game := *NewWordleGame()
	game.Start("start", 6)
    game.Update("patti")
	if game.maxwinstreak!=0{
		t.Fatalf("games maxwinsstreak not working")
	}
	game.Update("start")
	if game.maxwinstreak!=1{
		t.Fatalf("games maxwinstreak not working")
	}
	game.Start("happy", 6)
	game.Update("happy")
	if game.maxwinstreak!=2{
		t.Fatalf("games maxwinstreak not working")
	}
	game.Start("makes", 3)
	game.Update("happy")
	game.Update("start")
	game.Update("other")
	if game.maxwinstreak!=2{
		t.Fatalf("games maxwinstreak not working")
	}
	game.Start("makes", 3)
	game.Update("makes")
	game.Start("happy", 3)
	game.Update("happy")
	game.Start("paper", 3)
	game.Update("paper")
	if game.maxwinstreak!=3{
		t.Fatalf("games maxwinstreak not working")
	}
}
