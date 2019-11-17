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

func (c *client) Logout(token string) (bool, error) {

	client, err := c.Dial()
	if err != nil {
		return false, err
	}
	defer client.Close()
	var logoutResponse bool
	err = client.Call("InterfaceRPC.Logout", LogoutRequest{Token: token}, &logoutResponse)
	return logoutResponse, err
}
