package main

import (
	"testing"
)

func TestInitializeSnake(t *testing.T) {
	snake := initializeSnake(1)
	if snake.VelX != 1 {
		t.Errorf("Expected velocity 1, got %d", snake.VelX)
	}
}

func TestGenerateFruitPosition(t *testing.T) {
	snake := snake{X: 5, Y: 5}
	fruit := generateFruitPosition(10, 10, snake)
	if fruit.X == snake.X && fruit.Y == snake.Y {
		t.Errorf("Fruit is on the snake")
	}
}

func TestCreateState(t *testing.T) {
	snake := snake{X: 5, Y: 5}
	fruit := fruit{X: 2, Y: 3}
	state := createState("11", 10, 10, 0, fruit, snake)
	if state.GameID != "11" || state.Width != 10 || state.Height != 10 || state.Score != 0 || state.Fruit != fruit || state.Snake != snake {
		t.Errorf("State creation failed")
	}
}

func TestIsMoveValid_WithCorrectMove(t *testing.T) {
	if !isMoveValid(1) || !isMoveValid(0) || !isMoveValid(-1) {
		t.Errorf("Move validation failed")
	}
}

func TestIsMoveValid_WithIncorrectMove(t *testing.T) {
	if isMoveValid(2) {
		t.Errorf("Move validation failed")
	}
}

func TestIsReverseMove_WithCorrectMove(t *testing.T) {
	if isReverseMove(1, 1, 0, 0) {
		t.Errorf("Reverse move validation failed")
	}
}

func TestIsReverseMove_WithIncorrectMove(t *testing.T) {
	if !isReverseMove(1, -1, 0, 0) {
		t.Errorf("Reverse move validation failed")
	}
}

func TestIsDiagonalMove_WithCorrectMove(t *testing.T) {
	if isDiagonalMove(1, 0) {
		t.Errorf("Diagonal move validation failed")
	}
}

func TestIsDiagonalMove_WithIncorrectMove(t *testing.T) {
	if !isDiagonalMove(1, 1) {
		t.Errorf("Diagonal move validation failed")
	}
}

func TestIsWithinBoundaries_WithCorrectValues(t *testing.T) {
	if isWithinBoundaries(5, 5, 10, 10) {
		t.Errorf("Boundary check failed")
	}
}

func TestIsWithinBoundaries_WithIncorrectValues(t *testing.T) {
	if !isWithinBoundaries(-5, -5, 10, 10) {
		t.Errorf("Boundary check failed")
	}
}

func TestUpdateSnake(t *testing.T) {
	rb := requestBody{state: state{Snake: snake{X: 5, Y: 5, VelX: 1, VelY: 0}}}
	rb = updateSnake(rb, 6, 5, 1, 0)
	if rb.Snake.X != 6 || rb.Snake.Y != 5 || rb.Snake.VelX != 1 || rb.Snake.VelY != 0 {
		t.Errorf("Snake update failed")
	}
}

func TestIncrementScore(t *testing.T) {
	rb := requestBody{state: state{Score: 0}}
	rb = incrementScore(rb)
	if rb.Score != 1 {
		t.Errorf("Score increment failed")
	}
}
