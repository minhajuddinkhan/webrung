package iorpc

func (c *client) SetGameIDInToken(request JoinGameRequest) (bool, error) {

	client, err := c.Dial()
	if err != nil {
		return false, err
	}
	defer client.Close()
	var token string
	err = client.Call("InterfaceRPC.SetGameIDInToken", request, &token)
	if err != nil {
		return false, err
	}
	return true, nil

}
