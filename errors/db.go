package errors

import "fmt"

//ErrDBConnection ErrDBConnection
type ErrDBConnection struct {
	ConnectionString string
	baseErr          string
}

func (e *ErrDBConnection) Error() string {
	return fmt.Sprintf("error connecting %s. base error: %s", e.ConnectionString, e.baseErr)
}
