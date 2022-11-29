package instance

import "github.com/go-redis/redis/v8"

func (i *IAppContext) RedisSetup() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr: i.Cfg.RedisUri,
	})

	if _, err := client.Ping(i.Ctx).Result(); err != nil {
		panic(err)
	}

	return client
}
