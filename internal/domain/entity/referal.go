package entity

type Referal struct {
	name string
	tgID int64

	userCount int64
}

func NewReferal(tgID int64, opts ...Opt) *Referal {
	ref := &Referal{
		tgID: tgID,
	}

	for _, opt := range opts {
		opt(ref)
	}
	return ref
}

func (ref *Referal) GetName() string {
	return ref.name
}

func (ref *Referal) GetTgId() int64 {
	return ref.tgID
}

func (ref *Referal) IncUsersCount() {
	ref.userCount++
}

func test() {
	ref := NewReferal(1, WithName("maxim"))
	ref.IncUsersCount()
}
