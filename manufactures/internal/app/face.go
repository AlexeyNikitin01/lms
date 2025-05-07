package app

type apper interface {
	get() error
}

type AppManfs struct{}

func NewAppManfs() AppManfs {
	return AppManfs{}
}
