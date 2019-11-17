package iorpc

//JoinGameRequest join game request for iorung rpc service
type JoinGameRequest struct {
	GameID uint
	Token  string
}

//AddPlayerRequest add player request for iorung rpc service
type AddPlayerRequest struct {
	PlayerID uint
	GameID   uint
}

//AuthenticateResponse response on token authentication
type AuthenticateResponse struct {
	GameID   uint
	PlayerID uint
}

//DistributeCardsRequest rpc request protocol for distributing cards in a game
type DistributeCardsRequest struct {
	PlayerIds []uint
	GameID    uint
}

//LogoutRequest LogoutRequest
type LogoutRequest struct {
	Token string
}

type Player struct {
	Cards    []Card
	PlayerID uint
	GameID   uint
}
type Card struct {
	House  string
	Number int
}
