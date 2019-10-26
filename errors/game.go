package errors

import "fmt"

//ErrGameIDNotFound ErrGameIDNotFound
type ErrGameIDNotFound struct {
	GameID string
}

func (e *ErrGameIDNotFound) Error() string {
	return fmt.Sprintf("%s doesn't exist in db", e.GameID)
}

//ErrFailCreateGameInDb ErrFailCreateGameInDb
type ErrFailCreateGameInDb struct {
	Err error
}

func (e *ErrFailCreateGameInDb) Error() string {
	return e.Err.Error()
}
