package exporter

import (
	"encoding/json"
	"log"
	"siteanalyser/core"
)

func CalculateStats(site core.Site, stats *Stats) {
	stats.Total++
	stats.UnknownVisits++ // если будет статус 200, то вычтем

	// отчёт ещё не скачивался
	if site.PrCyAnalysis == nil || site.PrCyAnalysis.HttpStatus == 0 {
		stats.NoStatus++
		return
	}

	// подсчитываем количество плохих статусов
	status := site.PrCyAnalysis.HttpStatus
	cnt, ok := stats.Statuses[status]
	if ok {
		stats.Statuses[status] = cnt + 1
	} else {
		stats.Statuses[status] = 1
	}
	if status != 200 {
		return
	}

	// далее site.PrCyAnalysis.HttpStatus == 200
	stats.UnknownVisits-- // вычитаем, как обещали

	if site.PrCyAnalysis.PublicStatistics == nil {
		stats.NoVisits++
		return
	}

	// далее site.PrCyAnalysis.PublicStatistics есть
	stats.HasVisits++

	// подсчитываем количество источников
	source := site.PrCyAnalysis.PublicStatistics.SourceType
	cnt, ok = stats.Sources[source]
	if ok {
		stats.Sources[source] = cnt + 1
	} else {
		stats.Sources[source] = 1
	}
}

func LogStats(stats *Stats) {
	data, _ := json.MarshalIndent(stats, "", "  ")
	log.Printf("\n%s\n", data)
}

func ShowStats(config *core.Config, storage *core.SiteStorage) {
	sites, _ := core.LoadSites(storage)

	stats := NewStats()
	for site := range sites {
		CalculateStats(site, stats)
	}

	LogStats(stats)
}
