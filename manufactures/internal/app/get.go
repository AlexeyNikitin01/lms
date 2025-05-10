package app

import (
	"manufactures/internal/adapters/hh"
)

func (a AppManfs) data() error {
	return nil
}

func (a AppManfs) dataVacancies() (hh.VacanciesResponse, error) {
	return hh.VacanciesResponse{}, nil
}
