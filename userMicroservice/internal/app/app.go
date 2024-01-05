package app

type AppUser interface {

}

type appUser struct {

}

func CreateAppUser() AppUser {
	return &appUser{} 
}