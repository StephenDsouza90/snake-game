# Snake Challenge Solution

## Introduction

The Snake Challenge Solution is a web application that allows users to play a snake game. It provides two main functionalities: creating a new game and validating moves made by the snake within the game.

The application is built using the [Gin](https://gin-gonic.com/) web framework in Go programming language. Gin offers fast performance and supports JSON validation, making it suitable for handling HTTP requests and responses.

## Getting Started

To run the application locally, you have two options:

1. Use the standard Go command:

```bash
$ go run .
```

2. Use the Make command:

```bash
$ make run
```

Once the application is running, you'll see the message:

```bash
$ Server is running on localhost:8080
```

## Handlers

Handlers are responsible for managing incoming requests and generating appropriate responses. There are two handlers:

`newHandler`: Handles requests to create a new game with specified dimensions (width and height). It validates the input parameters, initializes a new game state, and returns the game state as a JSON response.

`validateHandler`: Validates the moves of the snake within an existing game. It takes the current game state and a series of moves (ticks) as input, validates the moves, and updates the game accordingly. If the snake eats the fruit, it increments the score and generates a new fruit position. It then returns the updated game state.

## Game Logic

The game logic consists of functions responsible for initializing a new game and processing moves within the game.

`getNewGame`: Initializes a new game with specified dimensions. It sets the initial position, velocity of the snake and places a fruit at a random position on the board.

`play`: Processes a series of moves within the game. It updates the snake's position based on its velocity and validates the moves. If the snake eats the fruit, it increments the score and generates a new fruit position.

## Testing

Unit tests are provided for both handlers and game logic. To run the tests locally, you can use the following commands:

1. Standard Go command:

```bash
$ go test ./... -v
```

2. Make command:

```bash
$ make test
```

## API Usage Examples

### Starting a new game

**GET Request**

```bash
$ curl --GET 'http://localhost:8080/new?w=6&h=6'
```

**Response**

```JSON
{
    "gameId": "1714933014754623000",
    "width": 6,
    "height": 6,
    "score": 0,
    "fruit": {
        "x": 3,
        "y": 5
    },
    "snake": {
        "x": 0,
        "y": 0,
        "velX": 1,
        "velY": 0
    }
}
```

**Display**

```
Game board visualization with snake and fruit
=============================================
Game ID: 1714933014754623000 Score: 0

SOOOOO
OOOOOO
OOOOOO
OOOOOO
OOOOOO
OOOFOO
```

The game is displayed visually in the terminal:

- **S** represents the **Snake**
- **F** represents the **Fruit**
- **O** represents an **Empty Cell** for the Snake to move

### Playing the game

**POST Request**

```bash
$ curl --POST 'http://localhost:8080/validate' \
--header 'Content-Type: application/json' \
--data '{
     "gameId": "1714933014754623000",
    "width": 6,
    "height": 6,
    "score": 0,
    "fruit": {
        "x": 3,
        "y": 5
    },
    "snake": {
        "x": 0,
        "y": 0,
        "velX": 1,
        "velY": 0
    },
    "ticks": [
        { "velX": 1, "velY": 0 },
        { "velX": 1, "velY": 0 },
        { "velX": 1, "velY": 0 },
        { "velX": 0, "velY": 1 },
        { "velX": 0, "velY": 1 },
        { "velX": 0, "velY": 1 },
        { "velX": 0, "velY": 1 },
        { "velX": 0, "velY": 1 }
    ]
}'
```

**Response**

```JSON
{
    "gameId": "1714933014754623000",
    "width": 6,
    "height": 6,
    "score": 1,
    "fruit": {
        "x": 1,
        "y": 2
    },
    "snake": {
        "x": 3,
        "y": 5,
        "velX": 0,
        "velY": 1
    }
}
```

**Display**

```
Move: 1

OSOOOO
OOOOOO
OOOOOO
OOOOOO
OOOOOO
OOOFOO

Move: 2

OOSOOO
OOOOOO
OOOOOO
OOOOOO
OOOOOO
OOOFOO

Move: 3

OOOSOO
OOOOOO
OOOOOO
OOOOOO
OOOOOO
OOOFOO

Move: 4

OOOOOO
OOOSOO
OOOOOO
OOOOOO
OOOOOO
OOOFOO

Move: 5

OOOOOO
OOOOOO
OOOSOO
OOOOOO
OOOOOO
OOOFOO

Move: 6

OOOOOO
OOOOOO
OOOOOO
OOOSOO
OOOOOO
OOOFOO

Move: 7

OOOOOO
OOOOOO
OOOOOO
OOOOOO
OOOSOO
OOOFOO

Move: 8

OOOOOO
OOOOOO
OOOOOO
OOOOOO
OOOOOO
OOOSOO

New Game

OOOOOO
OOOOOO
OFOOOO
OOOOOO
OOOOOO
OOOSOO
```