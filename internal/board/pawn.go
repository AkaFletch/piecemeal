package board

type Pawn interface {
	// Get the total pawn count of the game
	PawnCount() int

	// Get the pawn count of piecemeal
	PMPawnCount() int

	// Get the pawn count of the opponent
	OPPawnCount() int

	// Get the pawn bitboard for piecemeal
	PMPawnBitboard() int64

	// Get the pawn bitboard for the opponent
	OPPawnBitboard() int64
}

// Implementation of pawn pieces in the game
type pawn struct {
	pmPawnBitboard int64
	opPawnBitboard int64
}

// return piecemeals pawn bitboard
func (pawn pawn) PMPawnBitBoard() int64 {
	return pawn.pmPawnBitboard
}

// Return the opponent pawn bitboard
func (pawn pawn) OPPawnBitBoard() int64 {
	return pawn.opPawnBitboard
}
