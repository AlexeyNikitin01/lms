package httpgin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"lab/internal/app"
	"lab/internal/repository/pg/entity"
)

type LabResponse struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Lecture string `json:"lecture"`
	Formule string `json:"formule"`
}

func getLab(_ *app.Lab) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid lab ID"})
			return
		}

		lab, err := entity.FindLab(c, boil.GetContextDB(), id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "lab not found"})
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

type AirplaneModel struct {
	ID           int                `json:"id"`
	Name         string             `json:"name"`
	Manufacturer string             `json:"manufacturer"`
	Year         int                `json:"year"`
	Description  string             `json:"description"`
	Materials    []AirplaneMaterial `json:"materials"`
}

type AirplaneMaterial struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

func getAirplaneModel(_ *app.Lab) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		airplane, err := entity.AirplaneModels(
			qm.Where("id = ?", id),
			qm.Load(entity.AirplaneModelRels.AirplaneMaterials), // Подгружаем связанные материалы
		).One(c, boil.GetContextDB())

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Model not found"})
			return
		}

		response := gin.H{
			"id":           airplane.ID,
			"name":         airplane.Name,
			"manufacturer": airplane.Manufacturer,
			"year":         airplane.Year,
			"description":  airplane.Description,
			"materials":    make([]gin.H, len(airplane.R.AirplaneMaterials)),
		}

		for i, material := range airplane.R.AirplaneMaterials {
			response["materials"].([]gin.H)[i] = gin.H{
				"id":          material.ID,
				"name":        material.Name,
				"description": material.Description,
				"color":       material.Color,
			}
		}

		c.JSON(http.StatusOK, response)
	}
}
