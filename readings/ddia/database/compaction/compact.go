package compaction

func fillRows() {

}

func createAppendOnlyDatabase() {

}

// use a cron to periodically compact segments
func compactSegments(s1, s2 *segment) segment {
	combinedRows := make([]row, len(s1.rows)+len(s2.rows))
	copy(combinedRows[:len(s1.rows)], s1.rows)
	copy(combinedRows[len(s1.rows):], s2.rows)

	compactedRows := make([]row, len(combinedRows))
	inserted := map[int]bool{}
	newSize := 1
	for i := len(combinedRows) - 1; i > 0; i-- {
		currentRow := combinedRows[i]

		_, exists := inserted[currentRow.key]
		if exists {
			continue
		}

		inserted[currentRow.key] = true
		compactedRows[len(combinedRows)-newSize] = row{currentRow.key, currentRow.value}
		newSize += 1
	}

	return segment{
		rows: compactedRows[len(combinedRows)-newSize+1:],
	}
}
