package dao

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"helloworld/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDb, NewRedis, New)

// dao .
type Dao struct {
	db    *gorm.DB
	redis *redis.Client
	log   log.Logger
}

// New new a dao and return.
func New(r *redis.Client, db *gorm.DB, logger log.Logger) (d *Dao) {
	d = &Dao{
		db:    db,
		log:   logger,
		redis: r,
	}
	return d
}

//初始化数据库
func NewDb(c *conf.Data) (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       c.Database.Source,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		panic(err)
	}
	return db, nil
}

// 初始化redis客户端
func NewRedis(c *conf.Data) (*redis.Client, error) {
	r := redis.NewClient(&redis.Options{
		Addr: c.Redis.Addr,
		//Password: c.Redis.Password,
		//DB:       int(c.Redis.Db),
	})
	return r, nil
}
