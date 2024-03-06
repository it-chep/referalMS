package entity

type Referal struct {
	id             int64
	name           string
	tgID           int64
	adminId        int64
	username       string
	inServiceId    int64
	allUserCount   int64
	usersLastNDays int64
	referalLink    string
}

func NewReferal(tgID, adminId int64, name string, opts ...RefOpt) *Referal {
	ref := &Referal{
		tgID:    tgID,
		name:    name,
		adminId: adminId,
	}

	for _, opt := range opts {
		opt(ref)
	}
	return ref
}

func (ref *Referal) GetId() int64 {
	return ref.id
}

func (ref *Referal) GetName() string {
	return ref.name
}

func (ref *Referal) GetTgId() int64 {
	return ref.tgID
}

func (ref *Referal) GetReferalLink() string {
	return ref.referalLink
}

func (ref *Referal) GetInServiceId() int64 {
	return ref.inServiceId
}

func (ref *Referal) GetUsername() string {
	return ref.username
}

func (ref *Referal) GetAllUsers() int64 {
	return ref.allUserCount
}

func (ref *Referal) GetLastNDaysUsers() int64 {
	return ref.usersLastNDays
}

func (ref *Referal) SetId(id int64) {
	ref.id = id
}

func (ref *Referal) SetAllUsers(users int64) {
	ref.allUserCount = users
}

func (ref *Referal) SetLastNDays(users int64) {
	ref.usersLastNDays = users
}
