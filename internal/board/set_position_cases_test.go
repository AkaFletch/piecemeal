package board

type SetPositionTestCase struct {
	testName       string
	file           int
	rank           int
	expectedOutput uint64
}

var SetPositionTestCases = []SetPositionTestCase{
	{
		testName:       "Set top left to piece",
		file:           7,
		rank:           7,
		expectedOutput: 1 << 63,
	},
	{
		testName:       "Set top right to piece",
		file:           0,
		rank:           7,
		expectedOutput: 1 << 56,
	},
	{
		testName:       "Set bottom left to piece",
		file:           0,
		rank:           0,
		expectedOutput: 1,
	},
	{
		testName:       "Set bottom right to piece",
		file:           7,
		rank:           0,
		expectedOutput: 1 << 7,
	},
	{
		testName:       "Set rank 1 file 0",
		file:           0,
		rank:           1,
		expectedOutput: 1 << 8,
	},
	{
		testName:       "Set rank 7 file 0",
		file:           0,
		rank:           7,
		expectedOutput: 1 << (8 * 7),
	},
}
