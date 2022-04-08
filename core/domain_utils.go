package core

import (
	"strings"
)

// Парсит домены из строчек вида "https://www.site.ru/profiles/, othersite.com"
func ExtractDomains(input string) []string {
	domains := make([]string, 0)

	input = strings.ReplaceAll(input, ";", ",") // заменяем всевозможные разделилтели на запятую
	input = strings.ReplaceAll(input, " ", ",")
	input = strings.ReplaceAll(input, "\n", ",")

	values := strings.Split(input, ",")
	for _, value := range values {

		value = trimBefore(value, "@")  // берём только домен, если почтовый адрес
		value = trimBefore(value, "//") // очищаем схему, если урл

		value = trimAfter(value, "\\") // очищаем всё что полсе домена, если урл
		value = trimAfter(value, "/")
		value = trimAfter(value, "?")
		value = trimAfter(value, "#")

		idxFirst := strings.Index(value, ".")
		idxLast := strings.LastIndex(value, ".")
		if idxFirst > 0 && idxLast < len(value)-1 { // наличие точки внутри строки обязательно для домена
			domains = append(domains, value)
		}
	}

	return domains
}

func trimBefore(s, substr string) string {
	idx := strings.LastIndex(s, substr)
	if idx > -1 {
		return s[idx+len(substr):]
	}
	return s
}

func trimAfter(s, substr string) string {
	idx := strings.Index(s, substr)
	if idx > -1 {
		return s[:idx]
	}
	return s
}
