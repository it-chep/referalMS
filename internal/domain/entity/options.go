package entity

type Opt func(ref *Referal) *Referal

func WithName(name string) Opt {
	return func(ref *Referal) *Referal {
		ref.name = name
		return ref
	}
}
