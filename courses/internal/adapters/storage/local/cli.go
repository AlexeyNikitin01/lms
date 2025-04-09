package local

import (
	"course/internal/adapters/storage"
)

const PATH = "tmp/"

type Local struct {
	Path string
}

func NewLocal() storage.ICloud {
	return &Local{
		Path: PATH,
	}
}
