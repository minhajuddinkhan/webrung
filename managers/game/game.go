package game

import (
	"fmt"
	"strconv"

	"github.com/minhajuddinkhan/webrung/entities"
	"github.com/minhajuddinkhan/webrung/iorpc"
	"github.com/minhajuddinkhan/webrung/store"
	"github.com/minhajuddinkhan/webrung/store/models"
)

//GameManager GameManager
type GameManager interface {
	CreateGame(player *entities.Player) (game *entities.Game, err error)
	GetGame(gameID string) (*entities.Game, error)
	JoinGame(player *entities.Player, game *entities.Game) error
}

type gameManager struct {
	store    store.Store
	ioclient iorpc.Client
}

//NewGameManager returns a new game manager.
func NewGameManager(s store.Store, c iorpc.Client) GameManager {
	return &gameManager{store: s, ioclient: c}
}

func (g *gameManager) CreateGame(pl *entities.Player) (*entities.Game, error) {

	player := models.Player{Name: pl.Name}
	if err := player.SetID(pl.ID); err != nil {
		return nil, err
	}

	gameID, err := g.store.CreateGame(&player)
	if err != nil {
		return nil, err
	}
	gID, _ := strconv.Atoi(gameID)

	if err := g.store.JoinGame(&models.PlayersInGame{
		GameID:   uint(gID),
		PlayerID: player.ID,
	}); err != nil {
		return nil, err
	}

	updatedGame, err := g.store.GetGame(gameID)
	if err != nil {
		return nil, err
	}

	req := iorpc.AddPlayerRequest{PlayerID: pl.ID, GameID: gameID}
	if _, err := g.ioclient.AddPlayer(req); err != nil {
		return nil, err
	}
	return &entities.Game{
		GameID:        gameID,
		PlayersJoined: updatedGame.PlayersJoined,
	}, nil

}

func (g *gameManager) GetGame(gameID string) (*entities.Game, error) {
	game, err := g.store.GetGame(gameID)
	if err != nil {
		return nil, err

	}
	gameEntity := entities.Game{}
	gameEntity.Marshal(game)
	return &gameEntity, nil

}

func (g *gameManager) JoinGame(player *entities.Player, game *entities.Game) error {

	gameModel, err := game.ToModel()
	if err != nil {
		return err
	}
	playerModel, err := player.ToModel()
	if err != nil {
		return err
	}

	currentGame, err := g.store.GetGame(game.GameID)
	if err != nil {
		return err
	}

	if currentGame.PlayersJoined == 4 {
		return fmt.Errorf("already four players joined")
	}

	present, err := g.store.IsPlayerInGame(game.GameID, player.ID)
	if err != nil {
		return err
	}
	if present {
		return fmt.Errorf("player already present in game. cannot join again")
	}

	gameplay := models.PlayersInGame{
		Game:   *gameModel,
		Player: *playerModel,
	}

	if err := g.store.JoinGame(&gameplay); err != nil {
		return err
	}

	return g.store.IncrementPlayersJoined(game.GameID)
}
