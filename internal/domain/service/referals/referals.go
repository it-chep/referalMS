package referals

type ReferalsStorage interface {
}

type ReferalService struct {
	storage ReferalsStorage
	//WriteRepo
}
