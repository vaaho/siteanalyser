package exporter

import (
	"log"
	"siteanalyser/core/prcy"
	"strconv"
	"strings"
)

const (
	UnknownAnalysis = "n/a" // анализ не был скачен или был скачан с ошибками
	NoAnalysis      = "-"   // нет данных в анализе
)

type ExportColumns []string

func GetExportColumnsNames(columns ExportColumns) []string {
	return columns
}

func ParseExportColumns(source string) ExportColumns {
	return strings.Split(source, ",")
}

func getAnalysisStatus(analysis *prcy.Analysis) string {
	if analysis == nil {
		return UnknownAnalysis
	}
	if analysis.HttpStatus == 404 {
		return NoAnalysis
	}
	if analysis.HttpStatus != 200 {
		return UnknownAnalysis
	}
	if analysis.PublicStatistics == nil {
		return NoAnalysis
	}
	return ""
}

func GetExportColumnsData(columns ExportColumns, analysis *prcy.Analysis) []string {
	var data []string

	for _, column := range columns {
		value := getAnalysisStatus(analysis)
		if value == "" {
			switch column {
			case "HttpStatus":
				value = strconv.Itoa(analysis.HttpStatus)
			case "Error":
				value = analysis.Error
			case "Updated":
				value = extractDate(analysis.PublicStatistics.Updated)
			case "VisitsDaily":
				value = strconv.Itoa(analysis.PublicStatistics.VisitsDaily)
			case "VisitsWeekly":
				value = strconv.Itoa(analysis.PublicStatistics.VisitsWeekly)
			case "VisitsMonthly":
				value = strconv.Itoa(analysis.PublicStatistics.VisitsMonthly)
			case "PageViewsDaily":
				value = strconv.Itoa(analysis.PublicStatistics.PageViewsDaily)
			case "PageViewsWeekly":
				value = strconv.Itoa(analysis.PublicStatistics.PageViewsWeekly)
			case "PageViewsMonthly":
				value = strconv.Itoa(analysis.PublicStatistics.PageViewsMonthly)
			default:
				log.Fatalf("[FATAL] Invalid column name: %s", column)
			}
		}
		data = append(data, value)
	}

	return data
}

// Извлекает дату и даты и времени
// ex: 2020-03-10T17:38:08+03:00 => 2020-03-10
func extractDate(value string) string {
	res := strings.Split(value, "T")
	if len(res) == 2 {
		return res[0]
	}
	return ""
}
