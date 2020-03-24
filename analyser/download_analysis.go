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
)

func RequestApiAnalysis(site core.Site, api prcy.Api) (resSite core.Site, action RequestAction) {
	res, err := api.RequestAnalysis(site.Domain)
	if err != nil {
		// сетевая ошибка, ошибка api и т.п. => останавливаемся
		log.Fatalf("[FATAL] [%s] Network problem: %s", site.Domain, err.Error())
	}

	site.PrCyAnalysis = res

	if res.HttpStatus == http.StatusUnauthorized {
		// наткнулись на лимиты по api => ждём и ещё раз повторяем
		log.Printf("[LIMIT] [%s] Error: %s", site.Domain, res.Error)
		return site, ActionRepeat
	}

	if res.HttpStatus != http.StatusOK {
		// ошибка api => сохраняем результат и идём дальше
		log.Printf("[WARN] [%s] Status: %d, Error: %s", site.Domain, res.HttpStatus, res.Error)
		return site, ActionNext
	}

	if res.PublicStatistics != nil {
		log.Printf("[INFO] [%s] Analysis was downloaded, Date: %s", site.Domain, res.PublicStatistics.Updated)
	} else {
		log.Printf("[INFO] [%s] No analysys", site.Domain)
	}
	return site, ActionNext
}

func DownloadSitesAnalysis(sites <-chan core.Site, config *core.Config, api prcy.Api) <-chan core.Site {
	out := make(chan core.Site)

	go func() {
		count := 0
		for site := range sites {
			resSite, action := RequestApiAnalysis(site, api)

			// что-то пошло не так (например, упёрлись в лимиты) => ждём и повторяем запрос
			if action == ActionRepeat {
				sleep(config, 1)
				resSite, action = RequestApiAnalysis(site, api)

				if action == ActionRepeat {
					sleep(config, 2)
					resSite, action = RequestApiAnalysis(site, api)

					if action == ActionRepeat {
						sleep(config, 10)
						resSite, action = RequestApiAnalysis(site, api)

						if action == ActionRepeat {
							// не удалось повторно запросить успешный результат => останавливаемся
							log.Fatalf("[FATAL] [LIMIT] [%s] Too many retries", site.Domain)
						}
					}
				}
			}

			// результат получен => идём и идём дальше
			if action == ActionNext {
				count += 1
				out <- resSite
				sleep(config, 1)
			}
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
