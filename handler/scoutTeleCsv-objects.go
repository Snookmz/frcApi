package handler

type TeleCsv struct {
	ControlPanel ControlPanel `csv:"controlPanel"`
	EndGame EndGame `csv:"endGame"`
	Performance TelePerformance `csv:"performance"`
}

type ControlPanelCsv struct {
	FlPanelRotation bool `csv:"tele_flPanelRotation"`
	IdPanelRotationTime int `csv:"tele_idPanelRotationTime"`
	FlPanelPosition bool `csv:"tele_flPanelPosition"`
	IdPanelPositionTime int `csv:"tele_idPanelPositionTime"`
	NumPanelAttempt int `csv:"tele_numPanelAttempt"`
	NumPanelSuccess int `csv:"tele_numPanelSuccess"`
}

type EndGameCsv struct {
	FlPark bool `csv:"tele_tele_flPark"`
	IdClimb int `csv:"tele_idClimb"`
	IdClimbGrabTime int `csv:"tele_idClimbGrabTime"`
	IdClimbTime int `csv:"tele_idClimbTime"`
	IdClimbOutcome int `csv:"tele_idClimbOutcome"`
	IdClimbPos int `csv:"tele_idClimbPos"`
	NumClimbOthers int `csv:"tele_numClimbOthers"`
	FlClimbBalance bool `csv:"tele_flClimbBalance"`
	FlClimbCorrection bool `csv:"tele_flClimbCorrection"`
	FlClimbFall bool `csv:"tele_flClimbFall"`
}

type TelePerformanceCsv struct {
	NumCellAttempt int `csv:"tele_numCellAttempt"`
	NumCellSuccess int `csv:"tele_numCellSuccess"`
	FlOuter bool `csv:"tele_flOuter"`
	FlInner bool `csv:"tele_flInner"`
	FlLower bool `csv:"tele_flLower"`
}
