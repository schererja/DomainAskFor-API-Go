package util

import (
	"log"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

type Duration int64

const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)

func GetCachedSynonyms(word string) []string {
	SynonymPrefix := "synonyms"

	var cachedSynonyms []string
	if len(word) == 0 {
		return []string{}
	}
	redisClient := connectToRedis()
	foundSynonyms := redisClient.Get(SynonymPrefix + ":" + word)
	result, err := foundSynonyms.Result()
	if err != nil {
		log.Print(err)
	}
	cachedSynonyms = strings.Split(result, ",")

	return cachedSynonyms
}

func SetCachedSynonyms(word string, synonyms []string) bool {
	SynonymPrefix := "synonyms"

	if len(synonyms) <= 0 {
		return false
	}
	var duration time.Duration = 60 * 60 * 1000 * 1000 * 1000

	synonymsCombined := strings.Join(synonyms, ",")
	log.Print(synonymsCombined)
	redisClient := connectToRedis()
	redisClient.Set(SynonymPrefix+":"+word, synonymsCombined, duration)

	return true
}

func connectToRedis() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	return client
}
