package board

type Knight interface {
	// Get the total knight count of the game
	KnightCount() int

	// Get the knight count of piecemeal
	PMKnightCount() int

	// Get the knight count of the opponent
	OPKnightCount() int

	// Get the knight bitboard for piecemeal
	PMKnightBitboard() int64

	// Get the knight bitboard for the opponent
	OPKnightBitboard() int64
}

// Implementation of knight pieces in the game
type knight struct {
	pmKnightBitboard int64
	opKnightBitboard int64
}

// return piecemeals knight bitboard
func (knight knight) PMKnightBitBoard() int64 {
	return knight.pmKnightBitboard
}

// Return the opponent knight bitboard
func (knight knight) OPKnightBitBoard() int64 {
	return knight.opKnightBitboard
}
