package handler

type CommentsCsv struct {
	Drive Drive `csv:"drive"`
	QuickRatings QuickRatings `csv:"quickRatings"`
	Advice Advice `csv:"advice"`
	ShotType ShotType `csv:"shotType"`
	LoadType LoadType `csv:"loadType"`
	TxNotes string `csv:"comm_txNotes"`
}

type DriveCsv struct {
	FlAssist bool `csv:"comm_flAssist"`
	IdDriveRating int `csv:"comm_idDriveRating"`
	IdDefenceRating int `csv:"comm_idDefenceRating"`
}

type QuickRatingsCsv struct {
	FlAlliance bool `csv:"comm_flAlliance"`
	FlStrategy bool `csv:"comm_flStrategy"`
	FlOwnThing bool `csv:"comm_flOwnThing"`
	FlRecovery bool `csv:"comm_flRecovery"`
}

type AdviceCsv struct {
	FlWarning bool `csv:"comm_flWarning"`
	FlHighlight bool `csv:"comm_flHighlight"`
}

type ShotTypeCsv struct {
	FlShotFar bool `csv:"comm_flShotFar"`
	FlShotMid bool `csv:"comm_flShotMid"`
	FlShotNear bool `csv:"comm_flShotNear"`
	FlShotWall  bool `csv:"comm_flShotWall"`
}

type LoadTypeCsv struct {
	FlIntakeGround bool `csv:"comm_flIntakeGround"`
	FlIntakeHigh bool `csv:"comm_flIntakeHigh"`
	FlIntakeRobot bool `csv:"comm_flIntakeRobot"`
}
