package board

type Queen interface {
	// Get the total queen count of the game
	QueenCount() int

	// Get the queen count of piecemeal
	PMQueenCount() int

	// Get the queen count of the opponent
	OPQueenCount() int

	// Get the queen bitboard for piecemeal
	PMQueenBitboard() int64

	// Get the queen bitboard for the opponent
	OPQueenBitboard() int64
}

// Implementation of queen pieces in the game
type queen struct {
	pmQueenBitboard int64
	opQueenBitboard int64
}

// return piecemeals queen bitboard
func (queen queen) PMQueenBitBoard() int64 {
	return queen.pmQueenBitboard
}

// Return the opponent queen bitboard
func (queen queen) OPQueenBitBoard() int64 {
	return queen.opQueenBitboard
}
