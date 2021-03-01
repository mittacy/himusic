package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mittacy/himusic/service/paper/internal/conf"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewPaperRepo)

// Data .
type Data struct {
	mysqlDB *gorm.DB
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
	return &Data{
		mysqlDB: db,
	}, nil
}
