package core

import "strings"

const CsvSeparator = ";"
const CsvSeparatorRune = ';'

func ExtractCsvColumn(line string, colNum int) string {
	values := strings.Split(line, CsvSeparator)
	if colNum < len(values) {
		value := strings.TrimSpace(values[colNum])
		value = strings.Trim(value, "\"")
		return value
	}
	return ""
}
