package exporter

import (
	"log"
	"siteanalyser/core"
	"siteanalyser/core/prcy"
	"strconv"
	"strings"
)

const (
	UnknownAnalysis = "n/a" // анализ не был скачен или был скачан с ошибками
	NoAnalysis      = "no"  // нет данных в анализе
)

type ExportColumns []string

func GetExportColumnsNames(columns ExportColumns) string {
	result := core.CsvSeparator + strings.Join(columns, core.CsvSeparator)
	return result
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

func GetExportColumnsData(columns ExportColumns, analysis *prcy.Analysis) string {
	var data []string

	for _, column := range columns {
		value := getAnalysisStatus(analysis)
		if value == "" {
			switch column {
			case "SourceType":
				value = analysis.PublicStatistics.SourceType
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

	result := core.CsvSeparator + strings.Join(data, core.CsvSeparator)
	return result
}
