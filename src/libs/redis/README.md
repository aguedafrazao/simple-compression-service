# Redis package

To use it import to your project:
```
$ go get github.com/ABuarque/simple-compression-service/src/libs/redis
```

It provides an API to publish and subscribe to topics on a redis server instance. To use it first create a pointer to Client:
```
// New returns a new Client
func New() *Client {
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost"
	}
	var redispool *redis.Pool
	redispool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", fmt.Sprintf("%s:6379", redisHost))
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
```
To publish data to a topic the method Publish should be called:
```
// Publish publishes a new key value pair
func (c *Client) Publish(key string, value string) error {
	conn := c.pool.Get()
	_, err := conn.Do("PUBLISH", key, value)
	return err
}
```
To subscribe to a topic invoke Subscribe method:
```
// Subscribe subscribes a client to a topic
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
```
