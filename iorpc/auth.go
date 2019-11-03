package iorpc

func (c *client) Authenticate(token string) (gameID, playerID string, err error) {

	client, err := c.Dial()
	if err != nil {
		return "", "", err
	}
	defer client.Close()
	var authResponse AuthenticateResponse
	err = client.Call("InterfaceRPC.Authenticate", token, &authResponse)
	return authResponse.GameID, authResponse.PlayerID, err
}
