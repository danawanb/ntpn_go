package db

import "github.com/redis/go-redis/v9"

func NewRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		//Addr: "localhost:6379",
		Addr:     "redis:6379", // Alamat Redis
		Password: "",           // Password Redis (jika diperlukan)
		DB:       0,            // Indeks database Redis
	})

	return client
}
