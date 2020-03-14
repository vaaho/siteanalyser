package prcy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const defaultDelay = 1000
const analyseUrl = "https://a.pr-cy.ru/api/v1.1.0/analysis/base"

type Api struct {
	Token     string
	Delay     int
}

func NewApi(token string, delay int) Api {
	if delay == 0 {
		delay = defaultDelay
	}
	return Api{Token:token, Delay:delay}
}

func (a Api) analyseUrl(domain string) string {
	return analyseUrl + "/" + domain + "?key" + a.Token
}

func (a Api) requestAnalysis(domain string) (string, *ApiError) {
	body, err := a.requestAnalysisData(domain)
	if err != nil {
		return "", err
	}

	return "", nil
}

func (a Api) requestAnalysisData(domain string) (string, *ApiError) {
	url := a.analyseUrl(domain)

	resp, err := http.Get(url)
	if err != nil {
		return "", NewApiError(Network, err.Error())
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", NewApiError(Parse, err.Error())
	}

	return string(data), nil
}

func parseErrorResult(statusCode int, data byte[]) *ApiError {
	var errModel ErrorModel

	err = json.Unmarshal(data, &errModel)
	if ()
}

