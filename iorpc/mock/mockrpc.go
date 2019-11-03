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
	return "some-dummy-token", nil
}
func (mrpc *mockRPC) Authenticate(token string) (gameID, playerID string, err error) {
	return "1", "1", nil
}
