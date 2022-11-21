package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type User struct {
	ID        int
	FirstName string `json:"firstname" validate:"required, gte=3"`
	LastName  string `json:"lastname" validate:"required, gte=3"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func (user User) MarshalBinary() ([]byte, error) {
	return json.Marshal(user)
}

func newRedisClient() *redis.Client {
	var host = "localhost:6379"
	var password = ""

	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
	return client
}

func handleRegister(user User) bool {

	//fmt.Println(user)

	rdb := newRedisClient()
	err := rdb.Set(context.Background(), user.Username, user, 0).Err()
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func handleLogin(user User) bool {
	fmt.Println(user)

	rdb := newRedisClient()
	val, err := rdb.Get(context.Background(), user.Username).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}

	if val == "" {
		return false
	}

	var jsonData = []byte(val)
	var userRedis User
	var errs = json.Unmarshal(jsonData, &userRedis)
	if errs != nil {
		fmt.Println(err.Error())
		return false
	}

	if user.Password != userRedis.Password {
		fmt.Println("WRONG PASSWORD")
		return false
	}

	return true
}
