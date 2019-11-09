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
func (c *client) StartGame(request DistributeCardsRequest) (bool, error) {
	client, err := c.Dial()
	var started bool
	if err != nil {
		return started, err
	}
	defer client.Close()
	err = client.Call("InterfaceRPC.DistributeCards", request, &started)

	return started, err
}
