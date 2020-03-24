package main

import (
	"siteanalyser/core"
	"siteanalyser/uploader"
)

// Утилита для загрузки списка сайтов для анализа.
//
// На вход принимает CSV файл `./input.csv` со списком доменов для анализа.
// На выходе формирует папку `./sites/`, где для каждого домена создаётся JSON файл
// с результом анализа. По умолчанию пустой. В нём же сохраняется код ошибки, если анализ не удался.
// Важный параметр --column из которого берёт номер колонки с сайтами. По умолчанию колонка №0.
func main() {
	config := core.ParseConfigFromFlags()
	storage := core.NewSiteStorage(config.SitesDir)
	uploader.Upload(config, storage)
}
