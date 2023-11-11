# Snake Game

A project that includes a stateless API for a Snake Game, developed in the Go programming language.

## Description

- The snake will never be longer than length 1. It will not grow after eating a fruit.
- 1 single fruit at a time on the game board.
- If the snake hits the edge of the game bounds, the game is over.
- The snake will always start at position (0, 0), with a velocity of (1, 0).
- The gameId value will not be subject to validation and will be disregarded.

## Getting Started

### Dependencies

- Gorilla Mux


### Installing

```bash
git clone https://github.com/paimanbandi/snakegame
```

### Running

```bash
cd snakegame/cmd/snakegame
go build
./snakegame
```

### Testing

I provide the postman collection, which you can utilize to test the APIs.

