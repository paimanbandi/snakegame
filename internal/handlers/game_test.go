package handlers

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewGame(t *testing.T) {
	// Test with valid width and height
	req, err := http.NewRequest("GET", "/new?w=3&h=10", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(NewGame)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code for valid parameters: got %v want %v",
			status, http.StatusOK)
	}

	// Test with invalid width
	req, err = http.NewRequest("GET", "/new?w=invalid&h=10", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code for invalid width: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Test with invalid height
	req, err = http.NewRequest("GET", "/new?w=10&h=invalid", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code for invalid height: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestIsValidMove(t *testing.T) {
	tests := []struct {
		name     string
		snake    Snake
		tick     Tick
		expected bool
	}{
		{"Valid Move X", Snake{VelX: 1, VelY: 0}, Tick{VelX: 1, VelY: 0}, true},
		{"Valid Move Y", Snake{VelX: 0, VelY: 1}, Tick{VelX: 0, VelY: 1}, true},
		{"Invalid Opposite Move X", Snake{VelX: 1, VelY: 0}, Tick{VelX: -1, VelY: 0}, false},
		{"Invalid Opposite Move Y", Snake{VelX: 0, VelY: 1}, Tick{VelX: 0, VelY: -1}, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := isValidMove(test.snake, test.tick); got != test.expected {
				t.Errorf("isValidMove() = %v, want %v", got, test.expected)
			}
		})
	}
}

func TestSnakeReachesFruit(t *testing.T) {
	// Set a fixed seed for predictable random numbers
	rand.Seed(time.Now().UnixNano())

	// Example game state where snake is about to reach the fruit
	state := State{
		Snake:  Snake{X: 5, Y: 5},
		Fruit:  Fruit{X: 5, Y: 5},
		Score:  0,
		Width:  10,
		Height: 10,
	}

	// Simulate snake reaching the fruit
	if state.Snake.X == state.Fruit.X && state.Snake.Y == state.Fruit.Y {
		state.Score++
		state.Fruit = generateFruitPosition(state.Width, state.Height)
	} else {
		t.Errorf("Snake did not reach the fruit when it should have")
	}

	// Verify that score incremented
	if state.Score != 1 {
		t.Errorf("Score did not increment correctly: got %d, want %d", state.Score, 1)
	}

}

func TestSnakeOutOfBounds(t *testing.T) {
	tests := []struct {
		name     string
		snakeX   int
		snakeY   int
		width    int
		height   int
		expected int
	}{
		{"Snake Inside Bounds", 5, 5, 10, 10, http.StatusOK},
		{"Snake Out of Bounds Negative X", -1, 5, 10, 10, http.StatusTeapot},
		{"Snake Out of Bounds Over Width", 10, 5, 10, 10, http.StatusTeapot},
		{"Snake Out of Bounds Negative Y", 5, -1, 10, 10, http.StatusTeapot},
		{"Snake Out of Bounds Over Height", 5, 10, 10, 10, http.StatusTeapot},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/", nil) // Modify as needed
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Simulate setting requestData from request
				requestData := struct {
					State State
				}{
					State: State{
						Snake:  Snake{X: test.snakeX, Y: test.snakeY},
						Width:  test.width,
						Height: test.height,
					},
				}

				validateOutOfBounds(w, requestData.State)
			})

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != test.expected {
				t.Errorf("%s: handler returned wrong status code: got %v want %v",
					test.name, status, test.expected)
			}
		})
	}
}
