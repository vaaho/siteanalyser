package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtractDomains_1(t *testing.T) {
	d := ExtractDomains("https://www.somesite.ru/profile/7798752")
	assert.Equal(t, []string{"www.somesite.ru"}, d)
}

func TestExtractDomains_2(t *testing.T) {
	d := ExtractDomains("http://www.somesite.ru/")
	assert.Equal(t, []string{"www.somesite.ru"}, d)
}

func TestExtractDomains_3(t *testing.T) {
	d := ExtractDomains(" www.somesite.ru  ")
	assert.Equal(t, []string{"www.somesite.ru"}, d)
}

func TestExtractDomains_4(t *testing.T) {
	d := ExtractDomains("firstsite.org/, http://www.somesite.ru/, lastsite.com ")
	assert.Equal(t, []string{"firstsite.org", "www.somesite.ru", "lastsite.com"}, d)
}

func TestExtractDomains_5(t *testing.T) {
	d := ExtractDomains("")
	assert.Equal(t, []string{}, d)
}

func TestExtractDomains_6(t *testing.T) {
	d := ExtractDomains(",,")
	assert.Equal(t, []string{}, d)
}
