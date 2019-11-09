package iorpc

func (c *client) SetGameIDInToken(request JoinGameRequest) (bool, error) {

	client, err := c.Dial()
	if err != nil {
		return false, err
	}
	defer client.Close()
	var done bool
	err = client.Call("InterfaceRPC.SetGameIDInToken", request, &done)
	return done, nil

}
func (c *client) StartGame(request DistributeCardsRequest) (*DistributeCardsResponse, error) {
	client, err := c.Dial()
	if err != nil {
		return nil, err
	}
	defer client.Close()
	var resp DistributeCardsResponse
	err = client.Call("InterfaceRPC.DistributeCards", request, &resp)

	return &resp, err
}
