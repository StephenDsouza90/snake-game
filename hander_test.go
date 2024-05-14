package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func getTestEngine() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.Default()
}

func TestNewHandler_WithCorrectDimensions(t *testing.T) {
	r := getTestEngine()
	r.GET(NEW_GAME, newHandler)

	req, _ := http.NewRequest(http.MethodGet, "/new?w=10&h=10", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected HTTP 200 OK, got %d", resp.Code)
	}
}

func TestNewHandler_WithIncorrectDimensions(t *testing.T) {
	r := getTestEngine()
	r.GET(NEW_GAME, newHandler)

	req, _ := http.NewRequest(http.MethodGet, "/new?w=-10&h=-10", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusBadRequest {
		t.Errorf("Expected HTTP 200 OK, got %d", resp.Code)
	}
}

func TestGetBoardDimensions(t *testing.T) {
	gin.SetMode(gin.TestMode)

	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	context.Request = httptest.NewRequest("GET", "/?w=10&h=10", nil)

	width, height, err := getBoardDimensions(context)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if width != 10 || height != 10 {
		t.Errorf("Expected width 10 and height 10, got width %d and height %d", width, height)
	}
}

func TestIsFruitWithinBoard_WhenFruitIsWithinBoard(t *testing.T) {
	state := state{
		Width:  10,
		Height: 10,
		Fruit:  fruit{X: 5, Y: 5},
	}

	if !isFruitWithinBoard(state) {
		t.Errorf("Expected fruit to be within board")
	}
}

func TestIsFruitWithinBoard_WhenFruitIsNotWithinBoard(t *testing.T) {
	state := state{
		Width:  10,
		Height: 10,
		Fruit:  fruit{X: 11, Y: 11},
	}

	if isFruitWithinBoard(state) {
		t.Errorf("Expected fruit to be within board")
	}
}

func TestIsSnakeWithinBoard_WhenSnakeIsWithinBoard(t *testing.T) {
	state := state{
		Width:  10,
		Height: 10,
		Snake:  snake{X: 5, Y: 5},
	}

	if !isSnakeWithinBoard(state) {
		t.Errorf("Expected snake to be within board")
	}
}

func TestIsSnakeWithinBoard_WhenSnakeIsNotWithinBoard(t *testing.T) {
	state := state{
		Width:  10,
		Height: 10,
		Snake:  snake{X: 11, Y: 11},
	}

	if isSnakeWithinBoard(state) {
		t.Errorf("Expected snake to be within board")
	}
}

func TestValidateHandler_WithCorrectRequest(t *testing.T) {
	r := getTestEngine()
	r.POST(VALIDATE, validateHandler)

	state := state{
		GameID: "123",
		Width:  10,
		Height: 10,
		Score:  0,
		Fruit:  fruit{X: 3, Y: 3},
		Snake:  snake{X: 0, Y: 0, VelX: 1, VelY: 0},
	}

	rb := requestBody{
		state: state,
		Ticks: []tick{
			{VelX: 1, VelY: 0},
			{VelX: 1, VelY: 0},
			{VelX: 1, VelY: 0},
			{VelX: 0, VelY: 1},
			{VelX: 0, VelY: 1},
			{VelX: 0, VelY: 1},
		},
	}

	body, _ := json.Marshal(rb)
	buf := bytes.NewBuffer(body)
	req, _ := http.NewRequest(http.MethodPost, VALIDATE, buf)
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected HTTP 200 OK, got %d", resp.Code)
	}
}

func TestIsSnakeMoving_WhenSnakeIsMoving(t *testing.T) {
	rb := requestBody{
		Ticks: []tick{
			{VelX: 1, VelY: 0},
		},
	}

	if !isSnakeMoving(rb) {
		t.Errorf("Expected snake to be moving")
	}
}

func TestIsSnakeMoving_WhenSnakeIsNotMoving(t *testing.T) {
	rb := requestBody{
		Ticks: []tick{},
	}

	if isSnakeMoving(rb) {
		t.Errorf("Expected snake to be moving")
	}
}

func TestGameIdForSqlInjection_WithSqlInjection(t *testing.T) {
	if isUnixNanoTimestamp("SELECT * FROM games;") {
		t.Errorf("Expected no SQL injection")
	}
}

func TestGameIdForSqlInjection_WithoutSqlInjection(t *testing.T) {
	if !isUnixNanoTimestamp("111111") {
		t.Errorf("Expected no SQL injection")
	}
}

func TestGameIdForXSS_WithXSS(t *testing.T) {
	if isUnixNanoTimestamp("<script>alert('XSS')</script>") {
		t.Errorf("Expected no XSS")
	}
}

func TestGameIdForXSS_WithoutXSS(t *testing.T) {
	if !isUnixNanoTimestamp("111111") {
		t.Errorf("Expected no XSS")
	}
}
