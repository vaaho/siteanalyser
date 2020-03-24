package main

import (
	"siteanalyser/analyser"
	"siteanalyser/core"
	"siteanalyser/core/prcy"
)

// Запускает обновления анализов в PrCy. Количество обновлений ограничено,
// поэтому желательно эту процедуру делать ограниченными порциями
func main() {
	config := core.ParseConfigFromFlags()
	storage := core.NewSiteStorage(config.SitesDir)
	api := prcy.NewApi(config.PrCyToken)

	analyser.UpdateAnalysis(config, storage, api)
}
