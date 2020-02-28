package handler

type Tele struct {
	ControlPanel ControlPanel `json:"controlPanel"`
	EndGame EndGame `json:"endGame"`
	Performance TelePerformance `json:"performance"`
}

type ControlPanel struct {
	FlPanelRotation bool `json:"tele_flPanelRotation"`
	IdPanelRotationTime int `json:"tele_idPanelRotationTime"`
	FlPanelPosition bool `json:"tele_flPanelPosition"`
	IdPanelPositionTime int `json:"tele_idPanelPositionTime"`
	NumPanelAttempt int `json:"tele_numPanelAttempt"`
	NumPanelSuccess int `json:"tele_numPanelSuccess"`
}

type EndGame struct {
	FlPark bool `json:"tele_tele_flPark"`
	IdClimb int `json:"tele_idClimb"`
	IdClimbGrabTime int `json:"tele_idClimbGrabTime"`
	IdClimbTime int `json:"tele_idClimbTime"`
	IdClimbOutcome int `json:"tele_idClimbOutcome"`
	IdClimbPos int `json:"tele_idClimbPos"`
	NumClimbOthers int `json:"tele_numClimbOthers"`
	FlClimbBalance bool `json:"tele_flClimbBalance"`
	FlClimbCorrection bool `json:"tele_flClimbCorrection"`
	FlClimbFall bool `json:"tele_flClimbFall"`
}

type TelePerformance struct {
	NumCellAttempt int `json:"tele_numCellAttempt"`
	NumCellSuccess int `json:"tele_numCellSuccess"`
	FlOuter bool `json:"tele_flOuter"`
	FlInner bool `json:"tele_flInner"`
	FlLower bool `json:"tele_flLower"`
}
