package errors

import "fmt"

//ErrGameIDNotFound ErrGameIDNotFound
type ErrGameIDNotFound struct {
	GameID uint
}

func (e *ErrGameIDNotFound) Error() string {
	return fmt.Sprintf("%d doesn't exist in db", e.GameID)
}

//ErrFailCreateGameInDb ErrFailCreateGameInDb
type ErrFailCreateGameInDb struct {
	Err error
}

func (e *ErrFailCreateGameInDb) Error() string {
	return e.Err.Error()
}

//ErrGameAlreadyHosted ErrGameAlreadyHosted
type ErrGameAlreadyHosted struct {
	Err error
}

func (e *ErrGameAlreadyHosted) Error() string {
	if e.Err == nil {
		return fmt.Sprintf("game already hosted")
	}
	return e.Err.Error()
}

//ErrPlayerAlreadyJoinedInAnotherGame ErrPlayerAlreadyJoinedInAnotherGame
type ErrPlayerAlreadyJoinedInAnotherGame struct {
	Err error
}

func (e *ErrPlayerAlreadyJoinedInAnotherGame) Error() string {
	if e.Err == nil {
		return fmt.Sprintf("cannot create new game while you are already joined in another game")
	}
	return e.Err.Error()
}
