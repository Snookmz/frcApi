package handler

import (
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/database"
	"fmt"
)

type NsDatabase struct {}

func (d *NsDatabase) CreateClass (c database.Class) (newClass database.Class, err error) {
	return
}

func (d *NsDatabase) CreateConfigFilter (c database.ConfigFilter) (r database.ConfigFilter, err error) {
	return
}

func (d *NsDatabase) CreateConfigShaper (c database.ConfigShaper) (r database.ConfigShaper, err error) {
	return
}

func (d *NsDatabase) CreateGroup (g database.Group) (r database.Group, err error) {
	return
}

func (d *NsDatabase) CreateUser (name string, password string, groupId int) (err error) {
	return
}

func (d *NsDatabase) DeleteClass (classId int) (err error) {
	return
}

func (d *NsDatabase) DeleteConfigFilter (filterId int) (err error) {
	return
}

func (d *NsDatabase) DeleteConfigShaper (shaperId int) (err error) {
	return
}

func (d *NsDatabase) DeleteGroup (groupId int) (err error) {
	return
}

func (d *NsDatabase) EditClass (c database.Class) (err error) {
	return
}

func (d *NsDatabase) EditGroup (g database.Group) (err error) {
	return
}

func (d *NsDatabase) EditUser (userId int, u database.User) (err error) {
	return
}

func (d *NsDatabase) GetActivePolicy () (p database.Policy, err error) {
	p.PolicyId = 1
	return
}

func (d *NsDatabase) GetClasses (policyId int) (classes []database.Class, err error) {

	for i := 0; i < 10; i++ {
		var c database.Class
		c.PolicyId = 1
		c.ClassId = i
		c.FirstChildId = -i
		c.Name = fmt.Sprintf("class name #%v", i)
		classes = append(classes, c)
	}

	return
}

func (d *NsDatabase) GetConfigFilters (policyId int) (configFilters []database.ConfigFilter, err error) {
	return
}

func (d *NsDatabase) GetConfigShapers (policyId int) (configShapers []database.ConfigShaper, err error) {
	return
}

func (d *NsDatabase) GetDpiProtocols () (dpiProtocols []database.Dpi, err error) {
	return
}

func (d *NsDatabase) GetFilters (policyId int) (filters []database.Filter, err error) {

	for i := 0; i < 10; i++ {
		var f database.Filter
		f.PolicyId = 1
		f.DpiProtocol = fmt.Sprintf("dpi #%v", i)
		filters = append(filters, f)
	}

	return
}

func (d *NsDatabase) GetGroups () (groups []database.Group, err error) {
	return
}

func (d *NsDatabase) GetGroupFromGroupId (groupId int) (rG database.Group, err error) {
	return
}

func (d *NsDatabase) GetGroupIdFromGroupName (groupName string) (groupId int, err error) {
	return
}

func (d *NsDatabase) GetPolicyByPolicyId (policyId int) (p database.Policy, err error) {
	return
}

func (d *NsDatabase) GetPolicies () (policies []database.Policy, err error) {
	return
}

func (d *NsDatabase) GetShapers (policyId int) (shapers []database.Shaper, err error) {
	return
}

func (d *NsDatabase) GetUserStateFromSessionId (sessionId string) (u database.UserState, err error) {
	u.SessionId = "unit test session id"
	u.SignedIn = true
	return
}

func (d *NsDatabase) LoginUser (name string, password string) (sessionId string, err error) {
	sessionId = "unit test session id"
	return
}

func (d *NsDatabase) LogoutUser (name string) (err error) {
	return
}
