package handler

type Pit struct {
	ImperialUnits bool `json:"imperialUnits"`
	Event string `json:"event"`
	Details Details `json:"details"`
	RobotStats RobotStats `json:"robotStats"`
	PowerCells PowerCells `json:"powerCells"`
	Climb Climb `json:"climb"`
	ControlPanel PitControlPanel `json:"controlPanel"`
	Auto PitAuto `json:"auto"`
	TxPitNotes string `json:"txPitNotes"`
	Record Record `json:"record"`

}

type Details struct {
	IdTeam int`json:"idTeam"`
	Name string`json:"name"`
	TxScoutName string`json:"txScoutName"`
}

type RobotStats struct {
	NumWeight int`json:"numWeight"`
	NumHeight int`json:"numHeight"`
	ImgTeamUniform string`json:"imgTeamUniform"`
	ImgRobotFront string`json:"imgRobotFront"`
	ImgRobotSide string`json:"imgRobotSide"`
}

type PowerCells struct {
	FlCells bool`json:"flCells"`
	FlIntakeGround bool`json:"flIntakeGround"`
	FlIntakeHigh bool`json:"flIntakeHigh"`
	NumStorage int`json:"numStorage"`
	TxShooting string`json:"txShooting"`
	FlTargetLow bool`json:"flTargetLow"`
	FlTargetOuter bool`json:"flTargetOuter"`
	FlTargetInner bool`json:"flTargetInner"`

}

type Climb struct {
	FlClimb bool`json:"flClimb"` // can climb
	IdClimbType int`json:"idClimbType"`
	NumClimbHeight int`json:"numClimbHeight"`
	FlClimbSecure bool`json:"flClimbSecure"`
	IdClimbGrab int`json:"idClimbGrab"` // 1=NA, 2=slow, 3=Med, 4=fast
	IdClimbSpeed int`json:"idClimbSpeed"` // same as above
	FlClimbTilt bool`json:"flClimbTilt"`
	TxClimb string`json:"txClimb"`
	IdClimbPos int`json:"idClimbPos"` // preferred position 1=NA, 2=Any, 3=Inner, 4=Middle, 5=Outer
	FlClimbLevel bool`json:"flClimbLevel"` // can level generator
	FlClimbLevelSelf bool`json:"flClimbLevelSelf"`
	FlClimbLevelOther bool`json:"flClimbLevelOther"`
	FlClimbMove bool`json:"flClimbMove"`
	FlClimbOther bool`json:"flClimbOther"`
	NumClimbOther int`json:"numClimbOther"` // buddies #
}

type PitControlPanel struct {
	FlPanel bool`json:"flPanel"` // can manipulate control panel
	FlPanelBrake bool`json:"flPanelBrake"`
	FlPanelRotation bool`json:"flPanelRotation"`
	FlPanelPos bool`json:"flPanelPos"`
	FlPanelSensor bool`json:"flPanelSensor"`
	TxPanelSensor string`json:"txPanelSensor"` // notes

}

type PitAuto struct {
	FlAuto bool`json:"flAuto"` // can auto
	FlAutoLine bool`json:"flAutoLine"`
	FlAutoShoot bool`json:"flAutoShoot"` // can auto shoot
	NumAutoShoot int`json:"numAutoShoot"` // int of balls
	NumAutoLoad int`json:"numAutoLoad"` // pickup

}

type Record struct {
	DtCreated string`json:"dtCreated"`
	DtModified string`json:"dtModified"`
	TxComputerName string`json:"txComputerName"`
}

type FrcEvent struct {
	Key string `json:"key"`
	Name string `json:"name"`
	EventCode string `json:"event_code"`
	EventType int `json:"event_type"`
	//district District `json:"district"`
	City string `json:"city"`
	StateProv string `json:"state_prov"`
	Country string `json:"country"`
	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
	Year int `json:"year"`
	ShortName string `json:"short_name"`
	EventTypeString string `json:"event_type_string"`
	Week int `json:"week"`
	Address string `json:"address"`
	PostalCode string `json:"postal_code"`
	GmapsPlaceId string `json:"gmaps_place_id"`
	GmapsUrl string `json:"gmaps_url"`
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
	LocationName string `json:"location_name"`
	Timezone string `json:"timezone"`
	Website string `json:"website"`
	FirstEventId string `json:"first_event_id"`
	FirstEventCode string `json:"first_event_code"`
	//webcasts Webcast[] `json:""`
	DivisionKeys []string `json:"division_keys"`
	ParentEventKey string `json:"parent_event_key"`
	PlayoffType int `json:"playoff_type"`
	PlayoffTypeString string `json:"playoff_type_string"`
}

