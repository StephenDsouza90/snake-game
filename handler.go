package main

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

// newHandler is mapped to the /new path
func newHandler(context *gin.Context) {
	boardWidth, boardHeight, err := getBoardDimensions(context)
	if err != nil {
		errorHandler(context, http.StatusBadRequest, INVALID_NUMBER) // 400
		return
	}

	if !isBoardDimensionValid(boardWidth, boardHeight) {
		errorHandler(context, http.StatusBadRequest, INVALID_BOARD_DIMENSIONS) // 400
		return
	}

	state := getNewGame(boardWidth, boardHeight)

	if !isFruitWithinBoard(state) {
		errorHandler(context, http.StatusInternalServerError, FRUIT_OUT_OF_BOUNDS) // 500
		return
	}

	if !isSnakeWithinBoard(state) {
		errorHandler(context, http.StatusInternalServerError, SNAKE_OUT_OF_BOUNDS) // 500
		return
	}

	context.JSON(http.StatusOK, state)
}

// getBoardDimensions gets the board dimensions from the context.
func getBoardDimensions(context *gin.Context) (int, int, error) {
	boardWidth, err := strconv.Atoi(context.Query(WIDTH_PARAM))
	if err != nil && boardWidth <= 2 {
		return 0, 0, err
	}

	boardHeight, err := strconv.Atoi(context.Query(HEIGHT_PARAM))
	if err != nil && boardHeight <= 2 {
		return 0, 0, err
	}

	return boardWidth, boardHeight, nil
}

// isFruitWithinBoard checks if the fruit is within the board.
func isFruitWithinBoard(state state) bool {
	return state.Fruit.X < state.Width && state.Fruit.Y < state.Height
}

// isSnakeWithinBoard checks if the snake is within the board.
func isSnakeWithinBoard(state state) bool {
	return state.Snake.X < state.Width && state.Snake.Y < state.Height
}

// validateHandler is mapped to the /validate path
func validateHandler(context *gin.Context) {
	var rb requestBody
	if err := context.ShouldBindJSON(&rb); err != nil {
		errorHandler(context, http.StatusBadRequest, INVALID_BODY_REQUEST) // 400
		return
	}

	if !isUnixNanoTimestamp(rb.GameID) {
		errorHandler(context, http.StatusBadRequest, INVALID_GAME_ID) // 400
		return
	}

	if !isSnakeMoving(rb) {
		errorHandler(context, http.StatusBadRequest, SNAKE_NOT_MOVING) // 400
		return
	}

	fruitFound, rb, err := play(rb)
	if err != nil {
		handlePlayError(context, err) // 400, 418
		return
	}

	if !isFruitWithinBoard(rb.state) {
		errorHandler(context, http.StatusInternalServerError, FRUIT_OUT_OF_BOUNDS) // 500
		return
	}

	if !isSnakeWithinBoard(rb.state) {
		errorHandler(context, http.StatusInternalServerError, SNAKE_OUT_OF_BOUNDS) // 500
		return
	}

	if fruitFound {
		context.JSON(http.StatusOK, rb.state) // 200
		return
	}

	errorHandler(context, http.StatusNotFound, FRUIT_NOT_FOUND) // 404
}

// isSnakeMoving checks if the snake is moving.
func isSnakeMoving(rb requestBody) bool {
	return len(rb.Ticks) > 0
}

// handlePlayError handles the error from the play function.
func handlePlayError(context *gin.Context, err error) {
	switch err.Error() {
	case INVALID_MOVE:
		errorHandler(context, http.StatusBadRequest, INVALID_MOVE) // 400
	case REVERSE_MOVE:
		errorHandler(context, http.StatusTeapot, REVERSE_MOVE) // 418
	case DIAGONAL_MOVE:
		errorHandler(context, http.StatusTeapot, DIAGONAL_MOVE) // 418
	case OUT_OF_BOUNDS:
		errorHandler(context, http.StatusTeapot, OUT_OF_BOUNDS) // 418
	}
}

// Handle errors
func errorHandler(context *gin.Context, code int, message string) {
	context.JSON(code, gin.H{ERROR: message})
}

// isUnixNanoTimestamp checks if the string is a Unix nano timestamp.
// This is done to prevent SQL Injection and XSS attacks.
func isUnixNanoTimestamp(gameId string) bool {
	regex := regexp.MustCompile(`^\d+$`)
	return regex.MatchString(gameId)
}
