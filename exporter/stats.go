package exporter

type Stats struct {
	Total int // всего сайтов в хранилище

	// сумма должна равняться Total
	Status0     int // не было запроса заотчётом
	Status200   int // отчёт скачен
	Status403   int // доступ ограничен, ограничения по api
	Status404   int // нет базового анализа
	StatusOther int // иные ошибки

	// сумма должна равняться Total
	HasVisits     int // есть данные по PublicStatistics
	NoVisits      int // отчёт скачен, но данные по PublicStatistics отсутсвуют
	UnknownVisits int // отчёт не скачен или была ошибка

	// сумма должна равняться HasVisits
	Sources map[string]int // статистика по источникам данных
}
