package exporter

type Stats struct {
	Total int // всего сайтов в хранилище

	// сумма должна равняться Total
	NoStatus int         // не было запроса заотчётом
	Statuses map[int]int // статистика по статусам

	// сумма должна равняться Total, при этом HasVisits + NoVisits == Status 200
	HasVisits     int // есть данные по PublicStatistics
	NoVisits      int // отчёт скачен, но данные по PublicStatistics отсутсвуют
	UnknownVisits int // отчёт не скачен или была ошибка

	// сумма должна равняться HasVisits
	Sources map[string]int // статистика по источникам данных
}

func NewStats() *Stats {
	var stats Stats
	stats.Statuses = make(map[int]int)
	stats.Sources = make(map[string]int)
	return &stats
}
