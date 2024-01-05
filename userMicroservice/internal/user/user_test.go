package user

import (
	"testing"
	"time"
)

func TestPersonalData(t *testing.T) {
	u := User{}
	u.PersonalData.Email = "email@com.ru"
	u.PersonalData.Name = "Name"
	u.PersonalData.Surname = "Surname"
	u.PersonalData.Phone = "+79993432244"
	u.PersonalData.PlaceWork = "MAI"
	u.PersonalData.Position = "student"

	if u.PersonalData.Email != "email@com.ru" {
		t.Error("Not correct email")
	}

	if u.PersonalData.Name != "Name" {
		t.Error("Not correct name")
	}

	if u.PersonalData.Surname != "Surname" {
		t.Error("Not correct surname")
	}

	if u.PersonalData.Phone != "+79993432244" {
		t.Error("Not correct phone")
	}

	if u.PersonalData.PlaceWork != "MAI" {
		t.Error("Not correct place work")
	}

	if u.PersonalData.Position != "student" {
		t.Error("Not correct position")
	}
}

func TestUserID(t *testing.T) {
	u := User{}
	
	if u.ID != 0 {
		t.Error("Not correct id")
	}
}

func TestUserAdmin(t *testing.T) {
	u := User{}
	
	if u.Admin != false {
		t.Error("Admin not is true")
	}
}

func TestProgres(t *testing.T) {
	u := User{}
	s := time.Now()
	u.ProgressCourse.Name = "fiz-him"
	u.ProgressCourse.Progress = 50
	u.ProgressCourse.StartDate = s

	if u.ProgressCourse.StartDate != s {
		t.Error("not correct time")
	}

	if u.ProgressCourse.Progress == 50 && u.ProgressCourse.Progress > 100 && u.ProgressCourse.Progress < 0 {
		t.Error("not correct progress")
	}

	if u.ProgressCourse.Name != "fiz-him" {
		t.Error("not correct course name")
	}
}

func TestCompleted(t *testing.T) {
	u := User{}
	s := time.Now()
	u.CompletedCourse.Certificat = true
	u.CompletedCourse.Name = "fiz-him"
	u.CompletedCourse.SpendTime = s

	if u.CompletedCourse.Name != "fiz-him" {
		t.Error("not correct course name")
	}

	if u.CompletedCourse.SpendTime != s {
		t.Error("not correct spend course time")
	}

	if u.CompletedCourse.Certificat != true {
		t.Error("not correct certificate bool")
	}
}

func TestFavorite(t *testing.T) {
	u := User{}
	u.FavoritInfo.Topic = "polimer"
	u.FavoritInfo.Name = "fiz-him"
	u.FavoritInfo.Text = "text" 

	if u.FavoritInfo.Topic != "polimer" {
		t.Error("not correct topic")
	}

	if u.FavoritInfo.Text != "text" {
		t.Error("not correct text favorite")
	}

	if u.FavoritInfo.Name != "fiz-him" {
		t.Error("not correct name favorite")
	}
}
