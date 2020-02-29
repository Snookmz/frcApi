package handler

func convertScoutsToScoutsCsv (scouts []Scout) (scoutsCsv []ScoutCsv) {

	for _, v := range scouts {
		var s ScoutCsv
		s.ParentData.TxEvent = v.ParentData.TxEvent
		s.ParentData.DeviceName = v.ParentData.DeviceName

		s.ParentData.TeamDetails.NumMatch = v.ParentData.TeamDetails.NumMatch
		s.ParentData.TeamDetails.IdAlliance = v.ParentData.TeamDetails.IdAlliance
		s.ParentData.TeamDetails.IdDriveStation = v.ParentData.TeamDetails.IdDriveStation
		s.ParentData.TeamDetails.IdTeam = v.ParentData.TeamDetails.IdTeam
		s.ParentData.TeamDetails.TxScoutName = v.ParentData.TeamDetails.TxScoutName

	}


	return
}
