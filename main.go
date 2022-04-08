package main

import (
	"siteanalyser/analyser"
	"siteanalyser/core"
	"siteanalyser/core/prcy"
	"siteanalyser/exporter"
	"siteanalyser/importer"
)

func main() {
	config := core.ParseConfigFromFlags()
	switch config.Command {

	case core.Import:
		storage := core.NewSiteStorage(config.SitesDir)
		importer.Import(config, storage)

	case core.Analyse:
		storage := core.NewSiteStorage(config.SitesDir)
		api := prcy.NewApi(config.PrCyToken)
		analyser.DownloadAnalysis(config, storage, api)

	case core.UpdateAnalyse:
		storage := core.NewSiteStorage(config.SitesDir)
		api := prcy.NewApi(config.PrCyToken)
		analyser.UpdateAnalysis(config, storage, api)

	case core.Export:
		storage := core.NewSiteStorage(config.SitesDir)
		exporter.Export(config, storage)

	case core.Stats:
		storage := core.NewSiteStorage(config.SitesDir)
		exporter.ShowStats(config, storage)

	default:
		core.PrintCommandLineUsage()
	}
}
