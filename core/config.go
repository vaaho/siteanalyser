package core

import "flag"

type Config struct {
	InputFile     string
	OutputFile    string
	SitesDir      string
	PrCyToken     string
	PrCyDelay     int
}

// Считыает входные аргументы для программы
func ParseConfigFromFlags() *Config {
	config := &Config{}

	flag.StringVar(&config.InputFile, "input", "./input.csv", "Файл со списком сайтов для анализа")
	flag.StringVar(&config.OutputFile, "output", "./output.csv", "Файл с результатами анализа")
	flag.StringVar(&config.SitesDir, "sites", "./sites/", "Папка для сохранения результатов анализа сайтов")
	flag.StringVar(&config.PrCyToken, "prcy-token", "", "Ключ доступа к API PR-CY")
	flag.IntVar(&config.PrCyDelay, "prcy-delay", 0, "Задержка между запросами к API PR-CY в милисекундах")

	flag.Parse()

	return config
}
