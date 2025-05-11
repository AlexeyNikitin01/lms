package httpgin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"lab/internal/app"
	"lab/internal/repository/pg/entity"
)

type UpdateLabRequest struct {
	Content string `json:"content,omitempty"`
	Formule string `json:"formule,omitempty"`
}

func updateLab(_ *app.Lab) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid lab ID"})
			return
		}

		var req UpdateLabRequest
		if err = c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}

		lab, err := entity.FindLab(c, boil.GetContextDB(), id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "lab not found"})
			return
		}

		// Обновляем поля
		if req.Content != "" {
			lab.Lecture = req.Content
		}
		if req.Formule != "" {
			lab.Formule = req.Formule
		}

		if _, err = lab.Update(c, boil.GetContextDB(), boil.Infer()); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update lab"})
			return
		}

		response := LabResponse{
			ID:      lab.ID,
			Title:   lab.Title,
			Author:  lab.Author,
			Lecture: lab.Lecture,
			Formule: lab.Formule,
		}

		c.JSON(http.StatusOK, response)
	}
}
