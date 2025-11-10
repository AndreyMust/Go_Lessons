//go get -u github.com/go-redis/redis/v8
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Подключение к Redis
	client := redis.NewClient(&redis.Options{
//		Addr:     "91.217.80.209:6379", // адрес Redis-сервера
		Addr:     "192.168.1.9:6379", // адрес Redis-сервера
		Password: "",                // пароль (если установлен)
		DB:       0,                 // номер базы данных
	})

	defer client.Close()

	// Пример работы с кешем
	key := "example_key"

	fmt.Println("Start")

	value, err := client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		// Ключ не найден в кеше, проведем вычисление и сохраним результат
		result := "example_key_Values999"
		fmt.Println("Try to SET value")

		err := client.Set(context.Background(), key, result, 10*time.Second).Err()
		if err != nil {
			log.Fatal(err)
		}
		value = result
		fmt.Println("Calculated value:", value)
	} else if err != nil {
		// Обработка ошибок при получении данных из кеша
		log.Fatal(err)
	} else {
		// Ключ найден в кеше
		fmt.Println("Value from cache:", value)
	}

	fmt.Println("Start 100 Gets")
	// Засечем время выполнения запроса
	startTime := time.Now()

	get1, _ := client.Get(context.Background(), key).Result()

	fmt.Println("Start Get1:", get1)

	// Пример выполнения нескольких запросов для демонстрации высокой скорости работы с кеш-памятью
	for i := 0; i < 500; i++ {
		_, err := client.Get(context.Background(), key).Result()
		//fmt.Println("Get: ", i, g)
		if err != nil && err != redis.Nil {
			log.Fatal(err)
			fmt.Println("Error")
		}
	}

	// Выведем время выполнения
	elapsed := time.Since(startTime)
	fmt.Printf("Elapsed time for 500 cache lookups: %s\n", elapsed)
}