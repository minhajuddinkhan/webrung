package auth

import (
	"github.com/minhajuddinkhan/webrung/iorpc"
	"github.com/minhajuddinkhan/webrung/store"
)

//Manager manages authentication
type Manager interface {
	//Authenticate authenticates a player
	Authenticate(username string) (token string, err error)
}

type manager struct {
	ioclient iorpc.Client
	store    store.Store
}

//NewAuthManager returns new auth manager
func NewAuthManager(iorungrpc iorpc.Client, store store.Store) Manager {
	return &manager{ioclient: iorungrpc, store: store}
}
