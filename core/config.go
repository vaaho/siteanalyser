package core

import "flag"

type Config struct {
	InputFile   string
	OutputFile  string
	SitesDir    string
	HasHeader   bool
	SiteColumn  int
	PrCyToken   string // api key
	PrCyDelay   int    // задержка между запросами в милисекундах
	UpdateCount int
}

// Считыает входные аргументы для программы
func ParseConfigFromFlags() *Config {
	config := &Config{}

	flag.StringVar(&config.InputFile, "input", "./input.csv", "Файл со списком сайтов для анализа")
	flag.StringVar(&config.OutputFile, "output", "./output.csv", "Файл с результатами анализа")
	flag.StringVar(&config.SitesDir, "sites", "./sites/", "Папка для сохранения результатов анализа сайтов")
	flag.BoolVar(&config.HasHeader, "has-header", true, "Есть во входящем файле заголовок")
	flag.IntVar(&config.SiteColumn, "column", 0, "Номер колонки с доменом во входящем файле")
	flag.StringVar(&config.PrCyToken, "prcy-token", "", "Ключ доступа к API PR-CY")
	flag.IntVar(&config.PrCyDelay, "prcy-delay", 1000, "Задержка между запросами к API PR-CY в милисекундах")
	flag.IntVar(&config.UpdateCount, "update", 0, "Если указан, то запускает апдейт сатов, для которых не существует отчёта")

	flag.Parse()

	return config
}
