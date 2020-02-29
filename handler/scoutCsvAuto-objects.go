package handler

type ScoutAutoCsv struct {
	Auto AutoCsv `json:"auto"`
	Errors ErrorsCsv `json:"errors"`
	Performance AutoPerformanceCsv `json:"performance"`
}

type AutoCsv struct {
	FlStart bool `json:"auto_flStart"`
	FlBaseLine bool `json:"auto_flBaseLine"`
	NumCellLoad int `json:"auto_numCellLoad"`
}

type ErrorsCsv struct {
	FlFoul bool `json:"auto_flFoul"`
	FlRobotContact bool `json:"auto_flRobotContact"`
	FlLoseStartObject bool `json:"auto_flLoseStartObject"`
	FlCrossover bool `json:"auto_flCrossover"`
}

type AutoPerformanceCsv struct {
	NumCellAttempt int `json:"auto_numCellAttempt"`
	NumCellSuccess int `json:"auto_numCellSuccess"`
	FlOuter bool `json:"auto_flOuter"`
	FlInner bool `json:"auto_flInner"`
	FlLower bool `json:"auto_flLower"`
}
