package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "your_postgres_host"
	port     = 5432
	user     = "your_username"
	password = "your_password"
	dbname   = "your_database_name"
)

func main() {
	// Формирование строки подключения
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Открытие соединения с базой данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Проверка соединения с базой данных
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Выполнение запроса к базе данных
	var name string
	err = db.QueryRow("SELECT name FROM users WHERE id = $1", 5).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Name from the database: %s\n", name)
}

//docker

version: '3'

services:
  app:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    depends_on:
      - postgres
    environment:
      POSTGRES_HOST: postgres
      POSTGRES_PORT: 5432
      POSTGRES_USER: your_username
      POSTGRES_PASSWORD: your_password
      POSTGRES_DB: your_database_name

  postgres:
    image: postgres:latest
    environment:
      POSTGRES_USER: your_username
      POSTGRES_PASSWORD: your_password
      POSTGRES_DB: your_database_name
    ports:
      - "5432:5432"





А вот пример Dockerfile, который устанавливает драйвер в контейнере с приложением:

# Используем официальный образ Golang
FROM golang:latest

# Устанавливаем драйвер для PostgreSQL
RUN go get -u github.com/lib/pq

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем исходный код приложения
COPY . .

# Собираем приложение
RUN go build -o main .

# Запускаем приложение
CMD ["./main"]