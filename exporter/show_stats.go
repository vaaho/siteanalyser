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
	if site.PrCyAnalysis == nil {
		stats.Status0++
		return
	}

	if site.PrCyAnalysis.HttpStatus == 403 {
		stats.Status403++
		return
	}

	if site.PrCyAnalysis.HttpStatus == 404 {
		stats.Status404++
		return
	}

	if site.PrCyAnalysis.HttpStatus != 200 {
		stats.StatusOther++
		return
	}

	// далее site.PrCyAnalysis.HttpStatus == 200
	stats.Status200++
	stats.UnknownVisits-- // вычитаем, как обещали

	if site.PrCyAnalysis.PublicStatistics == nil {
		stats.NoVisits++
		return
	}

	// далее site.PrCyAnalysis.PublicStatistics есть
	stats.HasVisits++

	// подсчитываем количество источников
	source := site.PrCyAnalysis.PublicStatistics.SourceType
	cnt, ok := stats.Sources[source]
	if ok {
		stats.Sources[source] = cnt + 1
	} else {
		stats.Sources[source] = 1
	}
}

func LogStats(stats *Stats) {
	data, _ := json.Marshal(stats)
	log.Printf("%s\n", data)
}

func ShowStats(config *core.Config, storage *core.SiteStorage) {
	sites, _ := core.LoadSites(storage)

	var stats Stats
	stats.Sources = make(map[string]int)
	for site := range sites {
		CalculateStats(site, &stats)
	}

	LogStats(&stats)
}
