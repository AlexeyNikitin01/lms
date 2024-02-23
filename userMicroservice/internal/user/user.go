package user

import "time"

type User struct {
	Admin			bool
	ID              int
	Login           string
	Password        string
	PersonalData    Personal
	ProgressCourse  ProgressCourse
	CompletedCourse CompletedCourse
	FavoritInfo     FavoritInfo
}

type Personal struct {
	Name      string
	Surname   string
	Email     string
	Phone     string
	PlaceWork string
	Position  string
}

type ProgressCourse struct {
	Progress  int
	Name      string
	StartDate time.Time
}

type CompletedCourse struct {
	Certificat bool
	Name       string
	SpendTime  time.Time
}

type FavoritInfo struct {
	Text  string
	Topic string
	Name  string
}
