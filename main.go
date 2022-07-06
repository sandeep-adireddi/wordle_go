/*
package main is main package of this program
*/
package main

import (
	"fmt"
	"log"
	"math/rand"

	"time"

	wordlegame "example.com/wordleclone/Game_src"
	worddata "example.com/wordleclone/word_data"
)

const Instructions = `1)You have to guess the Wordle in six goes or less
2)A correct letter turns green, correct letter in the wrong place turns yellow, incorrect letter turns red
3)Letters can be used more than once
4)Every word you enter must be in the word list`

func isValidWord(userInput string, wordData []string) bool {
	for _, word := range wordData {
		if userInput == word {
			return true
		}

	}
	return false
}
func isValidInput(userInput string, wordData []string) (bool, string) {
	if len(userInput) != 5 {
		return false, "⚠️  Enter only 5 letter word ⚠️"
	}
	if !isValidWord(userInput, wordData) {
		return false, "⚠️  Not in word list ⚠️"
	}
	return true, ""
}
func colorPrintOutput(output, userInput string) {
	const Green, Yellow, Red, Reset = "\033[32m", "\033[33m", "\033[31m", "\033[0m"

	for pos, char := range output {

		switch {
		case char == '*':
			fmt.Printf(Green + string(userInput[pos]) + Reset)
		case char == '.':
			fmt.Printf(Yellow + string(userInput[pos]) + Reset)
		default:
			fmt.Printf(Red + string(userInput[pos]) + Reset)
		}
	}
}
func main() {
	words, err := worddata.Read_txt("sgb-words.txt")
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	seed := rand.NewSource(time.Now().UnixNano())
	random_genrator := rand.New(seed)
	game := wordlegame.NewWordleGame()
	fmt.Println("INSTRUCTIONS")
	fmt.Println(Instructions)
	want_to_play := "y"
	for want_to_play == "y" {
		random_word_index := random_genrator.Intn(len(words))

		game.Start(words[random_word_index], 6)
		var userInput string
		for !game.IsGameOver() {

			fmt.Println("Number of tries Left: ", game.GetTriesLeft())
			fmt.Scanf("%s", &userInput)
			if ok, err := isValidInput(userInput, words); ok {
				game.Update(userInput)
				colorPrintOutput(game.GetGameOutput(), userInput)
				fmt.Println()
				if game.IsWinner() {
					fmt.Println("🎉you won🎉")
					fmt.Printf("Played: %d Won: %d Current Streak: %d Max Streak: %d\n", game.GetPlayedCount(), game.GetWinCount(), game.GetWinStreakCount(), game.GetMaxWinStreakCount())
					break
				}
			} else {
				fmt.Println(err)
			}

		}
		if !game.IsWinner() && game.IsGameOver() {
			fmt.Println("Answer is: ", words[random_word_index])
			fmt.Printf("Played: %d Won: %d Current Streak: %d Max Streak: %d\n", game.GetPlayedCount(), game.GetWinCount(), game.GetWinStreakCount(), game.GetMaxWinStreakCount())
			fmt.Println("Better luck next time 🤝")
		}
		fmt.Printf("Wanna play again [y/n]: ")
		fmt.Scanf("%s", &want_to_play)
	}
}
