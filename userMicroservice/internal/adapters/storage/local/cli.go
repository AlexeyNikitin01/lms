package local

import (
	"lms-user/internal/adapters/storage"
)

const PATH = "tmp/"

type Local struct {
	Path string
}

func NewLocal() storage.IFace {
	return &Local{
		Path: PATH,
	}
}
