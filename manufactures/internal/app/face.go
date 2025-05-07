package app

import (
	"manufactures/internal/adapters/dadata"
)

type Apper interface {
	get() error
}

type AppManfs struct {
	D dadata.Dadata
}

func NewAppManfs(d dadata.Dadata) AppManfs {
	return AppManfs{
		D: d,
	}
}
