package handler

type ScoutAuto struct {
	Auto Auto `json:"auto"`
	Errors Errors `json:"errors"`
	Performance AutoPerformance `json:"performance"`
}

type Auto struct {
	FlStart bool `json:"auto_flStart"`
	FlBaseLine bool `json:"auto_flBaseLine"`
	NumCellLoad int `json:"auto_numCellLoad"`
}

type Errors struct {
	FlFoul bool `json:"auto_flFoul"`
	FlRobotContact bool `json:"auto_flRobotContact"`
	FlLoseStartObject bool `json:"auto_flLoseStartObject"`
	FlCrossover bool `json:"auto_flCrossover"`
}

type AutoPerformance struct {
	NumCellAttempt int `json:"auto_numCellAttempt"`
	NumCellSuccess int `json:"auto_numCellSuccess"`
	FlOuter bool `json:"auto_flOuter"`
	FlInner bool `json:"auto_flInner"`
	FlLower bool `json:"auto_flLower"`
}
