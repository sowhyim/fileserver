package cache

import (
	"context"
	"fileserver/model"
	"github.com/go-redis/redis/v8"
)

var engine *redis.Client

func Init(config *model.RedisConfig) {
	if config == nil {
		panic("init redis got a config with nil!")
	}
	engine = redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Username: config.Username,
		Password: config.Password, // no password set
		DB:       config.DB,       // use default DB
	})

	err := engine.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	// TODO log Init success
}
