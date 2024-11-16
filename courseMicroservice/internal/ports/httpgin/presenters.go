package httpgin

type CourseRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

type LectureRequest struct {
	Title    string `json:"title"`
	Lecture  string `json:"lecture"`
	CourseID int    `json:"course_id"`
}

type FindLecturesRequest struct {
	CourseID int `json:"course_id"`
}

type AllCoursesRequest struct {
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
}
