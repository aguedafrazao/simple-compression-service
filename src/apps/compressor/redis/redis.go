package redis

import (
	"fmt"
	"log"
	"os"

	"github.com/garyburd/redigo/redis"
)

// Client holds a pool connection
type Client struct {
	pool *redis.Pool
}

// New returns a new pool
func New() *Client {
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost"
	}
	var redispool *redis.Pool
	redispool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", fmt.Sprintf("%s:6379", redisHost))//"localhost:6379"
		},
	}
	conn := redispool.Get()
	defer conn.Close()
	_, err := conn.Do("PING")
	if err != nil {
		log.Fatalf("can't connect to the redis database, got error:\n%v", err)
	}
	return &Client{
		pool: redispool,
	}
}

// Publish publishes a new key value pair
func (c *Client) Publish(key string, value string) error {
	conn := c.pool.Get()
	_, err := conn.Do("PUBLISH", key, value)
	return err
}

// Subscribe subscribe
func (c *Client) Subscribe(key string, msg chan []byte) error {
	rc := c.pool.Get()
	psc := redis.PubSubConn{Conn: rc}
	if err := psc.PSubscribe(key); err != nil {
		return err
	}
	go func() {
		for {
			switch v := psc.Receive().(type) {
			case redis.PMessage:
				msg <- v.Data
			}
		}
	}()
	return nil
}
