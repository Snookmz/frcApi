package nsc

import (
	"encoding/json"
	"errors"
	"fmt"
)

func ReturnJsonFormat (format string) (j string, err error) {

	var q Query
	var b []byte
	if format == "activity" {
		q.Component = "stats"
		q.Item = "bandwidth_flow"
		q.Action = "query"
		q.Parameters.ClassId = 8
		q.Parameters.Direction = "inbound"
		q.Parameters.Duration = 300
		q.Parameters.Period = "current"
	} else if format == "topN" {
		q.Component = "stats"
		q.Item = "top_flows"
		q.Action = "query"
		q.Parameters.ClassId = 8
		q.Parameters.Direction = "inbound|outbound"
		q.Parameters.Duration = 300
		q.Parameters.Period = "current"
		q.Parameters.Target = "ip_src"
	} else if format == "rtt" {
		q.Component = "stats"
		q.Item = "rtt"
		q.Action = "query"
		q.Parameters.ClassId = 8
	} else if format == "bypass" {
		q.Action = "command"
		q.Component = "bypass"
		q.Command = "bypass|online"
	} else if format == "topClasses" {
		q.Component = "stats"
		q.Item = "top_classes"
		q.Action = "query"
		q.Parameters.ClassId = 8
		q.Parameters.Direction = "inbound|outbound"
		q.Parameters.Duration = 300
		q.Parameters.Period = "current"
	}

	b, err = json.Marshal(q)
	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from json.Marshal: %s", err.Error()))
	}

	j = string(b)


	return
}
