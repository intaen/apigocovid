package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"apigocovid/src/domain"
	"apigocovid/src/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

const (
	minIdleConns = 200
	poolSize     = 12000
	poolTimeout  = 240
)

type Redis struct {
	client *redis.Client
	ctx    *gin.Context
}

// Returns new redis client
func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.host") + ":" + viper.GetString("redis.port"),
		MinIdleConns: minIdleConns,
		PoolSize:     poolSize,
		PoolTimeout:  time.Duration(poolTimeout) * time.Second,
		Password:     viper.GetString("redis.password"), // no password set
		DB:           viper.GetInt("redis.db"),          // use default DB
	})

	return client
}

func CreateClient(client *redis.Client, ctx *gin.Context) domain.ClientRedis {
	return &Redis{client: client, ctx: ctx}
}

// GetRedisValue return byte
func (r *Redis) GetRedisValue(key string) ([]byte, error) {
	if r.client == nil {
		return nil, errors.New("nil redis pointer")
	}
	res, err := r.client.Get(r.ctx, key).Result()
	if err != nil || string(res) == "" {
		return nil, err
	}

	data := struct {
		Values json.RawMessage `json:"values"`
		Expire int             `json:"expire"`
	}{}

	err = json.Unmarshal([]byte(res), &data)
	if err != nil {
		return nil, err
	}
	return data.Values, nil
}

// SetRedisValue with Set
func (r *Redis) SetRedisValue(key string, value interface{}, expireTime int) {
	if r.client == nil {
		return
	}

	ttl := time.Duration(expireTime) * time.Second
	expireUnix := int(time.Now().Unix())

	valueString := utils.ConvertString(value)
	if valueString == "" {
		valueString = `""`
	}

	redisString := fmt.Sprintf(`{"values" : %s, "expire" : %d}`, valueString, expireUnix)
	r.client.Set(r.ctx, key, redisString, ttl)
}

// DelRedisValue with Set
func (r *Redis) DelRedisValue(key string) error {
	if r.client == nil {
		return errors.New("nil redis pointer")
	}

	res, err := r.client.Del(r.ctx, key).Result()
	if err != nil || string(fmt.Sprintf("%v", res)) == "" {
		return err
	}
	fmt.Printf("res del %+v \n", res)

	return nil
}
