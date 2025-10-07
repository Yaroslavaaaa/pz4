## Практическая работа №4. Вуйко Ярослава, ЭФМО-01-25

Маршрутизация с chi (альтернатива — gorilla/mux). Создание небольшого CRUD-сервиса «Список задач». 07.10.2025

## Описание проекта

Этот проект демонстрирует создание RESTful API сервера на Go со следующими возможностями:
- Создание и получение задач
- Изменение задачи
- Удаление задачи
- Middleware для логирования запросов

## Структура проекта

```
.
├── cmd/
│   └── server/
│       └── main.go              
├── internal/
│   ├── api/
│   │   ├── handlers_test.go
│   │   ├── handlers.go          
│   │   ├── middleware.go        
│   │   └── responses.go         
│   └── storage/
│       └── memory.go            
├── go.mod 
├── go.sum                      
├── README.md                    
└── requests.md                  
```


### Команда запуска:
```
go run ./cmd/server
```

### Сборка и запуск бинарного файла
```
go build -o server.exe ./cmd/server
.\server.exe
```

По умолчанию порт 8080, можно изменить с помощью переменной окружения $env:PORT


## Доступные эндпоинты:

http://localhost:8080/health
<img width="679" height="617" alt="image" src="https://github.com/user-attachments/assets/5181d1e8-32d7-4b06-aa90-63919f607b2b" />



http://localhost:8080/tasks(GET)
<img width="772" height="659" alt="image" src="https://github.com/user-attachments/assets/3468502e-5e83-4c18-823c-d5277a80595c" />




http://localhost:8080/tasks(POST)
<img width="992" height="605" alt="image" src="https://github.com/user-attachments/assets/34501915-b540-42c2-b135-c1c38c21896c" />



http://localhost:8080/tasks/{id}(GET)
<img width="784" height="617" alt="image" src="https://github.com/user-attachments/assets/cea89db8-290c-410d-b61a-4fe7798c1d1a" />



http://localhost:8080/tasks/{id}(PATCH)
<img width="855" height="639" alt="image" src="https://github.com/user-attachments/assets/318a4c50-ec37-4814-b8fb-5cf14c35dee5" />



http://localhost:8080/tasks/{id}(DELETE)
<img width="1064" height="578" alt="image" src="https://github.com/user-attachments/assets/026dd6fd-1a14-4738-b469-19e393f525e9" />




## Логирование

Сервер автоматически логирует все входящие запросы в формате:
ГГГГ/ММ/ДД ЧЧ:ММ:СС метод эндпоинт статус время

```
2025/10/01 15:15:15 POST /tasks 201 1.2ms
```
<img width="450" height="225" alt="image" src="https://github.com/user-attachments/assets/5cd2c166-a99f-43d8-a97d-448872710fd3" />






Роутер, здесь происходит версионирование эндпоинтов. Так же при необходимости можно добавить и v2, v3, ...
<img width="678" height="436" alt="image" src="https://github.com/user-attachments/assets/ee8d9a4b-728f-4732-b7b8-d4e6349e48ef" />


Middleware cors
<img width="914" height="401" alt="image" src="https://github.com/user-attachments/assets/42a6a7ad-f01e-4bd3-9d49-d557994f0803" />


Middleware logger
<img width="595" height="237" alt="image" src="https://github.com/user-attachments/assets/8ac75518-40ed-4303-950b-c62ad81b3ed7" />


Доступные пути к которым можно обращаться
<img width="753" height="599" alt="image" src="https://github.com/user-attachments/assets/23ec9dcc-493f-4bbe-bb18-6ac097f15009" />


Метод обрабатывающий GET запросы для получения списка задач с поддержкой пагинации и фильтрации
<img width="710" height="306" alt="image" src="https://github.com/user-attachments/assets/15e6edc9-dc32-40e5-9c90-706fc208a820" />


Метод обрабатывающий GET запросы для получения одной конкретной задачи по её ID.
<img width="1025" height="382" alt="image" src="https://github.com/user-attachments/assets/e0ce63dc-2185-4f70-acb4-51e13c53ebab" />


Метод обрабатывающий POST запросы для создания новой задачи. Здесь же происходит валидация поля title при создании.
<img width="1030" height="568" alt="image" src="https://github.com/user-attachments/assets/77946d76-e7f7-4746-ab4b-a3bddde30ee1" />


Метод обрабатывающий PATCH запросы для изменения задачи. Здесь так же происходит валидация поля title при создании.
<img width="1058" height="574" alt="image" src="https://github.com/user-attachments/assets/fc09556c-d71c-415e-aee4-08518523f236" />


## Обработка ошибок и коды ответа
- 200 OK - успешные GET, PUT запросы
- 201 Created - успешное создание задачи
- 204 No Content - успешное удаление
- 400 Bad Request - невалидные данные, некорректный ID
- 404 Not Found - задача не найдена
- 500 Internal Server Error - серверные ошибки


## Результаты тестирования API

| Маршрут                        | Метод  | Запрос                               | Ожидаемый ответ                      | Фактический ответ |
|--------------------------------|--------|--------------------------------------|--------------------------------------|-------------------|
| `/api/v1/tasks`                | POST   | `{"title": "Test"}`                  | `201 Created` + задача               | соответсвует      |
| `/api/v1/tasks`                | GET    | –                                    | `200 OK` + список                    | соответсвует      |
| `/api/v1/tasks?page=1&limit=1` | GET    | –                                    | `200 OK` + 1 задача                  | соответсвует      |
| `/api/v1/tasks?done=true`      | GET    | –                                    | `200 OK` + выполненные               | соответсвует      |
| `/api/v1/tasks/1`              | GET    | –                                    | `200 OK` + задача                    | соответсвует      |
| `/api/v1/tasks/999`            | GET    | –                                    | `404 Not Found` + сообщение об ошибке| соответсвует      |
| `/api/v1/tasks/1`              | PATCH  | `{"title": "Test", "done": true}`    | `200 OK` + обновлённая задача        | соответсвует      |
| `/api/v1/tasks/1`              | DELETE | –                                    | `204 No Content`                     | соответсвует      |
| `/api/v1/tasks`                | POST   | `{"title": ""}`                      | `400 Bad Request`                    | соответсвует      |
| `/api/v1/tasks/abc`            | GET    | –                                    | `400 Bad Request`                    | соответсвует      |



## Получилось сделать
- CRUD операции с задачами + валидация длины title
- Пагинация при запросе
- Фильтрация выполненных задач
- Сохранение данных в json файл
- Внедрение версионирования

## Сложности
- Организация сохранения данных при операциях
- Реализвция пагинации






