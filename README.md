# REST API для управления книгами (Go + Gin + GORM)

Простой RESTful API сервис, разработанный на Go, для выполнения CRUD (Create, Read, Update, Delete) операций над коллекцией книг.

В качестве веб-фреймворка используется [Gin](https://github.com/gin-gonic/gin), а для взаимодействия с базой данных PostgreSQL — ORM-библиотека [GORM](https://gorm.io/).

## Ключевые возможности

-   **Полный CRUD:** Реализация всех четырёх основных операций для управления ресурсами.
-   **RESTful архитектура:** Понятные и предсказуемые эндпоинты.
-   **Работа с PostgreSQL:** Надёжное хранение данных в реляционной базе данных.
-   **Автоматическая миграция:** GORM автоматически создаёт таблицу `books` при первом запуске.

## Технологический стек

-   **Go**: Язык программирования
-   **Gin Gonic**: Веб-фреймворк
-   **GORM**: ORM для работы с базой данных
-   **PostgreSQL**: Реляционная база данных
-   **Docker**: Контейнеризация приложения и базы данных

## Быстрый старт

Для запуска проекта вам понадобится установленный [Git](https://git-scm.com/) и [Docker](https://www.docker.com/) с Docker Compose.

### 1. Клонирование репозитория

```bash
git clone https://github.com/Chaice1/REST-CRUD--GORM.git
cd REST-CRUD--GORM
```

### 2. Настройка базы данных

Проект использует базу данных PostgreSQL. Самый простой способ запустить её — использовать Docker. Создайте в корне проекта файл `docker-compose.yml` со следующим содержимым:

```yaml
version: '3.8'
services:
  db:
    image: postgres:14-alpine
    restart: always
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=postgres
      - POSTGRES_DB=postgres
    ports:
      - "5432:5432"
    volumes:
      - db_data:/var/lib/postgresql/data

volumes:
  db_data:
```

Теперь запустите контейнер с базой данных:

```bash
docker-compose up -d
```
Эта команда поднимет PostgreSQL сервер, доступный по адресу `localhost:5432`, с данными для подключения, которые уже прописаны в коде.

### 3. Запуск приложения

После запуска базы данных можно запустить Go-приложение.

```bash
# Установка зависимостей
go mod tidy

# Запуск сервера
go run main.go
```

Сервер будет запущен по адресу `http://localhost:8080`.

## Документация по API

Ниже описаны все доступные эндпоинты.

---

### **`GET /books`**
Получает список всех книг.

-   **Метод:** `GET`
-   **URL:** `/books`
-   **Ответ (200 OK):**
    ```json
    [
        {
            "ID": 1,
            "Author": "George Orwell",
            "Title": "1984",
            "Description": "A dystopian novel."
        },
        {
            "ID": 2,
            "Author": "J.R.R. Tolkien",
            "Title": "The Lord of the Rings",
            "Description": "A high-fantasy novel."
        }
    ]
    ```

---

### **`GET /books/:id`**
Получает одну книгу по её уникальному идентификатору.

-   **Метод:** `GET`
-   **URL:** `/books/1`
-   **Ответ (200 OK):**
    ```json
    {
        "ID": 1,
        "Author": "George Orwell",
        "Title": "1984",
        "Description": "A dystopian novel."
    }
    ```
-   **Ответ (404 Not Found):** Если книга не найдена.

---

### **`POST /books`**
Создаёт новую книгу.

-   **Метод:** `POST`
-   **URL:** `/books`
-   **Тело запроса (JSON):**
    ```json
    {
        "author": "Aldous Huxley",
        "title": "Brave New World",
        "description": "Another dystopian novel."
    }
    ```
-   **Ответ (201 Created):**
    ```json
    {
        "ID": 3,
        "Author": "Aldous Huxley",
        "Title": "Brave New World",
        "Description": "Another dystopian novel."
    }
    ```

---

### **`PUT /books/:id`**
Обновляет информацию о существующей книге.

-   **Метод:** `PUT`
-   **URL:** `/books/1`
-   **Тело запроса (JSON):**
    ```json
    {
        "author": "George Orwell",
        "title": "1984",
        "description": "A classic dystopian social science fiction novel and cautionary tale."
    }
    ```
-   **Ответ (200 OK):**
    ```json
    {
        "ID": 1,
        "Author": "George Orwell",
        "Title": "1984",
        "Description": "A classic dystopian social science fiction novel and cautionary tale."
    }
    ```

---

### **`DELETE /books/:id`**
Удаляет книгу по её ID.

-   **Метод:** `DELETE`
-   **URL:** `/books/1`
-   **Ответ (204 No Content):** Пустое тело ответа, означающее успешное удаление.

## Возможные улучшения

-   **Конфигурация:** Вынести строку подключения к базе данных и порт сервера в переменные окружения (например, с использованием библиотеки `viper` или `godotenv`).
-   **Валидация:** Добавить валидацию входящих данных (например, чтобы поля `author` и `title` не были пустыми).
-   **Слоистая архитектура:** Разделить код на слои (handler, service, repository) для лучшей читаемости и поддерживаемости.
-   **Обработка ошибок:** Реализовать более детальную обработку ошибок и возвращать осмысленные сообщения.
