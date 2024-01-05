package user

import "time"

type User struct {
	Admin			bool
	ID              int
	Login           string
	Password        string
	PersonalData    PersonalStruct
	ProgressCourse  ProgressCourseStruct
	CompletedCourse CompletedCourseStruct
	FavoritInfo     FavoritInfoStruct
}

type PersonalStruct struct {
	Name      string
	Surname   string
	Email     string
	Phone     string
	PlaceWork string
	Position  string
}

type ProgressCourseStruct struct {
	Progress  int
	Name      string
	StartDate time.Time
}

type CompletedCourseStruct struct {
	Certificat bool
	Name       string
	SpendTime  time.Time
}

type FavoritInfoStruct struct {
	Text  string
	Topic string
	Name  string
}
