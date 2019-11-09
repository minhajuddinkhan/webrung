package mock

import (
	"net/rpc"

	"github.com/minhajuddinkhan/webrung/iorpc"
)

type mockRPC struct{}

//NewMockIORPCClient NewMockIORPCClient
func NewMockIORPCClient() iorpc.Client {
	return &mockRPC{}
}

func (mrpc *mockRPC) Dial() (*rpc.Client, error) {
	return nil, nil
}
func (mrpc *mockRPC) SetGameIDInToken(request iorpc.JoinGameRequest) (done bool, err error) {
	return true, nil
}
func (mrpc *mockRPC) AddPlayer(request iorpc.AddPlayerRequest) (token string, err error) {
	return "mock.jwt.token", nil
}
func (mrpc *mockRPC) Authenticate(token string) (gameID, playerID string, err error) {
	return "1", "1", nil
}

func (mrpc *mockRPC) StartGame(request iorpc.DistributeCardsRequest) (*iorpc.DistributeCardsResponse, error) {
	return &iorpc.DistributeCardsResponse{
		Players: make([]iorpc.Player, 4),
		GameID:  "1",
	}, nil
}
