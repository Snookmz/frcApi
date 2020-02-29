package handler

type Comments struct {
	Drive Drive `json:"drive"`
	QuickRatings QuickRatings `json:"quickRatings"`
	Advice Advice `json:"advice"`
	ShotType ShotType `json:"shotType"`
	LoadType LoadType `json:"loadType"`
	TxNotes string `json:"comm_txNotes"`
}

type Drive struct {
	FlAssist bool `json:"comm_flAssist"`
	IdDriveRating int `json:"comm_idDriveRating"`
	IdDefenceRating int `json:"comm_idDefenceRating"`
}

type QuickRatings struct {
	FlAlliance bool `json:"comm_flAlliance"`
	FlStrategy bool `json:"comm_flStrategy"`
	FlOwnThing bool `json:"comm_flOwnThing"`
	FlRecovery bool `json:"comm_flRecovery"`
}

type Advice struct {
	FlWarning bool `json:"comm_flWarning"`
	FlHighlight bool `json:"comm_flHighlight"`
}

type ShotType struct {
	FlShotFar bool `json:"comm_flShotFar"`
	FlShotMid bool `json:"comm_flShotMid"`
	FlShotNear bool `json:"comm_flShotNear"`
	FlShotWall  bool `json:"comm_flShotWall"`
}

type LoadType struct {
	FlIntakeGround bool `json:"comm_flIntakeGround"`
	FlIntakeHigh bool `json:"comm_flIntakeHigh"`
	FlIntakeRobot bool `json:"comm_flIntakeRobot"`
}
