package main

// state represents the game state.
type state struct {
	GameID string `json:"gameId"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Score  int    `json:"score"`
	Fruit  fruit  `json:"fruit"`
	Snake  snake  `json:"snake"`
}

// fruit represents the fruit position.
type fruit struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// snake represents the snake position.
type snake struct {
	X    int `json:"x"`
	Y    int `json:"y"`
	VelX int `json:"velX"`
	VelY int `json:"velY"`
}

// tick represents the snake velocity.
type tick struct {
	VelX int `json:"velX"`
	VelY int `json:"velY"`
}

// requestBody represents the request body.
type requestBody struct {
	state
	Ticks []tick `json:"ticks"`
}
