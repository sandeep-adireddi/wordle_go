package wordlegame

type wordlegame struct {
	answer       string
	userInput    string
	gameOutput   string
	tries        uint8
	triesLeft    uint8
	played       uint
	wins         uint
	winstreak    uint
	maxwinstreak uint
}

func createGameOutput(answer, userInput string) string {
	var output []rune
	for j := 0; j < len(userInput); j++ {
		output = append(output, '_')
	}

	answer_map := make(map[rune]int)

	for _, char := range answer {

		answer_map[char] += 1
	}

	for pos, char := range userInput {

		if userInput[pos] == answer[pos] {
			output[pos] = rune('*')
			answer_map[char] -= 1

		}

	}
	for pos, char := range userInput {
		if val, ok := answer_map[char]; ok && val != 0 && output[pos] != '*' {

			output[pos] = rune('.')
			answer_map[char] -= 1
		}
	}

	outputString := string(output)

	return outputString
}
func NewWordleGame() *wordlegame {

	return &wordlegame{answer: "", userInput: "", gameOutput: "", tries: 0, triesLeft: 0,
		played: 0, wins: 0, winstreak: 0, maxwinstreak: 0}
}
func (game *wordlegame) Start(answer string, tries uint8) {
	game.answer = answer
	game.tries = tries
	game.triesLeft = tries
	game.played += 1

}
func (game *wordlegame) Update(userInput string) {
	game.userInput = userInput
	game.triesLeft -= 1
	game.gameOutput = createGameOutput(game.answer, game.userInput)
	if game.IsWinner() {
		game.wins += 1
		game.winstreak += 1
		if game.maxwinstreak < game.winstreak {
			game.maxwinstreak = game.winstreak
		}
	}
	if game.IsGameOver() {
		game.winstreak = 0
	}

}
func (game *wordlegame) IsWinner() bool {
	var output string
	for i := 0; i < len(game.answer); i++ {
		output += "*"
	}
	return game.gameOutput == output
}

func (game *wordlegame) IsGameOver() bool {

	return game.triesLeft <= 0
}
func (game *wordlegame) GetGameOutput() string {

	return game.gameOutput
}
func (game *wordlegame) GetTriesLeft() uint8 {

	return game.triesLeft
}
func (game *wordlegame) GetPlayedCount() uint {
	return game.played
}
func (game *wordlegame) GetWinCount() uint {
	return game.wins
}
func (game *wordlegame) GetWinStreakCount() uint {
	return game.winstreak
}
func (game *wordlegame) GetMaxWinStreakCount() uint {
	return game.maxwinstreak
}
