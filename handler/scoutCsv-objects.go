package handler


type ScoutCsv struct {
	ParentData ParentDataCsv `csv:"parentData"`
	Auto ScoutAutoCsv `csv:"auto"`
	Tele TeleCsv `csv:"tele"`
	Comments Comments `csv:"comments"`
}

type ParentDataCsv struct {
	TxEvent string `csv:"txEvent"`
	DeviceName string `csv:"deviceName"`
	TeamDetails TeamDetailsCsv `csv:"teamDetails"`
	MatchSetup MatchSetupCsv `csv:"matchSetup"`
	Results ResultsCsv `csv:"results"`
}

type TeamDetailsCsv struct {
	NumMatch int `csv:"numMatch"`
	IdAlliance int `csv:"idAlliance"`
	IdDriveStation int `csv:"idDriveStation"`
	IdTeam string `csv:"idTeam"`
	TxScoutName string `csv:"txScoutName"`
}

type MatchSetupCsv struct {
	IdStartFacing int `csv:"idStartFacing"`
	IdStartPosition int `csv:"idStartPosition"`
	NumStartCells int `csv:"numStartCells"`
}

type ResultsCsv struct {
	FlRed bool `csv:"flRed"`
	FlYellow bool `csv:"flYellow"`
	FlCrash bool `csv:"flCrash"`
	FlRanking1 bool `csv:"flRanking1"`
	FlRanking2 bool `csv:"flRanking2"`
}
