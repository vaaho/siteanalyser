package core

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Command       Command
	InputFile     string
	OutputFile    string
	SitesDir      string
	HasHeader     bool
	SiteColumn    int
	PrCyToken     string // api key
	PrCyDelay     int    // задержка между запросами в милисекундах
	UpdateCount   int
	ExportColumns string
}

func newDefaultFlagSet(config *Config) *flag.FlagSet {
	flags := flag.NewFlagSet("default", flag.ExitOnError)

	flags.StringVar(&config.InputFile, "input", "./sites.csv", "Файл со списком сайтов для анализа")
	flags.StringVar(&config.OutputFile, "output", "./output.csv", "Файл с результатами анализа")
	flags.StringVar(&config.SitesDir, "sites", "./sites/", "Папка для сохранения результатов анализа сайтов")
	flags.BoolVar(&config.HasHeader, "has-header", true, "Есть во входящем файле заголовок")
	flags.IntVar(&config.SiteColumn, "column", 0, "Номер колонки с доменом во входящем файле")
	flags.StringVar(&config.PrCyToken, "prcy-token", "", "Ключ доступа к API PR-CY")
	flags.IntVar(&config.PrCyDelay, "prcy-delay", 1000, "Задержка между запросами к API PR-CY в милисекундах")
	flags.IntVar(&config.UpdateCount, "update", 0, "Если указан, то запускает апдейт сайтов, для которых не существует отчёта")
	flags.StringVar(&config.ExportColumns, "export", "VisitsMonthly,PageViewsMonthly,Updated", "Список колонок для экспорта")

	return flags
}

// Считыает входные аргументы для программы.
// Первый аргумент это команда, например import, export, analyse.
// Затем идут параметры команды. На данный момент параметры одинаковые для всех комманд.
func ParseConfigFromFlags() *Config {
	config := &Config{}

	if len(os.Args) < 2 {
		return config
	}

	flags := newDefaultFlagSet(config)
	flags.Parse(os.Args[2:])

	config.Command = ParseCommand(os.Args[1])

	return config
}

// Выводит экран помощи для консольной команды
func PrintCommandLineUsage() {
	usage := `Usage: 
  %s <command> [<options>]
Commands:
  %s (options: -input -column [-has-header])
        Загрузить из CSV файла список сайтов для анализа
  %s (options: -prcy-token)
        Скачать аналитику из сервиса PR-CY
  %s (options: -prcy-token [-update])
        Скачать обновление аналитики из сервиса PR-CY
  %s (options: -input -column -output [-export])
        Выгрузить аналитику по сайтам в CSV файл
  %s
        Показать статистику
  %s
        Экран помощи
Options:
`
	fmt.Fprintf(os.Stderr, usage, os.Args[0],
		Import,
		Analyse,
		UpdateAnalyse,
		Export,
		Stats,
		Help,
	)
	flags := newDefaultFlagSet(&Config{})
	flags.PrintDefaults()
}
