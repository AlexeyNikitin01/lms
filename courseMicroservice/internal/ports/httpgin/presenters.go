package httpgin

type CourseRequest struct {
	Name string `json:"name"`
}

type LectureRequest struct {
	Title    string `json:"title"`
	Lecture  string `json:"lecture"`
	CourseID int    `json:"course_id"`
}
