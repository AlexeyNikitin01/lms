package hh

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Vacancy представляет структуру вакансии
type Vacancy struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Salary struct {
		From     *int   `json:"from"` // Используем указатель, так как поле может отсутствовать
		To       *int   `json:"to"`
		Currency string `json:"currency"`
	} `json:"salary"`
	Employer struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"employer"`
	PublishedAt string `json:"published_at"`
}

// VacanciesResponse представляет ответ на запрос поиска вакансий
type VacanciesResponse struct {
	Items []Vacancy `json:"items"`
	Page  int       `json:"page"`
	Pages int       `json:"pages"`
	Found int       `json:"found"`
}

type Companies struct {
	Items []Company `json:"items"`
	Page  int       `json:"page"`
	Pages int       `json:"pages"`
	Found int       `json:"found"`
}

type Company struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// SearchVacanciesByEmployer ищет вакансии по названию предприятия
func (c *HHClient) SearchVacanciesByEmployer(employerName string) (*VacanciesResponse, error) {
	path := fmt.Sprintf("%s/employers?text=%s",
		baseURL, url.QueryEscape(employerName))
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var resultCompanies Companies
	if err = json.NewDecoder(resp.Body).Decode(&resultCompanies); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	path = fmt.Sprintf("%s/vacancies?employer_id=%s",
		baseURL, url.QueryEscape(resultCompanies.Items[0].ID))
	reqVacs, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	reqVacs.Header.Set("User-Agent", userAgent)
	reqVacs.Header.Set("Accept", "application/json")

	respVacs, err := c.httpClient.Do(reqVacs)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if respVacs.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(respVacs.Body)
		return nil, fmt.Errorf("API request failed with status %d: %s", respVacs.StatusCode, string(body))
	}

	var result VacanciesResponse
	if err = json.NewDecoder(respVacs.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &result, nil
}
