package main

import (
	"fmt"
	"math/rand"
	"time"
)

// getNewGame initializes a new game state.
func getNewGame(boardWidth int, boardHeight int) state {
	gameId := fmt.Sprint(time.Now().UnixNano())
	snake := initializeSnake(INITIAL_VELOCITY_X)
	fruit := generateFruitPosition(boardWidth, boardHeight, snake)
	state := createState(gameId, boardWidth, boardHeight, INITIAL_GAME_SCORE, fruit, snake)

	displayNewGame(state)

	return state
}

// initializeSnake initializes the snake with a given velocity.
func initializeSnake(velX int) snake {
	return snake{
		VelX: velX,
	}
}

// generateFruitPosition generates a random position for the fruit ensuring it is not on the snake.
func generateFruitPosition(boardWidth int, boardHeight int, snake snake) fruit {
	fruitX := rand.Intn(boardWidth)
	fruitY := rand.Intn(boardHeight)

	// if fruitX == snake.X && fruitY == snake.Y {
	//	return generateFruitPosition(boardWidth, boardHeight, snake)
	// }

	return fruit{X: fruitX, Y: fruitY}
}

// createState creates a new game state.
func createState(gamedId string, width int, height int, score int, fruit fruit, snake snake) state {
	return state{
		GameID: gamedId,
		Width:  width,
		Height: height,
		Score:  score,
		Fruit:  fruit,
		Snake:  snake,
	}
}

// play simulates the game based on the given ticks and returns the game state.
func isBoardDimensionValid(boardWidth int, boardHeight int) bool {
	return boardWidth > WIDTH_START && boardHeight > HEIGHT_START
}

// play simulates the game based on the given ticks and returns the game state.
func play(rb requestBody) (bool, requestBody, error) {
	var err error
	fruitFound := false

	for index, tick := range rb.Ticks {
		rb.Snake.X += tick.VelX
		rb.Snake.Y += tick.VelY

		fmt.Println("Move:", index+1)
		printGame(rb.state)

		if rb.Snake.X == rb.Fruit.X && rb.Snake.Y == rb.Fruit.Y {
			fruitFound = true
		}

		if err = validateMove(tick, rb); err != nil {
			break
		}

		rb.Snake.VelX = tick.VelX
		rb.Snake.VelY = tick.VelY
	}

	if fruitFound {
		rb = updateGame(rb)

		fmt.Println("New Game")
		printGame(rb.state)
	}

	return fruitFound, rb, err
}

// validateMove validates the snake's move.
func validateMove(tick tick, rb requestBody) error {
	if !isMoveValid(tick.VelX) || !isMoveValid(tick.VelY) {
		return fmt.Errorf(INVALID_MOVE)
	}

	if isReverseMove(rb.Snake.VelX, tick.VelX, rb.Snake.VelY, tick.VelY) {
		return fmt.Errorf(REVERSE_MOVE)
	}

	if isDiagonalMove(tick.VelX, tick.VelY) {
		return fmt.Errorf(DIAGONAL_MOVE)
	}

	if isWithinBoundaries(rb.Snake.X, rb.Snake.Y, rb.Width, rb.Height) {
		return fmt.Errorf(OUT_OF_BOUNDS)
	}

	return nil
}

// isValidMove checks if the snake moved one unit at a time.
func isMoveValid(target int) bool {
	availableMoves := []int{-1, 0, 1}
	for _, move := range availableMoves {
		if move == target {
			return true
		}
	}
	return false
}

// isReverseMove checks if the snake has made a 180-degree turn.
func isReverseMove(previousSnakeVelX int, tickX int, previousSnakeVelY int, tickY int) bool {
	return (-previousSnakeVelX == tickX && tickX != 0) || (-previousSnakeVelY == tickY && tickY != 0)
}

// isDiagonalMove checks if the snake has moved diagonally.
func isDiagonalMove(tickX int, tickY int) bool {
	return tickX == tickY
}

// isWithinBoundaries checks if the snake is within the board boundaries.
func isWithinBoundaries(snakeX int, snakeY int, width int, height int) bool {
	return (snakeX < 0 || snakeX >= width || snakeY < 0 || snakeY >= height)
}

// updateGame updates the game state.
func updateGame(rb requestBody) requestBody {
	rb = incrementScore(rb)
	rb = updateSnake(rb, rb.Snake.X, rb.Snake.Y, rb.Snake.VelX, rb.Snake.VelY)
	rb.Fruit = generateFruitPosition(rb.Width, rb.Height, rb.Snake)
	return rb
}

// incrementScore increases the game score.
func incrementScore(rb requestBody) requestBody {
	rb.Score++
	return rb
}

// updateSnake updates the snake's position and velocity.
func updateSnake(rb requestBody, snakeX int, snakeY int, snakeVelX int, snakeVelY int) requestBody {
	rb.Snake.X = snakeX
	rb.Snake.Y = snakeY
	rb.Snake.VelX = snakeVelX
	rb.Snake.VelY = snakeVelY
	return rb
}

// printGame display the game board visually in the terminal.
func printGame(state state) {
	fmt.Println()
	for yAxis := 0; yAxis < state.Height; yAxis++ {
		for xAxis := 0; xAxis < state.Width; xAxis++ {
			if yAxis == state.Snake.Y && xAxis == state.Snake.X {
				fmt.Print(SNAKE_SYMBOL)
			} else if yAxis == state.Fruit.Y && xAxis == state.Fruit.X {
				fmt.Print(FRUIT_SYMBOL)
			} else {
				fmt.Print(EMPTY_SYMBOL)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// displayNewGame displays the game board with the snake and fruit.
func displayNewGame(state state) {
	fmt.Println()
	fmt.Println("Game board visualization with snake and fruit")
	fmt.Println("=============================================")
	fmt.Println("Game ID:", state.GameID, "Score:", state.Score)
	printGame(state)
}
