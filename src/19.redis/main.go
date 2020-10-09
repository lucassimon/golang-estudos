package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v7"
)

var client *redis.Client

func init() {
	//Initializing redis
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr:     dsn, //redis port
		Password: "",
		DB:       3,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func increment() {
	client.Set("foo", 1, 0).Result()
	client.Incr("foo").Result()
	client.Incr("bar").Result()
	client.Incr("like").Result()
	client.Incr("dislike").Result()
	client.IncrBy("download", 50).Result()

	time.Sleep(1 * time.Second)

	res, err := client.Get("foo").Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func hashes() {
	// SELECT 3
	// HSET keyname key value key value ......
	// HGETALL keyname

	client.HSet("token", "app1", "1234", "app2", "324234")
	client.HIncrBy("token", "foo", 1)
	client.HIncrBy("token", "bar", 10)

	// specific key
	res, err := client.HGet("token", "foo").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(res)

	// get all keys and values
	pages, err := client.HGetAll("token").Result()

	if err != nil {
		panic(err)
	}

	for key, value := range pages {

		fmt.Println(key, value)
	}
}

func orderedSets() {

	jhon := redis.Z{Score: 1, Member: "jhon"}
	client.ZAdd("globalscoreboard", &jhon)
	client.ZIncr("globalscoreboard", &jhon)
	client.ZIncrBy("globalscoreboard", 20, "jhon")

	doe := redis.Z{Score: 7, Member: "doe"}
	client.ZAdd("globalscoreboard", &doe)

	zoe := redis.Z{Score: 5, Member: "zoe"}
	client.ZAdd("globalscoreboard", &zoe)

	// buscar a cardinalidade
	count, err := client.ZCard("globalscoreboard").Result()

	if err != nil {
		panic(err)
	}

	fmt.Println("cardinalidade", count)

	options := redis.ZRangeBy{Min: "5", Max: "+inf", Offset: 0, Count: 3}
	res, err := client.ZRevRangeByScore("globalscoreboard", &options).Result()
	fmt.Println(res)

	scores, err := client.ZRevRangeByScoreWithScores("globalscoreboard", &options).Result()
	if err != nil {
		panic(err)
	}

	for _, score := range scores {
		fmt.Println(score.Score, score.Member)
	}

	jhonPosition := client.ZRevRank("globalscoreboard", "jhon")
	fmt.Println(jhonPosition)
	doePosition := client.ZRevRank("globalscoreboard", "doe")
	fmt.Println(doePosition)
	zoePosition := client.ZRevRank("globalscoreboard", "zoe")
	fmt.Println(zoePosition)
}

func pubSub() {
	// SELECT 3
	// SUBSCRIBE "notifications"
	// PUBLISH "notifications" "Hello world"
	// PUBLISH "notifications" '{"id": 1, "message": "Hello world"}'
	sub := client.Subscribe("notifications")

	client.Publish("notifications", "Foo").Result()

	message, err := sub.ReceiveMessage()
	if err != nil {
		panic(err)
	}

	fmt.Println(message.Channel, message.Payload)
}

func patternPubSub() {
	// SELECT 3
	// PSUBSCRIBE "notifications:*"
	// PUBLISH "notifications:security" "system down"
	sub := client.PSubscribe("notifications:*")

	message, err := sub.ReceiveMessage()
	if err != nil {
		panic(err)
	}

	fmt.Println(message.Channel, message.Payload)
}

func service(subCh, pubCh, task string, notify bool) {
	// PUBLISH user:signup foo@gmail.com
	// event source
	sub := client.Subscribe(subCh)

	// listen an event
	message, err := sub.ReceiveMessage()
	if err != nil {
		panic(err)
	}
	// do something
	fmt.Printf("[%s] > %s : %s \n", message.Channel, task, message.Payload)

	// notify
	if notify {
		client.Publish(pubCh, message.Payload)
	}

}

func geoLocation() {
	// Blue Mosque: 41.005422 28.9746033
	// Ayasofia: 	41.008595 28.9778043
	// Tour de pise: 	43.7229559 10.3943973
	// Colisee de reme: 41.890206 12.4900203

	// GEOADD key lat long member
	// GEOADD places 41.005422 28.9746033 "blue-mosque"
	// GEOADD places 41.008595 28.9778043 "ayasofia"
	// GEOADD places 41.005422 28.9746033 "blue-mosque" 41.008595 28.9778043 "ayasofia" 43.7229559 10.3943973 "tour-de-pise" 41.890206 12.4900203 "colisee-de-rome"

	// distancia entre 2 pontos
	// GEODIST places "blue-mosque" "ayasofia" km

	// min lat lng
	// GEOHASH places "blue-mosque"

	// nearby
	// GEORADIUS places 41.005422 28.9746033 20 km WITHCOORD
	// GEORADIUS places 41.005422 28.9746033 20 km WITHCOORD COUNT 5
	// GEORADIUS places 41.005422 28.9746033 20 km WITHCOORD COUNT 7 WITHDIST

	// blueMosque := redis.GeoLocation{Name: "blue-mosque", Latitude: 41.005422, Longitude: 28.9746033}
	// client.GeoAdd("places", &blueMosque)

	res, err := client.GeoHash("places", "blue-mosque").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	dist, err := client.GeoDist("places", "blue-mosque", "ayasofia", "km").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(dist)

	query := redis.GeoRadiusQuery{
		Radius:    1000,
		Unit:      "km",
		WithDist:  true,
		WithCoord: true,
	}
	near, err := client.GeoRadius("places", 43.7229559, 10.3943973, &query).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(near)
}

func main() {
	increment()
	hashes()
	orderedSets()
	pubSub()
	// patternPubSub()

	// go service("user:signup", "user:confirm-email", "sending confirmation email", true)
	// go service("user:confirm-email", "user:welcome-email", "sending welcome email", true)
	// go service("user:welcome-email", "user:activation", "activating account", false)

	// fmt.Scanln()

	geoLocation()
}
