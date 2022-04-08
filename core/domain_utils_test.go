package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractDomains_Url_1(t *testing.T) {
	d := ExtractDomains("https://www.somesite.ru/profile/7798752")
	assert.Equal(t, []string{"www.somesite.ru"}, d)
}

func TestExtractDomains_Url_2(t *testing.T) {
	d := ExtractDomains("http://www.somesite.ru?param1=2")
	assert.Equal(t, []string{"www.somesite.ru"}, d)
}

func TestExtractDomains_Url_3(t *testing.T) {
	d := ExtractDomains("http://www.somesite.ru#hash")
	assert.Equal(t, []string{"www.somesite.ru"}, d)
}

func TestExtractDomains_Url_4(t *testing.T) {
	d := ExtractDomains("www.somesite.ru\\")
	assert.Equal(t, []string{"www.somesite.ru"}, d)
}

func TestExtractDomains_Separator_1(t *testing.T) {
	d := ExtractDomains("a.ru, b.ru; c.ru\nd.ru")
	assert.Equal(t, []string{"a.ru", "b.ru", "c.ru", "d.ru"}, d)
}

func TestExtractDomains_Separator_2(t *testing.T) {
	d := ExtractDomains(" www.somesite.ru  ")
	assert.Equal(t, []string{"www.somesite.ru"}, d)
}

func TestExtractDomains_Separator_3(t *testing.T) {
	d := ExtractDomains(",,;")
	assert.Equal(t, []string{}, d)
}

func TestExtractDomains_Email(t *testing.T) {
	d := ExtractDomains("mail@somesite.ru")
	assert.Equal(t, []string{"somesite.ru"}, d)
}

func TestExtractDomains_Empty(t *testing.T) {
	d := ExtractDomains("")
	assert.Equal(t, []string{}, d)
}

func TestExtractDomains_Dots(t *testing.T) {
	d := ExtractDomains("a.ru. b.ru .c.ru")
	assert.Equal(t, []string{"b.ru"}, d)
}

func TestExtractDomains_Utf(t *testing.T) {
	d := ExtractDomains("салют.рф")
	assert.Equal(t, []string{"салют.рф"}, d)
}
