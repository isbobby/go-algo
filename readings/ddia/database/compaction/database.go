package compaction

type appendOnlyDatabase struct {
	segments []segment
}

// needs to provide Write Lock
type segment struct {
	rows []row
}

type row struct {
	key   int
	value int
}
