package analyser

import (
	"siteanalyser/core"
)

func FilterSitesWithoutAnalysis(sites <-chan core.Site) <-chan core.Site {
	out := make(chan core.Site)

	go func() {
		for site := range sites {
			// нет анализа
			if site.PrCyAnalysis != nil {
				out <- site
			}
		}
		close(out)
	}()

	return out
}

func FilterSitesByHttpStatus(sites <-chan core.Site, statuses ...int) <-chan core.Site {
	out := make(chan core.Site)

	go func() {
		for site := range sites {
			// нет анализа
			if site.PrCyAnalysis == nil {
				continue
			}

			// матчим статус
			for _, status := range statuses {
				if status == site.PrCyAnalysis.HttpStatus {
					out <- site
					break
				}
			}
		}
		close(out)
	}()

	return out
}
