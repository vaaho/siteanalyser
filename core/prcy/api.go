package prcy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	neturl "net/url"
	"strings"
)

const (
	analyseBaseUrlPattern       = "https://a.pr-cy.ru/api/v1.1.0/analysis/base/%s?key=%s"
	analyseStatusBaseUrlPattern = "https://a.pr-cy.ru/api/v1.1.0/analysis/status/base/%s?key=%s"
	analyseUpdateBaseUrlPattern = "https://a.pr-cy.ru/api/v1.1.0/analysis/update/base/%s?key=%s"
	analyseTestsUrlPattern      = "https://a.pr-cy.ru/api/v1.1.0/analysis/base/%s?key=%s&excludeHistory=1&tests=%s"
)

type Api struct {
	Token string   // api key
	Tests []string // набор тестов для запроса
}

func NewApi(token string) Api {
	if token == "" {
		log.Fatal("Api key is required")
	}
	return Api{Token: token, Tests: TestNames}
}

func (a Api) analyseBaseUrl(domain string) string {
	return fmt.Sprintf(analyseBaseUrlPattern, domain, a.Token)
}

func (a Api) analyseStatusBaseUrl(domain string) string {
	return fmt.Sprintf(analyseStatusBaseUrlPattern, domain, a.Token)
}

func (a Api) analyseUpdateBaseUrl(domain string) string {
	return fmt.Sprintf(analyseUpdateBaseUrlPattern, domain, a.Token)
}

func (a Api) analyseTestsUrl(domain string, tests []string) string {
	return fmt.Sprintf(analyseTestsUrlPattern, domain, a.Token, strings.Join(tests, ","))
}

// Заправшивает базовый анализ
func (a Api) RequestAnalysis(domain string) (model *Analysis, err error) {
	url := a.analyseTestsUrl(domain, a.Tests)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	model = new(Analysis)
	if string(data) != "[]" { // в случае осутствия отчёта приходит пустой массив, мы его не парсим
		err = json.Unmarshal(data, model)
		if err != nil {
			return nil, err
		}
	}

	model.HttpStatus = resp.StatusCode

	return model, nil
}

// Заправшивает создание/обновление базового анализа
func (a Api) RequestAnalysisUpdate(domain string) (status int, err error) {
	url := a.analyseUpdateBaseUrl(domain)

	resp, err := http.PostForm(url, neturl.Values{})
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
