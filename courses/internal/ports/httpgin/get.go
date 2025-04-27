package httpgin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"course/internal/app"
)

func getAllCourses(app app.ICourseApp) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AllCoursesRequest
		err := c.BindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		courses, total, err := app.AllCourse(c, req.Limit, req.Offset)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		for _, course := range courses {
			course.PhotoURL = "https://storage.yandexcloud.net/lms-user/" + course.PhotoURL
		}

		c.JSON(http.StatusOK, gin.H{"courses": courses, "total": total})
	}
}

func getCourse(app app.ICourseApp) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		courseID, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
			return
		}

		course, err := app.GetCourse(c, courseID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		modules := []gin.H{}
		for _, m := range course.R.GetModules() {
			lectures := []gin.H{}
			for _, l := range m.R.GetLectures() {
				tests := []gin.H{}
				for _, t := range l.R.GetTests() {
					questions := []gin.H{}
					for _, q := range t.R.GetQuestions() {
						answers := []gin.H{}
						for _, a := range q.R.GetAnswers() {
							answers = append(answers, gin.H{
								"answer": gin.H{
									"answer_id":  a.ID,
									"text":       a.Text,
									"is_correct": a.IsCorrect,
								},
							})
						}
						questions = append(questions, gin.H{
							"question": gin.H{
								"question_id": q.ID,
								"text":        q.Text,
								"answers":     answers,
							},
						})
					}
					tests = append(tests,
						gin.H{
							"test": gin.H{
								"test_id":   t.ID,
								"name":      t.Name,
								"questions": questions,
							},
						})

				}
				lectures = append(lectures, gin.H{
					"lecture": gin.H{
						"lecture_id": l.ID,
						"text":       l.Lecture,
						"name":       l.Title,
						"tests":      tests,
					},
				})
			}
			modules = append(modules,
				gin.H{
					"module": gin.H{
						"module_id":   m.ID,
						"name_module": m.Name,
						"lectures":    lectures,
					},
				})
		}

		c.JSON(
			http.StatusOK,
			gin.H{
				"id":      courseID,
				"course":  course,
				"modules": modules,
			},
		)
	}
}
