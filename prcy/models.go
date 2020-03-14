package prcy

type TestModel struct {
	name    string // Название теста
	updated string // Время последнего обновления
}

type AnalysisModel struct {
	domain     string // Домен
	created    string // Время создания анализа
	started    string // Время начала обновления
	updated    string // Время окончания обновления
	isUpdating bool   // Обновляется ли анализ
	isExpired  bool   // Устарел ли анализ
}

type UpdateResultModel struct {
	started	string // Время начала обновления
}

type ErrorModel struct {
	message string // Описание ошибки
}

