package iorpc

func (c *client) AddPlayer(req AddPlayerRequest) (string, error) {

	var token string
	client, err := c.Dial()
	if err != nil {
		return token, err
	}
	defer client.Close()

	err = client.Call("InterfaceRPC.AddPlayer", req, &token)
	if err != nil {
		return token, err
	}
	return token, nil

}
