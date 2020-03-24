package exporter

import (
	"bufio"
	"os"
	"siteanalyser/core"
)

func ReadInputByLine(sourceFile string) <-chan string {
	out := make(chan string)

	go func() {
		file, err := os.Open(sourceFile)
		core.FailOnError(err)
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			out <- line
		}

		close(out)
	}()

	return out
}

func WriteOutputByLine(destFile string, lines <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		file, err := os.Create(destFile)
		core.FailOnError(err)
		defer file.Close()

		for line := range lines {
			_, err := file.WriteString(line + "\n")
			core.FailOnError(err)
			out <- line
		}

		close(out)
	}()

	return out
}

/*
func AppendExportColumns(lines <-chan string, columns ExportColumns, storage *core.SiteStorage, hasColumnsRow bool) <-chan string {
	out := make(chan string)

	go func() {
		if hasColumnsRow {
			firstLine := <-lines
			outLine := firstLine + GetExportColumnsNames(columns)
			//log.Printf("[LINE] " + outLine)
			out <- outLine
		}
		for line := range lines {
			outLine := line + GetExportColumnsData(line, columns)
			//log.Printf("[LINE] " + outLine)
			out <- outLine
		}
		close(out)
	}()

	return out
}
*/

func Export(config *core.Config, storage *core.SiteStorage) {
	//columns := NewExportColumns(config.ExportColumns)
	//if columns.Total == 0 {
	//	log.Printf("[EXPORT] Нечего экспортировать")
	//	return
	//}
	//
	//lines := ReadInputByLine(config.InputFile)
	//lines = AppendExportColumns(lines, columns, storage, !config.NoColumnsRow)
	//lines = WriteOutputByLine(config.OutputFile, lines)
	//
	//for _ = range lines {
	//}
}
