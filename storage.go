package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
)

type Storage struct {
	Response string `json:response`
	FileID string `json:file_id`
}

var client *redis.Client

func InitStorage(url string, port int) {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", url, port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func SaveResponse(data Storage) {
	byteData,_:=json.Marshal(data)
	client.HSet("file_id", "payload", byteData)
}
