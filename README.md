# Тестовое задание в SberDS

Необходимо написать сервис на Go, удовлетворяющий критериям:
1.	Используем framework gin или net/http
2.	Сервис получает конфигурацию из окружения

      a.	Порт, на котором запускается приложение

      b.	Данные для подключения к БД (хост, порт, пользователь, пароль, схема)
3.	Сервис имеет connection к БД (Postgres, структура БД описана ниже)
4.	Сервис имеет 3 endpoint (обрабатываем http запросы)
      a.	GET host:port/api/v1/get_report

      •	В запросе передается report_id

      •	Необходимо получить отчет из БД, соответствующий данному id

      •	Ответ {“error_msg”: “”, “report_info”:  “”}, где error_msg – сообщение об ошибке (пустое, если ошибки не было), report_info – содержимое поля report_info

      b.	POST host:port/api/v1/set_report

      •	В запросе передаются данные для добавления в БД в формате {“report_info”: “”}, где report_info – содержимое поля report_info. сreation_time и report_id формируются автоматически

      •	Необходимо записать отчет в БД

      •	Ответ {“error_msg”: “”}, где error_msg – сообщение об ошибке (пустое, если ошибки не было)

      c.	GET host:port/api/v1/get_observation_time

      •	В запросе передается model_id

      •	Необходимо получить максимальное промежуток времени (в кол-ве дней) в течении которых не было поступлений отчетов по модели (если запись по модели была всего одна, то считаем по текущий день)

      •	Ответ {“error_msg”: “”, “max_observation_period”: “”}, где error_msg – сообщение об ошибке (пустое, если ошибки не было), max_observation_period – максимальный промежуток во время наблюдения модели

      
Дополнительное задание

1.	Добавить логирование

      a.	Время пришедшего запроса

      b.	Отправителя запроса

      c.	Endpoint запроса

2.	Написать Dockerfile

