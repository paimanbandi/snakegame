package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"snakegame/internal/validations"
	"snakegame/pkg/utils"
)

func NewGame(w http.ResponseWriter, r *http.Request) {
	validations.IsGetMethod(w, r)

	width, err := strconv.Atoi(r.URL.Query().Get("w"))
	if err != nil {
		http.Error(w, "Invalid width", http.StatusBadRequest)
		return
	}

	height, err := strconv.Atoi(r.URL.Query().Get("h"))
	if err != nil {
		http.Error(w, "Invalid height", http.StatusBadRequest)
		return
	}

	//initialize game state
	state := State{
		GameID: utils.GenerateSecureID(),
		Width:  width,
		Height: height,
		Score:  0,
		Fruit:  generateFruitPosition(width, height),
		Snake:  Snake{X: 0, Y: 0, VelX: 1, VelY: 0},
	}

	utils.MakeResponse(w, r, state)
}

func generateFruitPosition(width int, height int) Fruit {
	return Fruit{
		X: rand.Intn(width),
		Y: rand.Intn(height),
	}
}

func ValidateGame(w http.ResponseWriter, r *http.Request) {
	validations.IsPostMethod(w, r)

	//decode request body into game state and ticks
	var requestData struct {
		State State  `json:"state"`
		Ticks []Tick `json:"ticks"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	validateTicksAndUpdateGameState(w, r, &requestData)
	utils.MakeResponse(w, r, requestData.State)
}

func isValidMove(snake Snake, tick Tick) bool {
	if (snake.VelX != 0 && snake.VelX == -tick.VelX) || (snake.VelY != 0 && snake.VelY == -tick.VelY) {
		return false
	}
	return true
}

func validateOutOfBounds(w http.ResponseWriter, state State) {
	if state.Snake.X < 0 || state.Snake.X >= state.Width || state.Snake.Y < 0 || state.Snake.Y >= state.Height {
		http.Error(w, "Game is over, snake went out of bounds.", http.StatusTeapot)
		return
	}
}

func validateTicksAndUpdateGameState(w http.ResponseWriter, r *http.Request, requestData *struct {
	State State  `json:"state"`
	Ticks []Tick `json:"ticks"`
}) {
	for _, tick := range requestData.Ticks {
		//	validate the move
		if !isValidMove(requestData.State.Snake, tick) {
			http.Error(w, "Game is over, snake made an invalid move.", http.StatusTeapot)
			return
		}

		// update the snake's position and velocity
		requestData.State.Snake.X += tick.VelX
		requestData.State.Snake.Y += tick.VelY
		requestData.State.Snake.VelX += tick.VelX
		requestData.State.Snake.VelY += tick.VelY

		// check if the snake is out of bounds
		validateOutOfBounds(w, requestData.State)

		// check if the snake has reached the fruit
		if requestData.State.Snake.X == requestData.State.Fruit.X && requestData.State.Snake.Y == requestData.State.Fruit.Y {
			// increment score & generate new fruit position
			requestData.State.Score++
			requestData.State.Fruit = generateFruitPosition(requestData.State.Width, requestData.State.Height)
			break
		}
	}
}
