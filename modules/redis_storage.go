package modules

import (
	"github.com/go-redis/redis"
	"github.com/topfreegames/pitaya/config"
)


type RedisStorage struct {
	Base
	config *config.Config
	option *redis.Options
	Client *redis.Client
}

func NewRedisStorage(conf *config.Config) *RedisStorage {
	r := &RedisStorage{
		config: conf,
	}
	r.configure()
	return r
}

func (r *RedisStorage) configure()  {
	r.option = &redis.Options{
		Addr:               r.config.GetString("pitaya.modules.redisstorage.Client.addr"),
		DB:                 r.config.GetInt("pitaya.modules.redisstorage.Client.db"),
		DialTimeout:        r.config.GetDuration("pitaya.modules.redisstorage.Client.dialtimeout"),
	}
}

func (r *RedisStorage) Init() error {
	r.Client = redis.NewClient(r.option)
	_,err := r.Client.Ping().Result()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisStorage) Shutdown() error {
	err := r.Client.Close()
	if err != nil {
		return err
	}

	return nil
}


