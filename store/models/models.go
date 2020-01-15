package models

func Models() []interface{} {

	return []interface{}{

		Card{},
		Game{},
		PlayersInGame{},
		JoinableGame{},
		Player{},
	}
}
