package core

import (
	"strings"
)

// Парсит домены из строчек вида "https://www.site.ru/profiles/, othersite.com"
func ExtractDomains(input string) []string {
	domains := make([]string, 0)

	input = strings.ReplaceAll(input, ";", ",")
	values := strings.Split(input, ",")
	for _, value := range values {
		value = strings.TrimSpace(value)
		value = strings.TrimPrefix(value, "http://")
		value = strings.TrimPrefix(value, "https://")
		values2 := strings.Split(value, "/")
		if len(values2) > 0 && len(values2[0]) > 0 {
			domains = append(domains, values2[0])
		}
	}

	return domains
}
