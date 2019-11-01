package errors

import "fmt"

type ErrPlayerIDNotFound struct {
	PlayerID string
}

func (e *ErrPlayerIDNotFound) Error() string {
	return fmt.Sprintf("%s doesn't exist in db", e.PlayerID)
}

type ErrFailCreatePlayerInDb struct {
	Err error
}

func (e *ErrFailCreatePlayerInDb) Error() string {
	return e.Err.Error()
}
