package app

import (
	"lab/internal/adapters/postgres"
)

type Laber interface{}

type Lab struct {
	PG *postgres.LabPostgres
}

func NewLab(p *postgres.LabPostgres) *Lab {
	return &Lab{
		PG: p,
	}
}
