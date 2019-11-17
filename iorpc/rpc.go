package iorpc

import (
	"fmt"
	"net/rpc"

	"github.com/minhajuddinkhan/webrung/config"
)

//Client Client
type Client interface {
	Dial() (*rpc.Client, error)
	SetGameIDInToken(request JoinGameRequest) (done bool, err error)
	AddPlayer(request AddPlayerRequest) (token string, err error)
	Authenticate(token string) (gameID, playerID uint, err error)
	StartGame(request DistributeCardsRequest) (bool, error)
}

//NewIOClient client to communicate with the IORung server
func NewIOClient(conf *config.Conf) (Client, error) {
	return &client{
		protocol: "tcp",
		connStr:  fmt.Sprintf("%s:%s", conf.IORung.Host, conf.IORung.Port),
	}, nil
}

type client struct {
	connStr  string
	protocol string
}

func (c *client) Dial() (*rpc.Client, error) {
	return rpc.DialHTTP(c.protocol, c.connStr)
}
