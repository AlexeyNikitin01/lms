package httpgin

import (
	"course/internal/repository/pg/entity"
)

type CourseRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image,omitempty"`
}

type LectureRequest struct {
	Title    string `json:"title"`
	Lecture  string `json:"lecture"`
	CourseID int64  `json:"course_id"`
}

type FindLecturesRequest struct {
	CourseID int64 `json:"course_id"`
}

type AllCoursesRequest struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type Answer struct {
	AnswerID  int64  `json:"answer_id"`
	Text      string `json:"text"`
	IsCorrect bool   `json:"is_correct"`
}

type Question struct {
	QuestionID int64    `json:"question_id"`
	Text       string   `json:"text"`
	Answers    []Answer `json:"answers"`
}

type Test struct {
	TestID    int64      `json:"test_id"`
	Name      string     `json:"name"`
	Questions []Question `json:"questions"`
}

type Lecture struct {
	LectureID int64  `json:"lecture_id"`
	Name      string `json:"name"`
	Text      string `json:"text"`
	Tests     []Test `json:"tests"`
}

type Module struct {
	ModuleID int64     `json:"module_id"`
	Name     string    `json:"name"`
	Lectures []Lecture `json:"lectures"`
}

type Course struct {
	CourseID int64    `json:"course_id"`
	Modules  []Module `json:"modules"`
}

func convertToEntityCourse(course Course, courseID int64) *entity.Course {
	var entityModules entity.ModuleSlice

	for _, m := range course.Modules {
		var entityLectures entity.LectureSlice

		for _, l := range m.Lectures {
			var entityTests entity.TestSlice

			for _, t := range l.Tests {
				var entityQuestions entity.QuestionSlice

				for _, q := range t.Questions {
					entityQuestions = append(entityQuestions, &entity.Question{
						ID:   q.QuestionID,
						Text: q.Text,
					})

					var answers entity.AnswerSlice

					for _, a := range q.Answers {
						answers = append(answers, &entity.Answer{
							ID:        a.AnswerID,
							Text:      a.Text,
							IsCorrect: a.IsCorrect,
						})
					}
				}

				newTest := &entity.Test{
					ID:   t.TestID,
					Name: t.Name,
				}

				newTest.R = newTest.R.NewStruct()
				newTest.R.Questions = entityQuestions

				entityTests = append(entityTests, newTest)
			}

			newLecture := &entity.Lecture{
				ID:      l.LectureID,
				Title:   l.Name,
				Lecture: l.Text,
			}
			newLecture.R = newLecture.R.NewStruct()
			newLecture.R.Tests = entityTests

			entityLectures = append(entityLectures, newLecture)
		}

		newModule := &entity.Module{
			ID:       m.ModuleID,
			Name:     m.Name,
			CourseID: courseID,
		}
		newModule.R = newModule.R.NewStruct()
		newModule.R.Lectures = entityLectures

		entityModules = append(entityModules, newModule)
	}

	newCourse := entity.Course{}
	newCourse.R = newCourse.R.NewStruct()
	newCourse.R.Modules = entityModules

	return &newCourse
}
