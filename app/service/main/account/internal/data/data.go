package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gomodule/redigo/redis"
	"github.com/google/wire"
	"github.com/mittacy/himusic/service/account/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	db      *gorm.DB
	redisDB redis.Conn
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, error) {
	log := log.NewHelper("data", logger)

	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		log.Errorf("failed opening connection to mysql: %v", err)
		return nil, err
	}

	redisDb, err := redis.Dial(c.Redis.Network, c.Redis.Addr)
	if err != nil {
		log.Errorf("failed opening connection to redis: %v", err)
		return nil, err
	}

	return &Data{
		db:      db,
		redisDB: redisDb,
	}, nil
}
