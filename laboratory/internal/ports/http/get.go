package httpgin

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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

func getAI(a *app.Lab) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем ID лекции из параметров URL
		id := c.Param("id")

		// Проверяем, что ID является числом
		lectureID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lecture ID"})
			return
		}

		// Загружаем лекцию со всеми связанными данными
		lecture, err := entity.LecturesAis(
			qm.Where("id = ?", lectureID),
			qm.Load(
				qm.Rels(
					entity.LecturesAiRels.LectureAnalysisResults,
					entity.AnalysisResultRels.AnalysisDefectFindings,
					entity.DefectFindingRels.DefectType,
				),
			),
			qm.Load(
				qm.Rels(
					entity.LecturesAiRels.LectureAnalysisResults,
					entity.AnalysisResultRels.AnalysisRecommendations,
				),
			),
		).One(c, boil.GetContextDB())

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Lecture not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}

		// Формируем ответ
		response := gin.H{
			"id":         lecture.ID,
			"title":      lecture.Title,
			"author":     lecture.Author,
			"content":    lecture.Content,
			"created_at": lecture.CreatedAt,
			"analyses":   make([]gin.H, 0, len(lecture.R.LectureAnalysisResults)),
		}

		// Обрабатываем анализы
		for _, analysis := range lecture.R.LectureAnalysisResults {
			analysisData := gin.H{
				"id":               analysis.ID,
				"image_path":       analysis.ImagePath,
				"analysis_date":    analysis.AnalysisDate,
				"confidence_score": analysis.ConfidenceScore,
				"defects":          make([]gin.H, 0, len(analysis.R.AnalysisDefectFindings)),
				"recommendations":  make([]gin.H, 0, len(analysis.R.AnalysisRecommendations)),
			}

			// Обрабатываем дефекты
			for _, df := range analysis.R.AnalysisDefectFindings {
				analysisData["defects"] = append(analysisData["defects"].([]gin.H), gin.H{
					"type":        df.R.DefectType.Name,
					"probability": df.Probability,
					"description": df.R.DefectType.Description,
					"severity":    df.Severity,
				})
			}

			// Обрабатываем рекомендации
			for _, rec := range analysis.R.AnalysisRecommendations {
				analysisData["recommendations"] = append(analysisData["recommendations"].([]gin.H), gin.H{
					"text":     rec.RecommendationText,
					"priority": rec.Priority,
				})
			}

			response["analyses"] = append(response["analyses"].([]gin.H), analysisData)
		}

		c.JSON(http.StatusOK, response)
	}
}
