package entity

type User struct {
	name        string
	tgID        int64
	adminId     int64
	referalId   int64
	username    string
	inServiceId int64
	referalLink string
}

func NewUser(tgId, adminId int64, name string, opts ...UserOpt) *User {
	u := &User{
		tgID:    tgId,
		adminId: adminId,
		name:    name,
	}

	for _, opt := range opts {
		opt(u)
	}

	return u
}

func (usr *User) GetName() string {
	return usr.name
}

func (usr *User) GetTgId() int64 {
	return usr.tgID
}

func (usr *User) GetReferalLink() string {
	return usr.referalLink
}

func (usr *User) GetInServiceId() int64 {
	return usr.inServiceId
}

func (usr *User) GetUsername() string {
	return usr.username
}

func (usr *User) GetReferalId() int64 {
	return usr.referalId
}
