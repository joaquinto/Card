# Card Game

A card game is any game using playing cards as the primary device with which the game is played, be they traditional or game-specific. Countless card games exist, including families of related games (such as poker). A card game is played with a deck or pack of playing cards which are identical in size and shape.

This project implements the standard 52 deck of cards with suits comprising of Diamonds, Clubs, Spades and Hearts.

___

**Cloning this repository**

```bash
git clone https://github.com/joaquinto/Card.git
```
___

**Getting started**
```bash
go get ./...
cp .env.sample .env
#  Adjust the .env file if needed
go run cmd/main.go
```
___

**Testing the API**

You can locally run the test by running `go test -v ./...` in the cloned repository directory.
___

**Build binary file**

To build the card binary file, run `go build -o /build/card ./cmd/main.go` then run `./build/card` to run the program.
___

**Build docker image**

```bash
docker build ./build
```
___

**Features Implemented**
1. Create new deck.
2. Open deck.
3. draw a card.
___

**API Information**

|METHOD  |DESCRIPTION                        |ENDPOINT                                  |
|------- |-----------------------------------|------------------------------------------|
|POST    |Create new deck                            |/api/v1/decks?shuffled=true&cards=AS,KD,AC,2C,KH                      |
|GET    |Open deck                            |/api/v1/decks/:deckId                       |
|PATCH    |Draw a card           |/api/v1/decks/:deckId/draw-card?count=2                        |
