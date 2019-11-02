package errors

import "fmt"

//ErrPlayerIDNotFound ErrPlayerIDNotFound
type ErrPlayerIDNotFound struct {
	PlayerID string
}

//ErrPlayerNotFound ErrPlayerNotFound
type ErrPlayerNotFound struct{}

func (e *ErrPlayerNotFound) Error() string {
	return fmt.Sprintf("player not found")
}

func (e *ErrPlayerIDNotFound) Error() string {
	return fmt.Sprintf("%s doesn't exist in db", e.PlayerID)
}

//ErrFailCreatePlayerInDb ErrFailCreatePlayerInDb
type ErrFailCreatePlayerInDb struct {
	Err error
}

func (e *ErrFailCreatePlayerInDb) Error() string {
	return e.Err.Error()
}
