package types

import "fmt"

type TooManyTasksError struct {
	Count int
}

func (e *TooManyTasksError) Error() string {
	return fmt.Sprintf("too many tasks: %d", e.Count)
}
