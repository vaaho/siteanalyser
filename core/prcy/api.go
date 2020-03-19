package prcy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	analyseBaseUrlPattern  = "https://a.pr-cy.ru/api/v1.1.0/analysis/base/%s?key=%s"
	analyseTestsUrlPattern = "https://a.pr-cy.ru/api/v1.1.0/analysis/base/%s?key=%s&excludeHistory=1&tests=%s"
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

func (a Api) analyseTestsUrl(domain string, tests []string) string {
	return fmt.Sprintf(analyseTestsUrlPattern, domain, a.Token, strings.Join(tests, ","))
}

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

	err = json.Unmarshal(data, model)
	if err != nil {
		return nil, err
	}

	model.HttpStatus = resp.StatusCode

	return model, nil
}
