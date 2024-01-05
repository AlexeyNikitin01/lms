package user

type Course struct {
	ID       int
	Name     string
	Progress int
}

type Info struct {
	ID   int
	Name string
	Text string
}

type User struct {
	ID          int
	Name        string
	Email       string
	Login       string
	Password    string
	Course      []Course
	InfoFavorit []Info
}
