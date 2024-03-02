package usercases

type ReferalService interface {
	CreateReferal()
	GetReferalById()
	GetReferalByLink()
	GenerateReferalLink()
}

type UserService interface {
}

type ReferalUseCases struct {
	userService    UserService
	referalService ReferalService
}

func (u ReferalUseCases) CreateReferal() {
	u.referalService.CreateReferal()
}

func (u ReferalUseCases) GetReferalById() {
	u.referalService.GetReferalById()
}

func (u ReferalUseCases) GetReferalByLink() {
	u.referalService.GetReferalByLink()
}

func (u ReferalUseCases) GenerateReferalLink() {
	u.referalService.GenerateReferalLink()
}
