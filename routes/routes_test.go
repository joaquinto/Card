package routes

import (
	"bytes"
	"card/config"
	"card/domain"
	"card/services"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

type testTable struct {
	name       string
	methods    string
	url        string
	statusCode int
}

func TestCreateEndpoints(t *testing.T) {
	cfg := config.LoadConfig()

	if cfg.Port == "" {
		cfg.Port = "8000"
	}

	e := echo.New()
	srv := httptest.NewServer(Run(e))
	defer srv.Close()

	services.Decks = []*domain.Deck{
		{ID: "414c59ae-51a9-4d15-a1bc-8983ccdeea19", Shuffled: false, Remaining: 52, Cards: services.CreateCard()},
		{ID: "4f84dabf-f4b6-4282-8ab9-d7a1a5cb7700", Shuffled: true, Remaining: 52, Cards: services.ShuffleCard(services.CreateCard())},
		{ID: "c7f83a96-6a16-46b3-84cb-4e4c09ce6505", Shuffled: false, Remaining: 5, Cards: services.MakeCardSelection(services.CreateCard(), []string{"AS", "KD", "AC", "2C", "KH"})},
		{ID: "28041656-40f3-46b0-bf91-4580fc7a8ccb", Shuffled: true, Remaining: 6, Cards: services.ShuffleCard(services.MakeCardSelection(services.CreateCard(), []string{"AS", "KD", "AC", "2C", "KH", "3C"}))},
		{ID: "62af6132-4f9b-4e13-b004-2a6f0d8573cd", Shuffled: true, Remaining: 5, Cards: services.ShuffleCard(services.MakeCardSelection(services.CreateCard(), []string{"AS", "KD", "AC", "2C", "KH"}))},
		{ID: "61e7bab6-0588-474d-ac3e-4b587f10f25d", Shuffled: false, Remaining: 6, Cards: services.MakeCardSelection(services.CreateCard(), []string{"AS", "KD", "AC", "2C", "KH", "3C"})},
		{ID: "ee9ade4b-0579-4a07-b9d9-cec720e5efd5", Shuffled: false, Remaining: 7, Cards: services.MakeCardSelection(services.CreateCard(), []string{"AS", "KD", "AC", "2C", "KH", "3C", "4H"})},
		{ID: "6ec0ce07-6b14-4643-a0d6-114a47025078", Shuffled: true, Remaining: 7, Cards: services.ShuffleCard(services.MakeCardSelection(services.CreateCard(), []string{"AS", "KD", "AC", "2C", "KH", "3C", "4H"}))},
		{ID: "2b67df25-94ca-438f-9f01-fec44035ff85", Shuffled: false, Remaining: 9, Cards: services.MakeCardSelection(services.CreateCard(), []string{"AS", "KD", "AC", "2C", "KH", "3C", "4H", "JC", "5D"})},
		{ID: "5232c187-3381-4980-b05b-384b44e2a381", Shuffled: true, Remaining: 9, Cards: services.ShuffleCard(services.MakeCardSelection(services.CreateCard(), []string{"AS", "KD", "AC", "2C", "KH", "3C", "4H", "JC", "5D"}))},
	}

	var tt = []*testTable{
		// Create a deck
		{name: "Create a shuffled deck of cards with success", methods: http.MethodPost, url: "/api/v1/decks?shuffled=true", statusCode: 201},
		{name: "Create a deck of card with error for invalid route/endpoint", methods: http.MethodPost, url: "/api/v1/deck", statusCode: 404},
		{name: "Create a shuffled and selected deck of cards with successful", methods: http.MethodPost, url: "/api/v1/decks?shuffled=true&cards=AS,KD,AC,2C,KH", statusCode: 201},
		{name: "Create a shuffled and selected deck of cards with error for invalid card code", methods: http.MethodPost, url: "/api/v1/decks?shuffled=true&cards=AS,KDH,AC,2C,KH", statusCode: 400},
		{name: "Create a shuffled and selected deck of cards with error for invalid shuffled value", methods: http.MethodPost, url: "/api/v1/decks?shuffled=tr&cards=AS,KD,AC,2C,KH", statusCode: 400},
		{name: "Create a non-shuffled and selected deck of cards with successful", methods: http.MethodPost, url: "/api/v1/decks?shuffled=false&cards=AS,KD,AC,2C,KH", statusCode: 201},
		{name: "Create a selected deck of cards with successful", methods: http.MethodPost, url: "/api/v1/decks?cards=AS,KD,AC,2C,KH", statusCode: 201},
		{name: "Create a deck of cards with success", methods: http.MethodPost, url: "/api/v1/decks", statusCode: 201},

		// Open Deck
		{name: "Open a deck with success", methods: http.MethodGet, url: "/api/v1/decks/4f84dabf-f4b6-4282-8ab9-d7a1a5cb7700", statusCode: 200},
		{name: "Open a deck with error for invalid deck id", methods: http.MethodGet, url: "/api/v1/decks/ae7558a5-b0a0-4f0e-9c6e-89d2134f0f40", statusCode: 404},
		{name: "Open a deck with error 404 for invalid route/endpoint", methods: http.MethodGet, url: "/api/v1/decks/", statusCode: 404},

		// Draw Card
		{name: "Draw card with success", methods: http.MethodPatch, url: "/api/v1/decks/28041656-40f3-46b0-bf91-4580fc7a8ccb/draw-card?count=2", statusCode: 200},
		{name: "Draw card with error for count less than 1", methods: http.MethodPatch, url: "/api/v1/decks/28041656-40f3-46b0-bf91-4580fc7a8ccb/draw-card?count=0", statusCode: 400},
		{name: "Draw card with error for not passing the count query parameter", methods: http.MethodPatch, url: "/api/v1/decks/28041656-40f3-46b0-bf91-4580fc7a8ccb/draw-card", statusCode: 400},
		{name: "Draw card with error 404 for invalid deck id", methods: http.MethodPatch, url: "/api/v1/decks/ae7558a5-b0a0-4f0e-9c6e-89d2134f0f40/draw-card?count=2", statusCode: 404},
		{name: "Draw card with error 404 for invalid route/endpoint", methods: http.MethodPatch, url: "/api/v1/decks/ae7558a5-b0a0-4f0e-9c6e-89d2134f0f40/draw", statusCode: 404},
	}

	for _, tr := range tt {
		t.Run(tr.name, func(t *testing.T) {
			uri := fmt.Sprintf("%s%s", srv.URL, tr.url)
			recbyte, _ := json.Marshal("")
			req, err := http.NewRequest(tr.methods, uri, bytes.NewReader(recbyte))
			if err != nil {
				t.Errorf("Could not create request: %v", err)
			}

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Errorf("Could not send request: %v", err)
			}

			if res.StatusCode != tr.statusCode {
				t.Errorf("Expected status %v; got %v", tr.statusCode, res.StatusCode)
			}
		})
	}

}
