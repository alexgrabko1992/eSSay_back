# Backend для системы автоматической проверки сочинений ЕГЭ

## Репозиторий фронтенда
https://github.com/stelluchchka/check-essay-website.git

## Описание проекта

Backend-часть системы автоматической проверки сочинений ЕГЭ разработана на языке Go и предоставляет REST API для взаимодействия с веб-приложением. Система позволяет пользователям писать, проверять и публиковать сочинения, а также взаимодействовать с другими пользователями через лайки и комментарии. Модераторы могут проверять сочинения и обрабатывать апелляции.

## Основные функции

- **Регистрация и аутентификация пользователей**: Пользователи могут регистрироваться, входить в систему и управлять своими данными.
- **Управление сочинениями**: Пользователи могут создавать, редактировать, проверять и публиковать сочинения.
- **Проверка сочинений**: Сочинения автоматически проверяются по критериям ЕГЭ, а модераторы могут обрабатывать апелляции.
- **Социальные функции**: Пользователи могут ставить лайки и оставлять комментарии к сочинениям.

## Технологии

- **Язык программирования**: Go (Golang)
- **База данных**: PostgreSQL
- **Фреймворк**: Стандартная библиотека Go
- **Дополнительные технологии**: Docker, Kafka (для очередей)

## Установка и запуск

### Предварительные требования

- Установите [Go](https://golang.org/) версии 1.16 и выше.
- Установите [Docker](https://www.docker.com/) и [Docker Compose](https://docs.docker.com/compose/install/).

### Шаги для запуска

1. **Клонируйте репозиторий:**

   ```bash
   git clone https://github.com/yourusername/ege-backend-go.git
   cd ege-backend-go
   ```
2. **Настройте переменные окружения:**
Создайте файл .env в корне проекта и добавьте в него следующие переменные:
    ```env
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=postgres
    DB_PASSWORD=1234
    DB_NAME=essay
    SECRET_KEY=SECRET_KEYSECRET_KEYSECRET_KEY
    ```
3. **Соберите и запустите сервис:**
    ```bash
    go run src/cmd/app/main.go
    ```
4. **Сервис будет доступен по адресу:** http://localhost:8080

## API

### Пользователи

- POST /users/login: Создание сессии (аутентификация).
- GET /users/logout: Удаление сессии (выход из системы).
- GET /users/count: Получение количества пользователей.
- GET /users/:id : Получение информации о пользователе.
- PUT /users/:id : Изменение информации о пользователе.
- POST /users: Регистрация нового пользователя.

### Сочинения

- GET /essays: Список всех опубликованных сочинений.
- GET /essays/count: Получение количества опубликованных сочинений.
- GET /essays/:id : Чтение сочинения.
- GET /users/me/essays: Список своих сочинений.
- POST /essays: Создание черновика сочинения.
- PUT /essays/:id : Обновление сочинения.
- PUT /essays/:id /save: Проверка сочинения.
- PUT /essays/:id /appeal: Подача апелляции.
- PUT /essays/:id /publish: Публикация сочинения.

### Апелляции

- GET /appeal/essays: Список поданных на апелляцию сочинений.
- POST /results/:id : Проверка сочинения модератором.

### Лайки и комментарии

- GET /likes/:id : Количество лайков на сочинении.
- GET /comments/:id : Список комментариев под сочинением.
- POST /likes/:id : Поставить лайк на сочинение.
- POST /comments/:id : Добавить комментарий к сочинению.

### Варианты

- GET /variants/:id : Чтение текста варианта.
- GET /variants/count: Получение количества вариантов.
- POST /variants: Добавление варианта.

## Структура проекта

- src/cmd/app/main.go: Основной файл сервиса, содержащий точку входа.
- src/internal/app/: Пакет с настройкой сервиса.
- src/internal/config/: Пакет где прописаны все настройки.
- src/internal/database/: Пакет с подключением к БД.
- src/internal/middleware/: Пакет с middleware.
- src/internal/models/: Пакет с моделями данных.
- src/internal/services/: Пакет с бизнес-логикой сервиса.
- src/internal/transport/: Пакет с обработчиками HTTP-запросов.
- scripts/: SQL-файлы для миграций базы данных.

## Документация

- Научно-исследовательская работа [(НИР)](docs/ИУ5-71б_Саркисн_С_З_НИР2024.pdf)
