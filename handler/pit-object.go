package handler

type Pit struct {
	ImperialUnits bool `json:"imperialUnits"`
	Event FrcEvent `json:"event"`
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
	IdTeam string`json:""`
	Name string`json:""`
	TxScoutName string`json:""`

}

export class RobotStats {
numWeight number`json:""`
numHeight number`json:""`
imgTeamUniform string`json:""`
imgRobotFront string`json:""`
imgRobotSide string`json:""`
constructor() {
this.numWeight = 0`json:""`
this.numHeight = 0`json:""`
this.imgTeamUniform = ''`json:""`
this.imgRobotFront = ''`json:""`
this.imgRobotSide = ''`json:""`
}
}

export class PowerCells {
flCells boolean`json:""` // manipulate
flIntakeGround boolean`json:""` // ground intake
flIntakeHigh boolean`json:""`
numStorage number`json:""`
txShooting string`json:""`
flTargetLow boolean`json:""`
flTargetOuter boolean`json:""`
flTargetInner boolean`json:""`
constructor() {
this.flCells = false`json:""`
this.flIntakeGround = false`json:""`
this.flIntakeHigh = false`json:""`
this.numStorage = 0`json:""`
this.txShooting = ''`json:""`
this.flTargetLow = false`json:""`
this.flTargetOuter = false`json:""`
this.flTargetInner = false`json:""`
}
}

export class Climb {
flClimb boolean`json:""` // can climb
idClimbType string`json:""`
numClimbHeight number`json:""`
flClimbSecure boolean`json:""`
idClimbGrab number`json:""` // 1=NA, 2=slow, 3=Med, 4=fast
idClimbSpeed number`json:""` // same as above
flClimbTilt boolean`json:""`
txClimb string`json:""`
idClimbPos number`json:""` // preferred position 1=NA, 2=Any, 3=Inner, 4=Middle, 5=Outer
flClimbLevel boolean`json:""` // can level generator
flClimbLevelSelf boolean`json:""`
flClimbLevelOther boolean`json:""`
flClimbMove boolean`json:""`
flClimbOther boolean`json:""`
numClimbOther number`json:""` // buddies #

constructor() {
this.flClimb = false`json:""`
this.idClimbType = ''`json:""`
this.numClimbHeight = 0`json:""`
this.flClimbSecure = false`json:""`
this.idClimbGrab = 1`json:""`
this.idClimbSpeed = 1`json:""`
this.flClimbTilt = false`json:""`
this.txClimb = ''`json:""`
this.idClimbPos = 1`json:""`
this.flClimbLevel = false`json:""`
this.flClimbLevelSelf = false`json:""`
this.flClimbLevelOther = false`json:""`
this.flClimbMove = false`json:""`
this.flClimbOther = false`json:""`
this.numClimbOther = 0`json:""`

}
}

export class ControlPanel {
flPanel boolean`json:""` // can manipulate control panel
flPanelBrake boolean`json:""`
flPanelRotation boolean`json:""`
flPanelPos boolean`json:""`
flPanelSensor boolean`json:""`
txPanelSensor string`json:""` // notes
constructor() {
this.flPanel = false`json:""`
this.flPanelBrake = false`json:""`
this.flPanelRotation = false`json:""`
this.flPanelPos = false`json:""`
this.flPanelSensor = false`json:""`
this.txPanelSensor = ''`json:""`
}
}

export class Auto {
flAuto boolean`json:""` // can auto
flAutoLine boolean`json:""`
flAutoShoot boolean`json:""` // can auto shoot
numAutoShoot number`json:""` // number of balls
numAutoLoad number`json:""` // pickup
constructor() {
this.flAuto = false`json:""`
this.flAutoLine = false`json:""`
this.flAutoShoot = false`json:""`
this.numAutoShoot = 0`json:""`
this.numAutoLoad = 0`json:""`
}
}

export class Record {
dtCreated string`json:""`
dtModified string`json:""`
txComputerName string`json:""`
constructor() {
this.dtCreated = ''`json:""`
this.dtModified = ''`json:""`
this.txComputerName = ''`json:""`
}
}


export class TeamMember {
team Team`json:""`
id number`json:""`
firstName string`json:""`
lastName string`json:""`
constructor() {
this.team = new Team()`json:""`
this.id = 0`json:""`
this.firstName = ''`json:""`
this.lastName = ''`json:""`
}

}

