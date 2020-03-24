package analyser

import (
	"log"
	"net/http"
	"siteanalyser/core"
	"siteanalyser/core/prcy"
)

func UpdateApiAnalysis(site core.Site, api prcy.Api) {
	status, err := api.RequestAnalysisUpdate(site.Domain)
	if err != nil {
		// сетевая ошибка, ошибка api и т.п. => останавливаемся
		log.Fatalf("[FATAL] [%s] Network problem: %s", site.Domain, err.Error())
	}

	if status != http.StatusOK {
		// наткнулись на лимиты по api => ждём и ещё раз повторяем
		log.Fatalf("[FATAL] [%s] Update error: status = %d", site.Domain, status)
	}

	log.Printf("[INFO] [%s] Start updating", site.Domain)
}

func UpdateSitesAnalysis(sites <-chan core.Site, config *core.Config, api prcy.Api) <-chan core.Site {
	out := make(chan core.Site)

	go func() {
		count := 0
		for site := range sites {
			if count >= config.UpdateCount {
				break // останавливаемся, если дастигли максиумам обновлений
			}

			UpdateApiAnalysis(site, api)
			count += 1
		}

		log.Printf("[FINISHED] Updated %d analysis", count)
		close(out)
	}()

	return out
}

func UpdateAnalysis(config *core.Config, storage *core.SiteStorage, api prcy.Api) {
	sites, _ := core.LoadSites(storage)

	sites = FilterSitesByHttpStatus(sites, http.StatusNotFound)
	sites = UpdateSitesAnalysis(sites, config, api)
	// nothing to save

	// главный цикл ожидания, заканчивается только когда все сайты будут обработаны
	for range sites {
	}
}
