package nscAbstraction

import "bitbucket.org/turbosoftnetworks/netscope-api-v2/nsc"

type NscFunc interface {
	GetActivity (q nsc.Query) (d nsc.Data, err error)
	GetBypassState () (d nsc.DataBypass, err error)
	GetLicense () (d nsc.DataLicense, err error)
	GetRtt (q nsc.Query) (d nsc.DataRtt, err error)
	GetTopN (q nsc.Query) (d nsc.DataTopN, err error)
	GetTopClasses (q nsc.Query) (d nsc.DataTopClasses, err error)

	SetBypassState (q nsc.Query) (d nsc.DataBypass, err error)
}

type RegisteredNscFuncAbs struct {
	NscFunctions NscFunc
}

func NewRegisteredNscFuncAbs (nsc NscFunc) *RegisteredNscFuncAbs {
	return &RegisteredNscFuncAbs{nsc}
}

func (n RegisteredNscFuncAbs) GetActivity (q nsc.Query) (d nsc.Data, err error) {
	d, err = n.NscFunctions.GetActivity(q)
	return
}

func (n RegisteredNscFuncAbs) GetBypassState () (d nsc.DataBypass, err error) {
	d, err = n.NscFunctions.GetBypassState()
	return
}

func (n RegisteredNscFuncAbs) GetTopClasses (q nsc.Query) (d nsc.DataTopClasses, err error) {
	d, err = n.NscFunctions.GetTopClasses(q)
	return
}

func (n RegisteredNscFuncAbs) GetLicense () (d nsc.DataLicense, err error) {
	d, err = n.NscFunctions.GetLicense()
	return
}

func (n RegisteredNscFuncAbs) GetRtt (q nsc.Query) (d nsc.DataRtt, err error) {
	d, err = n.NscFunctions.GetRtt(q)
	return
}

func (n RegisteredNscFuncAbs) GetTopN (q nsc.Query) (d nsc.DataTopN, err error) {
	d, err = n.NscFunctions.GetTopN(q)
	return
}

func (n RegisteredNscFuncAbs) SetBypassState (q nsc.Query) (d nsc.DataBypass, err error) {
	d, err = n.NscFunctions.SetBypassState(q)
	return
}
