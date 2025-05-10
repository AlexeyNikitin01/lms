package app

import (
	"manufactures/internal/adapters/dadata"
	"manufactures/internal/adapters/hh"
)

type Apper interface {
	data() error
}

type AppManfs struct {
	D  dadata.Dadata
	HH hh.HHClient
}

func NewAppManfs(d dadata.Dadata, hh hh.HHClient) AppManfs {
	return AppManfs{
		D:  d,
		HH: hh,
	}
}
