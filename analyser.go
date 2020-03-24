package main

import (
	"siteanalyser/analyser"
	"siteanalyser/core"
	"siteanalyser/core/prcy"
)

func main() {
	config := core.ParseConfigFromFlags()
	storage := core.NewSiteStorage(config.SitesDir)
	api := prcy.NewApi(config.PrCyToken)

	analyser.DownloadAnalysis(config, storage, api)
}
