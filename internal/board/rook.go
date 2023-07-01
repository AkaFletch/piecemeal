package board

type Rook interface {
	// Get the total rook count of the game
	RookCount() int

	// Get the rook count of piecemeal
	PMRookCount() int

	// Get the rook count of the opponent
	OPRookCount() int

	// Get the rook bitboard for piecemeal
	PMRookBitboard() uint64

	// Get the rook bitboard for the opponent
	OPRookBitboard() uint64

	// Set PM rook at a given position
	PMSetAt(file int, rank int)

	// Set opponent rook at a given position
	OPSetAt(file int, rank int)
}

// Implementation of rook pieces in the game
type rook struct {
	pmRookBitboard uint64
	opRookBitboard uint64
}

// return piecemeals rook bitboard
func (rook rook) PMRookBitboard() uint64 {
	return rook.pmRookBitboard
}

// Return the opponent rook bitboard
func (rook rook) OPRookBitboard() uint64 {
	return rook.opRookBitboard
}

// Forcefully set a rook for PM's colour into a square
func (rook *rook) PMSetAt(file int, rank int) error {
	if file > 7 || rank > 7 || file < 0 || rank < 0 {
		return ErrOffBoard
	}
	bitSet := uint64(1 << ((8 * rank) + file))
	rook.pmRookBitboard |= bitSet
	return nil
}

// Forcefully set a rook for the opponents colour into a square
func (rook *rook) OPSetAt(file int, rank int) error {
	if file > 7 || rank > 7 || file < 0 || rank < 0 {
		return ErrOffBoard
	}
	bitSet := uint64(1 << ((8 * rank) + file))
	rook.opRookBitboard |= bitSet
	return nil
}
