package entity

type RefOpt func(ref *Referal) *Referal

func WithRefUsername(username string) RefOpt {
	return func(ref *Referal) *Referal {
		ref.username = username
		return ref
	}
}

func WithRefReferalLink(referalLink string) RefOpt {
	return func(ref *Referal) *Referal {
		ref.referalLink = referalLink
		return ref
	}
}

func WithRefId(id int64) RefOpt {
	return func(ref *Referal) *Referal {
		ref.id = id
		return ref
	}
}

func WithRefInServiceId(inServiceId int64) RefOpt {
	return func(ref *Referal) *Referal {
		ref.inServiceId = inServiceId
		return ref
	}
}

func WithRefAllUsersCount(allUsers int64) RefOpt {
	return func(ref *Referal) *Referal {
		ref.allUserCount = allUsers
		return ref
	}
}

func WithRefUsersLastNDays(users int64) RefOpt {
	return func(ref *Referal) *Referal {
		ref.usersLastNDays = users
		return ref
	}
}

type UserOpt func(usr *User) *User

func WithUsrReferalId(referalId int64) UserOpt {
	return func(usr *User) *User {
		usr.referalId = referalId
		return usr
	}
}

func WithUsrUsername(username string) UserOpt {
	return func(usr *User) *User {
		usr.username = username
		return usr
	}
}

func WithUsrInServiceId(inServiceId int64) UserOpt {
	return func(usr *User) *User {
		usr.inServiceId = inServiceId
		return usr
	}
}

func WithUsrReferalLink(referalLink string) UserOpt {
	return func(usr *User) *User {
		usr.referalLink = referalLink
		return usr
	}
}

type AdmOpts func(adm *Admin) *Admin

func WithAdmId(id int64) AdmOpts {
	return func(adm *Admin) *Admin {
		adm.id = id
		return adm
	}
}

func WithPassword(password string) AdmOpts {
	return func(adm *Admin) *Admin {
		adm.password = password
		return adm
	}
}
