package iorpc

//JoinGameRequest join game request for iorung rpc service
type JoinGameRequest struct {
	GameID string
	Token  string
}

//AddPlayerRequest add player request for iorung rpc service
type AddPlayerRequest struct {
	PlayerID string
	GameID   string
}

//AuthenticateResponse response on token authentication
type AuthenticateResponse struct {
	GameID   string
	PlayerID string
}

//DistributeCardsRequest rpc request protocol for distributing cards in a game
type DistributeCardsRequest struct {
	PlayerIds []string
	GameID    string
}

type Player struct {
	Cards    []Card
	PlayerID string
	GameID   string
}
type Card struct {
	House  string
	Number int
}
