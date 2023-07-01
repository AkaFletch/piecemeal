package board

// Functions around bishops in a given board
type Bishop interface {
	// Get the total bishop count of the game
	BishopCount() int

	// Get the bishop count of piecemeal
	PMBishopCount() int

	// Get the bishop count of the opponent
	OPBishopCount() int

	// Get the bishop bitboard for piecemeal
	PMBishopBitboard() int64

	// Get the bishop bitboard for the opponent
	OPBishopBitboard() int64
}

// Implementation of bishop pieces in the game
type bishop struct {
	pmBishopBitboard int64
	opBishopBitboard int64
}

// return piecemeals bishop bitboard
func (bishop bishop) PMBishopBitBoard() int64 {
	return bishop.pmBishopBitboard
}

// Return the opponent bishop bitboard
func (bishop bishop) OPBishopBitBoard() int64 {
	return bishop.opBishopBitboard
}
