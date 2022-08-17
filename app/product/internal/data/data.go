package data

import (
	"cpx-backend/app/product/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewProductRepo, NewDB, NewRedis)

// Data .
type Data struct {
	db    *gorm.DB      // mysql
	redis *redis.Client // redis
}

// NewData .
func NewData(logger log.Logger, gormClient *gorm.DB, redisClient *redis.Client) (*Data, func(), error) {
	log.NewHelper(log.With(logger, "module", "user-service/data"))
	data := &Data{
		gormClient,
		redisClient,
	}
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		d, _ := data.db.DB()
		if err := d.Close(); err != nil {
			log.Error("close gorm error:", err)
		}

		if err := data.redis.Close(); err != nil {
			log.Error("close redis error:", err)
		}
	}

	return data, cleanup, nil
}

func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{
		Logger: gLogger.Default.LogMode(gLogger.Info),
		NamingStrategy: schema.NamingStrategy{
			//表前缀
			TablePrefix: conf.Database.Prefix,
			//表复数禁用
			SingularTable: true,
		},
	})

	if err != nil {
		log.NewHelper(log.With(logger, "DB", "open DB failed!"))
	}
	return db
}

func NewRedis(conf *conf.Data, logger log.Logger) *redis.Client {
	return nil
}
