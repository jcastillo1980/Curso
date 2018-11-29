package main

import (
	"log"
	"time"

	"github.com/go-redis/redis"
)

func init() {
	log.Println("En Marcha go Redis !!")
}

func isConnectRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "oficina.xuitec.com:20053",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	defer client.Close()

	log.Println(pong, err)
}

func setKey(key, value string) {
	client := redis.NewClient(&redis.Options{
		Addr:     "oficina.xuitec.com:20053",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Println("Error: ", err)
	}

	defer client.Close()

	err = client.Set(key, value, 0).Err()
	if err != nil {
		log.Println(err)
	}
}

func getKey(key string) string {
	client := redis.NewClient(&redis.Options{
		Addr:     "oficina.xuitec.com:20053",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Println("Error: ", err)
	}

	defer client.Close()

	vari, err := client.Get(key).Result()
	if err != nil {
		log.Println(err)
		return ""
	}

	return vari

}

func sub(canal string) {
	client := redis.NewClient(&redis.Options{
		Addr:     "oficina.xuitec.com:20053",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Println("Error: ", err)
	}

	defer client.Close()

	pubsub := client.Subscribe(canal)

	_, err = pubsub.Receive()
	if err != nil {
		log.Println("Error SUB: ", err)
	}

	ch := pubsub.Channel()

	time.AfterFunc(2*time.Minute, func() {
		_ = pubsub.Close()
		log.Println("Final !!")
	})

	for msg := range ch {
		log.Printf("[%s]->[%s]\r\n", msg.Channel, msg.Payload)
	}

}

func addList(key string, values ...string) {
	client := redis.NewClient(&redis.Options{
		Addr:     "oficina.xuitec.com:20053",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Println("Error: ", err)
	}

	defer client.Close()

	_, err = client.RPush(key, values).Result()
	if err != nil {
		log.Println("Error: ", err)
	}
}

func extList(key string) string {
	client := redis.NewClient(&redis.Options{
		Addr:     "oficina.xuitec.com:20053",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Println("Error: ", err)
	}

	defer client.Close()

	val, err := client.LPop(key).Result()
	if err != nil {
		log.Println("Error: ", err)
		return ""
	}

	return val
}

func bloqExtList(key string) []string {
	client := redis.NewClient(&redis.Options{
		Addr:     "oficina.xuitec.com:20053",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Println("Error: ", err)
	}

	defer client.Close()

	val, err := client.BLPop(0, key).Result()
	if err != nil {
		log.Println("Error: ", err)
		return []string{""}
	}

	return val

}

func main() {
	setKey("variable1", "Esto es la variable 1")
	setKey("Mitiempo", time.Now().String())

	log.Printf("K1: [%s]\r\n", getKey("K1"))

	// sub("CC")
	// addList("ll", fmt.Sprintf("%d", time.Now().Unix()))

	for {
		log.Println(bloqExtList("ll"))
	}
}
