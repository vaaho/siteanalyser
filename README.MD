# Анализатор посищаемости сайтов на основе PR-CY  

### Ипользование

1) Сперва загружаем список сайтов из фала:
```
siteanalyser.exe import -input=sites.csv -column=0
```

2) Затем скачиваем статистику PR-CY. Количество скачиваний ограничено лимитами PR-CY API. Лучше всего эту команду запускать по cron'у раз в час.
```
siteanalyser.exe analyse -prcy-token=5VlB5YUprVlF1MHk0dHgvM3lnTnRMdz0
```

3) В конце выгружаем статистику в новые колонки обратно в файл:
```
siteanalyser.exe export -input=sites.csv -column=0 -output=sites.result.csv
```

Полное описание параметров:

```
Usage: 
  siteanalyser.exe <command> [<options>]

Commands:
  import          (options: -input -column [-has-header])
        Загрузить из CSV файла список сайтов для анализа
  analyse         (options: -prcy-token)
        Скачать аналитику из сервиса PR-CY
  update-analyse  (options: -prcy-token [-update])
        Скачать обновление аналитики из сервиса PR-CY
  export          (options: -input -column -output [-export])
        Выгрузить аналитику по сайтам в CSV файл
  stats
        Показать статистику
  help
        Экран помощи

Options:
  -column int
        Номер колонки с доменом во входящем файле
  -export string
        Список колонок для экспорта (default "VisitsMonthly,PageViewsMonthly,Updated")
  -has-header
        Есть во входящем файле заголовок (default true)
  -input string
        Файл со списком сайтов для анализа (default "./sites.csv")
  -output string
        Файл с результатами анализа (default "./output.csv")
  -prcy-delay int
        Задержка между запросами к API PR-CY в милисекундах (default 1000)
  -prcy-token string
        Ключ доступа к API PR-CY
  -sites string
        Папка для сохранения результатов анализа сайтов (default "./sites/")
  -update int
        Если указан, то запускает апдейт сайтов, для которых не существует отчёта
```

### Сборка

```
go build -o bin\siteanalyser.exe main.go
```
