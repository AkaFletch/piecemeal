package board

import (
	"errors"
)

var ErrOffBoard = errors.New("Provided position was off the board")
