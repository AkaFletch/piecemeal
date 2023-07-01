package board

type King interface {
	// Get the king bitboard for piecemeal
	PMKingBitboard() int64

	// Get the king bitboard for the opponent
	OPKingBitboard() int64
}

type king struct {
	pmKingBitBoard int64
	opKingBitBoard int64
}

// Get piecemeals king bitboard
func (king king) PMKingBitBoard() int64 {
	return king.pmKingBitBoard
}

// Get the opponent king bitboard
func (king king) OPKingBitBoard() int64 {
	return king.opKingBitBoard
}
