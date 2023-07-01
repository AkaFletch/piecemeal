package board

import "testing"

func TestPMSetAt(t *testing.T) {
	for _, test := range SetPositionTestCases {
		t.Run(test.testName, func(t *testing.T) {
			rook := rook{}
			pmErr := rook.PMSetAt(test.file, test.rank)
			if pmErr != nil {
				t.Errorf("Unexpected err: %q", pmErr)
			}
			if rook.pmRookBitboard != test.expectedOutput {
				t.Errorf(
					"Piece was not set in correct bitboard location, expected %064b got %064b",
					test.expectedOutput,
					rook.PMRookBitboard(),
				)
			}
			opErr := rook.OPSetAt(test.file, test.rank)
			if opErr != nil {
				t.Errorf("Unexpected err: %q", pmErr)
			}
			if rook.opRookBitboard != test.expectedOutput {
				t.Errorf(
					"Piece was not set in correct bitboard location, expected %064b got %064b",
					test.expectedOutput,
					rook.OPRookBitboard(),
				)
			}
		})
	}
}
