package player

import (
	"github.com/minhajuddinkhan/webrung/entities"
	"github.com/minhajuddinkhan/webrung/store"
	"github.com/minhajuddinkhan/webrung/store/models"
)

//Manager player manager
type Manager interface {
	//CreatePlayer creates a new player
	CreatePlayer(player *entities.Player) (*entities.Player, error)

	//GetPlayer gets a player by id
	GetPlayer(playerID string) (*entities.Player, error)
}

type playerManager struct {
	store store.Store
}

//NewPlayerManager NewPlayerManager
func NewPlayerManager(store store.Store) Manager {
	return &playerManager{store: store}
}

func (m *playerManager) CreatePlayer(player *entities.Player) (*entities.Player, error) {

	p := models.Player{Name: player.Name}
	playerID, err := m.store.CreatePlayer(&p)
	if err != nil {
		return nil, err
	}
	player.ID = playerID
	return player, nil
}

func (m *playerManager) GetPlayer(playerID string) (*entities.Player, error) {
	player, err := m.store.GetPlayer(playerID)
	if err != nil {
		return nil, err
	}
	return &entities.Player{
		ID:   playerID,
		Name: player.Name,
	}, nil
}
