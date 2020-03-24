package uploader

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"siteanalyser/core"
)

// Загружает список доменов из csv-файла
func LoadDomains2(sourceFile string, columnNum int, hasColumnsRow bool) <-chan string {
	file, err := os.Open(sourceFile)
	core.FailOnError(err)
	defer file.Close()

	// используем map как set, чтобы сохранять только уникальные домены
	domains := make(map[string]bool)

	reader := csv.NewReader(file)
	core.FailOnError(err)
	reader.Comma = core.CsvSeparatorRune

	if hasColumnsRow {
		_, err := reader.Read()
		core.FailOnError(err)
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		core.FailOnError(err)

		lineDomains := core.ExtractDomains(record[columnNum])
		for _, domain := range lineDomains {
			domains[domain] = true
		}
	}

	// превращаем set в канал
	out := make(chan string, len(domains))
	for domain := range domains {
		out <- domain
	}

	close(out)
	return out
}

// Загружает список доменов из csv-файла
func LoadDomains(sourceFile string, columnNum int, hasColumnsRow bool) <-chan string {
	file, err := os.Open(sourceFile)
	core.FailOnError(err)
	defer file.Close()

	// используем map как set, чтобы сохранять только уникальные домены
	domains := make(map[string]bool)

	scanner := bufio.NewScanner(file)
	if hasColumnsRow {
		scanner.Scan()
	}

	for scanner.Scan() {
		value := core.ExtractCsvColumn(scanner.Text(), columnNum)
		lineDomains := core.ExtractDomains(value)

		for _, domain := range lineDomains {
			domains[domain] = true
		}
	}

	err = scanner.Err()
	core.FailOnError(err)

	// превращаем set в канал
	out := make(chan string, len(domains))
	for domain := range domains {
		out <- domain
	}

	close(out)
	return out
}

// Фильтрует канал с доменами на предмет уже скаченных файлов
func FilterEmptySites(sites <-chan core.Site, storage *core.SiteStorage) <-chan core.Site {
	out := make(chan core.Site)

	go func() {
		for site := range sites {
			// нет анализа
			if site.PrCyAnalysis == nil {
				out <- site
			}
		}
		close(out)
	}()

	return out
}

func Upload(config *core.Config, storage *core.SiteStorage) {
	domains := LoadDomains2(config.InputFile, config.SiteColumn, config.HasHeader)

	sites := core.LoadSitesByDomains(domains, storage)
	sites = FilterEmptySites(sites, storage)
	sites = core.SaveSites(sites, storage)

	// главный цикл ожидания, заканчивается только когда все сайты будут обработаны
	for range sites {
	}
}
