package handler


type Scout struct {
	ParentData ParentData `json:"parentData"`
	Auto ScoutAuto `json:"auto"`
	Tele Tele `json:"tele"`
	Comments Comments `json:"comments"`
}

type ParentData struct {
	TxEvent string `json:"txEvent"`
	DeviceName string `json:"deviceName"`
	TeamDetails TeamDetails `json:"teamDetails"`
	MatchSetup MatchSetup `json:"matchSetup"`
	Results Results `json:"results"`
}

type TeamDetails struct {
	NumMatch int `json:"numMatch"`
	IdAlliance int `json:"idAlliance"`
	IdDriveStation int `json:"idDriveStation"`
	IdTeam string `json:"idTeam"`
	TxScoutName string `json:"txScoutName"`
}

type MatchSetup struct {
	IdStartFacing int `json:"idStartFacing"`
	IdStartPosition int `json:"idStartPosition"`
	NumStartCells int `json:"numStartCells"`
}

type Results struct {
	FlRed bool `json:"flRed"`
	FlYellow bool `json:"flYellow"`
	FlCrash bool `json:"flCrash"`
	FlRanking1 bool `json:"flRanking1"`
	FlRanking2 bool `json:"flRanking2"`
}
