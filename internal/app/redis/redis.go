package redis

import (
	"InternetApps_5sem/internal/app/config"
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

const servicePrefix = "cargo_iss_5."

type Client struct {
	cfg    config.RedisConfig
	client *redis.Client
}

func New(cfg config.RedisConfig) (*Client, error) {
	client := &Client{}
	client.cfg = cfg

	redisClient := redis.NewClient(&redis.Options{
		Password:    cfg.Password,
		Username:    cfg.User,
		Addr:        cfg.Host + ":" + strconv.Itoa(cfg.Port),
		DB:          0,
		DialTimeout: cfg.DialTimeout,
		ReadTimeout: cfg.ReadTimeout,
	})

	client.client = redisClient

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("cant ping redis: %w", err)
	}

	return client, nil
}

func (c *Client) Close() error {
	return c.client.Close()
}
