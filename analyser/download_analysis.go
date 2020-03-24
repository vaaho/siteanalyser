package analyser

import (
	"log"
	"net/http"
	"siteanalyser/core"
	"siteanalyser/core/prcy"
)

type RequestAction string

const (
	ActionNext   = "next"
	ActionRepeat = "repeat"
	ActionStop   = "repeat"
)

func RequestApiAnalysis(site *core.Site, api prcy.Api) RequestAction {
	res, err := api.RequestAnalysis(site.Domain)
	if err != nil {
		// сетевая ошибка, ошибка api и т.п. => останавливаемся
		log.Fatalf("[FATAL] [%s] Network problem: %s", site.Domain, err.Error())
	}

	site.PrCyAnalysis = res

	if res.HttpStatus == http.StatusTooManyRequests {
		// наткнулись на ограничение => ждём и ещё раз повторяем
		log.Printf("[WARN] [LIMIT] [%s] Error: %s", site.Domain, res.Error)
		return ActionRepeat
	}

	if res.HttpStatus == http.StatusBadRequest ||
		res.HttpStatus == http.StatusNotFound {
		// домен плохой или нет анализа по домену => идём дальше
		log.Printf("[INFO] [%s] Bad domain. Status: %d, Error: %s", site.Domain, res.HttpStatus, res.Error)
		return ActionNext
	}

	if res.HttpStatus != http.StatusOK {
		// непонятная ошибка api => останавливаемся
		log.Printf("[ERROR] [%s] Status: %d, Error: %s", site.Domain, res.HttpStatus, res.Error)
		return ActionStop
	}

	return ActionNext
}

func DownloadSitesAnalysis(sites <-chan core.Site, config *core.Config, api prcy.Api) <-chan core.Site {
	out := make(chan core.Site)

	go func() {
		count := 0
		for site := range sites {
			action := RequestApiAnalysis(&site, api)

			// что-то пошло не так (например, упёрлись в лимиты) => ждём и повторяем запрос
			if action == ActionRepeat {
				sleep(config, 2)
				action = RequestApiAnalysis(&site, api)

				if action == ActionRepeat {
					sleep(config, 4)
					action = RequestApiAnalysis(&site, api)
				}
			}

			// не удалось повторно запросить результат => останавливаемся
			if action != ActionNext {
				log.Fatalf("[FATAL] [%s] [%d] Download failed", site.Domain, count+1)
			}

			// логируем факт скачки анаиза
			if site.PrCyAnalysis.PublicStatistics != nil {
				log.Printf("[INFO] [%s] [%d] Downloaded, date: %s", site.Domain, count+1, site.PrCyAnalysis.PublicStatistics.Updated)
			} else if site.PrCyAnalysis.HttpStatus == http.StatusOK {
				log.Printf("[INFO] [%s] [%d] No analysis", site.Domain, count+1)
			}

			// результат получен => ждём и идём дальше
			count += 1
			out <- site
			sleep(config, 1)
		}

		log.Printf("[FINISHED] Downloaded %d analysys", count)
		close(out)
	}()

	return out
}

func DownloadAnalysis(config *core.Config, storage *core.SiteStorage, api prcy.Api) {
	sites, _ := core.LoadSites(storage)

	sites = FilterSitesWithoutAnalysis(sites)
	sites = DownloadSitesAnalysis(sites, config, api)
	sites = core.SaveSites(sites, storage)

	// главный цикл ожидания, заканчивается только когда все сайты будут обработаны
	for range sites {
	}
}
