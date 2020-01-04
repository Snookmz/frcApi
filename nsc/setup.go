package nsc

import (
	"github.com/fatih/pool"
)

type Nsc struct {
	Pool pool.Pool
}

func NewNsc (pool pool.Pool) *Nsc {
	return &Nsc{Pool: pool}
}

type Query struct {
	Component string `json:"component"`
	Action string `json:"action"`
	Item string `json:"item"`
	Command string `json:"command"`
	Parameters Parameters `json:"parameters"`
}

type Parameters struct {
	ClassId int `json:"class_id"`
	Direction string `json:"direction"`
	Period string `json:"period"`
	Duration int `json:"duration"`
	Target string `json:"target"`
	IpSrc string `json:"ip_src"`
	IpDst string `json:"ip_dst"`
	PortDst int `json:"pt_dst"`
	PortSrc int `json:"pt_src"`
	IpProtocol string `json:"ip_prt"`
	Dscp string `json:"dscp"`
	HttpUrl string `json:"http_url"`
	HttpDomain string `json:"http_domain"`
	DpiProtocol string `json:"dpi_protocol"`
	DomainUser string `json:"domain_user"`
	VlanTag string `json:"vlan_tag"`
}

type Data struct {
	Result Result `json:"result"`
	Direction string `json:"direction"`
	Status string `json:"status"`
}

type DataBypass struct {
	Result ResultBypass `json:"result"`
	Status string `json:"status"`
}

type DataRtt struct {
	Result ResultRtt `json:"result"`
	Direction string `json:"direction"`
	Status string `json:"status"`
}


type DataTopN struct {
	Result ResultTopN `json:"result"`
	Direction string `json:"direction"`
	Status string `json:"status"`
}

type DataTopClasses struct {
	Result ResultTopClasses `json:"result"`
	Direction string `json:"direction"`
	Status string `json:"status"`
}

type DataLicense struct {
	Result ResultLicense `json:"result"`
	Status string `json:"status"`
}

type Result struct {
	Data []string `json:"data"`
	Time []string `json:"time"`
	Error string `json:"error"`
}

type ResultBypass struct {
	State string `json:"state"`
	Error string `json:"error"`
}

type ResultTopN struct {
	Targets []string `json:"targets"`
	Volumes []string `json:"volumes"`
	Sum string `json:"sum"`
	Error string `json:"error"`
}

type ResultTopClasses struct {
	ClassIds []string `json:"class_ids"`
	ClassNames []string `json:"class_names"`
	Volumes []string `json:"volumes"`
	Sum string `json:"sum"`
}

type ResultRtt struct {
	Rtt []string `json:"rtt"`
	Time []string `json:"time"`
	Error string `json:"error"`
}

type ResultLicense struct {
	Validity string `json:"validity"`
	ExpirationDate string `json:"expiration-date"`
	HardwareId string `json:"hardware-id"`
	OwnerName string `json:"owner-name"`
	Classes string `json:"classes"`
	Bandwidth string `json:"bandwidth"`
	Alerts string `json:"alerts"`
	Reports string `json:"reports"`
	ShapingEnabled string `json:"shaping-enabled"`
	Error string `json:"error"`
}