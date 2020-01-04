package handler

import "bitbucket.org/turbosoftnetworks/netscope-api-v2/nsc"

type Nsc struct {}

func (n Nsc) GetActivity (q nsc.Query) (d nsc.Data, err error) {
	return
}

func (n Nsc) GetBypassState () (d nsc.DataBypass, err error) {
	return
}

func (n Nsc) GetLicense () (d nsc.DataLicense, err error) {
	return
}

func (n Nsc) GetRtt (q nsc.Query) (d nsc.DataRtt, err error) {
	return
}

func (n Nsc) GetTopN (q nsc.Query) (d nsc.DataTopN, err error) {
	return
}

func (n Nsc) GetTopClasses (q nsc.Query) (d nsc.DataTopClasses, err error) {
	return
}

func (n Nsc) SetBypassState (q nsc.Query) (d nsc.DataBypass, err error) {
	return
}
