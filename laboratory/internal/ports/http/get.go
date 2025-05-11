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
	ID                 int                `json:"id"`
	Name               string             `json:"name"`
	Manufacturer       string             `json:"manufacturer"`
	Year               int                `json:"year"`
	Description        string             `json:"description"`
	LectureDescription string             `json:"lecture_description"`
	Materials          []AirplaneMaterial `json:"materials"`
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

		// Получаем модель самолета с загруженными материалами
		airplane, err := entity.AirplaneModels(
			qm.Where("id = ?", id),
			qm.Load(entity.AirplaneModelRels.AirplaneMaterials),
		).One(c, boil.GetContextDB())
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Model not found"})
			return
		}

		response := AirplaneModel{
			ID:                 airplane.ID,
			Name:               airplane.Name,
			Manufacturer:       airplane.Manufacturer,
			Year:               airplane.Year.Int,
			Description:        airplane.Description.String,
			LectureDescription: airplane.LectureDescription.String,
		}

		// Обработка связанных материалов
		for _, mat := range airplane.R.AirplaneMaterials {
			response.Materials = append(response.Materials, AirplaneMaterial{
				ID:          mat.ID,
				Name:        mat.Name,
				Description: mat.Description.String,
				Color:       mat.Color.String,
			})
		}

		// Возвращаем результат
		c.JSON(http.StatusOK, response)
	}
}
