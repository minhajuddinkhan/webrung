package iorpc

func (c *client) Authenticate(token string) (gameID, playerID uint, err error) {

	client, err := c.Dial()
	if err != nil {
		return 0, 0, err
	}
	defer client.Close()
	var authResponse AuthenticateResponse
	err = client.Call("InterfaceRPC.Authenticate", token, &authResponse)
	return authResponse.GameID, authResponse.PlayerID, err
}
