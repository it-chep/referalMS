package entity

import "time"

type Admin struct {
	id                int64
	login             string
	password          string
	integrationsToken string
	integratorId      int64
	lastLogin         time.Time
	registrationTime  time.Time
}

func NewAdmin(login, password, integrationsToken string, opts ...AdmOpts) *Admin {
	return &Admin{
		login:             login,
		password:          password,
		integrationsToken: integrationsToken,
		lastLogin:         time.Now(),
	}
}

func (adm *Admin) GetId() int64 {
	return adm.id
}

func (adm *Admin) GetLogin() string {
	return adm.login
}

func (adm *Admin) GetPassword() string {
	return adm.password
}

func (adm *Admin) GetIntegrationsToken() string {
	return adm.integrationsToken
}

func (adm *Admin) GetIntegratorId() int64 {
	return adm.integratorId
}

type WinnersFilter struct {
	limit        int
	daysInterval int
}

func NewWinnersFilter(limit, daysInterval int) *WinnersFilter {
	return &WinnersFilter{
		limit:        limit,
		daysInterval: daysInterval,
	}
}

func (wf *WinnersFilter) GetLimit() int {
	return wf.limit
}

func (wf *WinnersFilter) GetDaysInterval() int {
	return wf.daysInterval
}
