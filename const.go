package main

const (
	// Server address
	SERVER = "localhost:8080"

	// API path
	NEW_GAME = "/new"
	VALIDATE = "/validate"

	// API parameters
	WIDTH_PARAM  = "w"
	HEIGHT_PARAM = "h"

	// Board Symbols
	SNAKE_SYMBOL = "S"
	FRUIT_SYMBOL = "F"
	EMPTY_SYMBOL = "O"

	// Initial values
	INITIAL_GAME_SCORE = 0
	INITIAL_VELOCITY_X = 1
	WIDTH_START        = 2
	HEIGHT_START       = 2

	// Messages
	INVALID_MOVE             = "invalid move"
	REVERSE_MOVE             = "snake made a 180-degree turn"
	DIAGONAL_MOVE            = "snake moved diagonal"
	OUT_OF_BOUNDS            = "snake is out of bounds"
	FRUIT_NOT_FOUND          = "fruit not found"
	SNAKE_NOT_MOVING         = "snake is not moving"
	INVALID_BOARD_DIMENSIONS = "invalid board dimensions"
	INVALID_BODY_REQUEST     = "invalid body request"
	SNAKE_OUT_OF_BOUNDS      = "snake is out of bounds"
	FRUIT_OUT_OF_BOUNDS      = "fruit is out of bounds"
	INVALID_NUMBER           = "not a number"
	INVALID_GAME_ID          = "invalid game id"
	METHOD_NOT_ALLOWED       = "method not allowed"

	// Keys
	ERROR = "error"
)
