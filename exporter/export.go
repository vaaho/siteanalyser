package exporter

import (
	"encoding/csv"
	"io"
	"os"
	"siteanalyser/core"
	"siteanalyser/core/prcy"
)

func ReadInputByRow(sourceFile string, hasHeader bool) (<-chan []string, <-chan []string) {
	headerOut := make(chan []string)
	dataOut := make(chan []string)

	go func() {
		file, err := os.Open(sourceFile)
		core.FailOnError(err)
		defer file.Close()

		reader := csv.NewReader(file)
		core.FailOnError(err)
		reader.Comma = core.CsvSeparatorRune

		if hasHeader {
			header, err := reader.Read()
			core.FailOnError(err)
			headerOut <- header
		}
		close(headerOut)

		for {
			row, err := reader.Read() // read splited csv-line
			if err == io.EOF {        // exit on the end
				break
			}
			core.FailOnError(err)

			dataOut <- row
		}

		close(dataOut)
	}()

	return dataOut, headerOut
}

func WriteOutputByRow(destFile string, rows <-chan []string, headers <-chan []string) <-chan []string {
	out := make(chan []string)

	go func() {
		file, err := os.Create(destFile)
		core.FailOnError(err)
		defer file.Close()

		writer := csv.NewWriter(file)
		core.FailOnError(err)
		writer.Comma = core.CsvSeparatorRune
		writer.UseCRLF = true
		defer writer.Flush()

		for header := range headers {
			err := writer.Write(header)
			core.FailOnError(err)
		}

		for row := range rows {
			err := writer.Write(row)
			core.FailOnError(err)
			out <- row
		}

		close(out)
	}()

	return out
}

func AppendHeaderExportColumns(headers <-chan []string, columns ExportColumns) <-chan []string {
	out := make(chan []string)

	go func() {
		for header := range headers {
			// формируем новые колонки
			data := GetExportColumnsNames(columns)

			// расширяем строку новыми колокнами с конца
			for _, item := range data {
				header = append(header, item)
			}

			out <- header
		}
		close(out)
	}()

	return out
}

func AppendExportColumns(rows <-chan []string, domainColumnNum int, columns ExportColumns, storage *core.SiteStorage) <-chan []string {
	out := make(chan []string)

	go func() {
		for row := range rows {
			// пустой анализ. если у строки нет доменов, то анализ будет пустым
			var analysis = prcy.NewAnalysis()

			// извлекаем из колонки список доменов
			domains := core.ExtractDomains(row[domainColumnNum])

			// загружаем статистику и скалдываем её по всем сайтам
			for _, domain := range domains {
				site := storage.Load(domain)
				accumulateStatistics(analysis, site.PrCyAnalysis)
			}

			// формируем список дополнительных колонок
			data := GetExportColumnsData(columns, analysis)

			// расширяем строку новыми колокнами с конца
			for _, item := range data {
				row = append(row, item)
			}

			out <- row
		}
		close(out)
	}()

	return out
}

// В случае если одной строке соотвтевует несколько сайтов, мы суммируем статистику по всем сайтам
func accumulateStatistics(res *prcy.Analysis, item *prcy.Analysis) {
	if item == nil {
		return // нечего агрегировать
	}

	// есть не пустая статистика => она больше
	if item.PublicStatistics != nil {
		if res.PublicStatistics == nil || item.PublicStatistics.VisitsMonthly > res.PublicStatistics.VisitsMonthly {
			res.HttpStatus = item.HttpStatus
			res.Error = item.Error
			res.PublicStatistics = item.PublicStatistics
		}
		return
	}

	// стаистики нет, но есть статус
	if res.HttpStatus != 200 {
		res.HttpStatus = item.HttpStatus
		res.Error = item.Error
	}
}

func Export(config *core.Config, storage *core.SiteStorage) {
	exportColumns := ParseExportColumns(config.ExportColumns)

	rows, headers := ReadInputByRow(config.InputFile, config.HasHeader)
	headers = AppendHeaderExportColumns(headers, exportColumns)
	rows = AppendExportColumns(rows, config.SiteColumn, exportColumns, storage)
	rows = WriteOutputByRow(config.OutputFile, rows, headers)

	for _ = range rows {
	}
}
