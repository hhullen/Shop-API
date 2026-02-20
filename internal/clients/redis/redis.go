package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"shopapi/internal/supports"
	"time"

	go_redis "github.com/redis/go-redis/v9"
)

//go:generate mockgen -destination=redis_mock.go -package=redis github.com/redis/go-redis/v9 UniversalClient

const (
	requestTimeout    = time.Millisecond * 200
	defaultExpiration = time.Minute * 10

	host_secret_path     = "./secrets/redis_host.txt"
	port_secret_path     = "./secrets/redis_port.txt"
	password_secret_path = "./secrets/redis_password.txt"
)

func NewRedisConn(ctx context.Context) (*go_redis.Client, error) {
	host, err := supports.ReadSecret(host_secret_path)
	if err != nil {
		return nil, err
	}
	port, err := supports.ReadSecret(port_secret_path)
	if err != nil {
		return nil, err
	}
	password, err := supports.ReadSecret(password_secret_path)
	if err != nil {
		return nil, err
	}

	if !supports.IsInContainer() {
		host = "localhost"
	}

	c := go_redis.NewClient(&go_redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       0,
	})

	if err := c.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return c, nil
}

type Client struct {
	client go_redis.UniversalClient
}

func NewClient(conn *go_redis.Client) *Client {
	return buildClient(conn)
}

func buildClient(c go_redis.UniversalClient) *Client {
	return &Client{client: c}
}

func (c *Client) Read(key string, v any) (bool, error) {
	ctx, cancel := getCtx()
	defer cancel()

	data, err := c.client.Get(ctx, key).Bytes()
	if err != nil {
		if errors.Is(err, go_redis.Nil) {
			return false, nil
		}
		return false, err
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *Client) Write(key string, v any) error {
	ctx, cancel := getCtx()
	defer cancel()

	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	return c.client.Set(ctx, key, data, defaultExpiration).Err()
}

func getCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), requestTimeout)
}
