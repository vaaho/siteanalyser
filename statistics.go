package main

import (
	"siteanalyser/core"
	"siteanalyser/exporter"
)

func main() {
	config := core.ParseConfigFromFlags()
	storage := core.NewSiteStorage(config.SitesDir)

	exporter.ShowStats(config, storage)
}
