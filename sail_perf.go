package sail_perf

// Keep track of performance samples
type ScoreBoard struct {
	samples [][]float64
}

func New(resolution int, samples int) ScoreBoard {
	return ScoreBoard{make([][]float64, resolution)}
}

// Add a sample to the scoreboard (if it is good enough)
func AddSample(board ScoreBoard, heading int, performance float64) {
	target := board.samples[heading]
	if target == nil {

	}
}

